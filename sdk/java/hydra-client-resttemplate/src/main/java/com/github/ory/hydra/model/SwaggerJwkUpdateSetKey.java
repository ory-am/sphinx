/*
 * ORY Hydra
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * OpenAPI spec version: latest
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package com.github.ory.hydra.model;

import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;
import com.github.ory.hydra.model.JSONWebKey;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

/**
 * SwaggerJwkUpdateSetKey
 */
@javax.annotation.Generated(value = "io.swagger.codegen.languages.JavaClientCodegen", date = "2019-04-02T13:01:09.037+02:00")
public class SwaggerJwkUpdateSetKey {
  @JsonProperty("Body")
  private JSONWebKey body = null;

  @JsonProperty("kid")
  private String kid = null;

  @JsonProperty("set")
  private String set = null;

  public SwaggerJwkUpdateSetKey body(JSONWebKey body) {
    this.body = body;
    return this;
  }

   /**
   * Get body
   * @return body
  **/
  @ApiModelProperty(value = "")
  public JSONWebKey getBody() {
    return body;
  }

  public void setBody(JSONWebKey body) {
    this.body = body;
  }

  public SwaggerJwkUpdateSetKey kid(String kid) {
    this.kid = kid;
    return this;
  }

   /**
   * The kid of the desired key in: path
   * @return kid
  **/
  @ApiModelProperty(required = true, value = "The kid of the desired key in: path")
  public String getKid() {
    return kid;
  }

  public void setKid(String kid) {
    this.kid = kid;
  }

  public SwaggerJwkUpdateSetKey set(String set) {
    this.set = set;
    return this;
  }

   /**
   * The set in: path
   * @return set
  **/
  @ApiModelProperty(required = true, value = "The set in: path")
  public String getSet() {
    return set;
  }

  public void setSet(String set) {
    this.set = set;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SwaggerJwkUpdateSetKey swaggerJwkUpdateSetKey = (SwaggerJwkUpdateSetKey) o;
    return Objects.equals(this.body, swaggerJwkUpdateSetKey.body) &&
        Objects.equals(this.kid, swaggerJwkUpdateSetKey.kid) &&
        Objects.equals(this.set, swaggerJwkUpdateSetKey.set);
  }

  @Override
  public int hashCode() {
    return Objects.hash(body, kid, set);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SwaggerJwkUpdateSetKey {\n");
    
    sb.append("    body: ").append(toIndentedString(body)).append("\n");
    sb.append("    kid: ").append(toIndentedString(kid)).append("\n");
    sb.append("    set: ").append(toIndentedString(set)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
  
}

