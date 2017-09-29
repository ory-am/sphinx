/**
 * Hydra OAuth2 & OpenID Connect Server
 * Please refer to the user guide for in-depth documentation: https://ory.gitbooks.io/hydra/content/   Hydra offers OAuth 2.0 and OpenID Connect Core 1.0 capabilities as a service. Hydra is different, because it works with any existing authentication infrastructure, not just LDAP or SAML. By implementing a consent app (works with any programming language) you build a bridge between Hydra and your authentication infrastructure. Hydra is able to securely manage JSON Web Keys, and has a sophisticated policy-based access control you can use if you want to. Hydra is suitable for green- (new) and brownfield (existing) projects. If you are not familiar with OAuth 2.0 and are working on a greenfield project, we recommend evaluating if OAuth 2.0 really serves your purpose. Knowledge of OAuth 2.0 is imperative in understanding what Hydra does and how it works.   The official repository is located at https://github.com/ory/hydra   ### Important REST API Documentation Notes  The swagger generator used to create this documentation does currently not support example responses. To see request and response payloads click on **\"Show JSON schema\"**: ![Enable JSON Schema on Apiary](https://storage.googleapis.com/ory.am/hydra/json-schema.png)   The API documentation always refers to the latest tagged version of ORY Hydra. For previous API documentations, please refer to https://github.com/ory/hydra/blob/<tag-id>/docs/api.swagger.yaml - for example:  0.9.13: https://github.com/ory/hydra/blob/v0.9.13/docs/api.swagger.yaml 0.8.1: https://github.com/ory/hydra/blob/v0.8.1/docs/api.swagger.yaml
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.HydraOAuth2OpenIdConnectServer);
  }
}(this, function(expect, HydraOAuth2OpenIdConnectServer) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new HydraOAuth2OpenIdConnectServer.SwaggerJwkCreateSet();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('SwaggerJwkCreateSet', function() {
    it('should create an instance of SwaggerJwkCreateSet', function() {
      // uncomment below and update the code to test SwaggerJwkCreateSet
      //var instane = new HydraOAuth2OpenIdConnectServer.SwaggerJwkCreateSet();
      //expect(instance).to.be.a(HydraOAuth2OpenIdConnectServer.SwaggerJwkCreateSet);
    });

    it('should have the property body (base name: "Body")', function() {
      // uncomment below and update the code to test the property body
      //var instane = new HydraOAuth2OpenIdConnectServer.SwaggerJwkCreateSet();
      //expect(instance).to.be();
    });

    it('should have the property set (base name: "set")', function() {
      // uncomment below and update the code to test the property set
      //var instane = new HydraOAuth2OpenIdConnectServer.SwaggerJwkCreateSet();
      //expect(instance).to.be();
    });

  });

}));
