-- Migration generated by the command below; DO NOT EDIT.
-- hydra:generate hydra migrations gen

CREATE INDEX hydra_oauth2_flow_login_session_id_idx ON hydra_oauth2_flow (login_session_id);
