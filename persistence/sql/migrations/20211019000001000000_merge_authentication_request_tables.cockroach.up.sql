-- Migration generated by the command below; DO NOT EDIT.
-- hydra:generate hydra migrations gen
-- TODO use nulls instead of default values
-- TODO split the consent initialized state into two states: login used and consent initialized
CREATE TABLE hydra_oauth2_flow
(
    login_challenge           character varying(40) NOT NULL,
    requested_scope           text NOT NULL,
    login_verifier            character varying(40) NOT NULL,
    login_csrf                character varying(40) NOT NULL,
    subject                   character varying(255) NOT NULL,
    request_url               text NOT NULL,
    login_skip                boolean NOT NULL,
    client_id                 character varying(255) NOT NULL,
    requested_at              timestamp without time zone DEFAULT now() NOT NULL,
    oidc_context              text NOT NULL,
    login_session_id          character varying(40),
    requested_at_audience     text DEFAULT ''::text,
    login_initialized_at      timestamp without time zone,

    state                     INTEGER      NOT NULL DEFAULT 0,

    login_remember boolean,
    login_remember_for integer,
    login_error text,
    acr text,
    login_authenticated_at timestamp without time zone,
    login_was_used boolean,
    forced_subject_identifier character varying(255) DEFAULT ''::character varying,
    context text DEFAULT '{}'::text,
    amr text DEFAULT ''::text,

    consent_challenge_id character varying(40),
    consent_verifier character varying(40),
    consent_skip boolean DEFAULT false NOT NULL,
    consent_csrf character varying(40),

    granted_scope text,
    consent_remember boolean DEFAULT false NOT NULL,
    consent_remember_for integer,
    consent_error text,
    session_access_token text DEFAULT '{}'::text NOT NULL,
    session_id_token text DEFAULT '{}'::text NOT NULL,
    consent_was_used boolean DEFAULT false NOT NULL,
    granted_at_audience text DEFAULT ''::text,
    consent_handled_at timestamp without time zone

    CHECK (
        state = 128 OR
        state = 129 OR
        state = 1 OR
        (state = 2 AND (
            login_remember IS NOT NULL AND
            login_remember_for IS NOT NULL AND
            login_error IS NOT NULL AND
            acr IS NOT NULL AND
            login_was_used IS NOT NULL AND
            context IS NOT NULL AND
            amr IS NOT NULL
        )) OR
        (state = 3 AND (
            login_remember IS NOT NULL AND
            login_remember_for IS NOT NULL AND
            login_error IS NOT NULL AND
            acr IS NOT NULL AND
            login_was_used IS NOT NULL AND
            context IS NOT NULL AND
            amr IS NOT NULL
        )) OR
        (state = 4 AND (
            login_remember IS NOT NULL AND
            login_remember_for IS NOT NULL AND
            login_error IS NOT NULL AND
            acr IS NOT NULL AND
            login_was_used IS NOT NULL AND
            context IS NOT NULL AND
            amr IS NOT NULL AND

            consent_challenge_id IS NOT NULL AND
            consent_verifier IS NOT NULL AND
            consent_skip IS NOT NULL AND
            consent_csrf IS NOT NULL
        )) OR
        (state = 5 AND (
            login_remember IS NOT NULL AND
            login_remember_for IS NOT NULL AND
            login_error IS NOT NULL AND
            acr IS NOT NULL AND
            login_was_used IS NOT NULL AND
            context IS NOT NULL AND
            amr IS NOT NULL AND

            consent_challenge_id IS NOT NULL AND
            consent_verifier IS NOT NULL AND
            consent_skip IS NOT NULL AND
            consent_csrf IS NOT NULL AND

            granted_scope IS NOT NULL AND
            consent_remember IS NOT NULL AND
            consent_remember_for IS NOT NULL AND
            consent_error IS NOT NULL AND
            session_access_token IS NOT NULL AND
            session_id_token IS NOT NULL AND
            consent_was_used IS NOT NULL
        ))
    )
);
