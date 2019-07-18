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
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

/**
 * Swaggeroauth2TokenParameters
 */
@javax.annotation.Generated(value = "io.swagger.codegen.languages.JavaClientCodegen", date = "2019-07-18T19:30:59.329+02:00")
public class Swaggeroauth2TokenParameters {
  @JsonProperty("client_id")
  private String clientId = null;

  @JsonProperty("code")
  private String code = null;

  @JsonProperty("grant_type")
  private String grantType = null;

  @JsonProperty("redirect_uri")
  private String redirectUri = null;

  public Swaggeroauth2TokenParameters clientId(String clientId) {
    this.clientId = clientId;
    return this;
  }

   /**
   * in: formData
   * @return clientId
  **/
  @ApiModelProperty(value = "in: formData")
  public String getClientId() {
    return clientId;
  }

  public void setClientId(String clientId) {
    this.clientId = clientId;
  }

  public Swaggeroauth2TokenParameters code(String code) {
    this.code = code;
    return this;
  }

   /**
   * in: formData
   * @return code
  **/
  @ApiModelProperty(value = "in: formData")
  public String getCode() {
    return code;
  }

  public void setCode(String code) {
    this.code = code;
  }

  public Swaggeroauth2TokenParameters grantType(String grantType) {
    this.grantType = grantType;
    return this;
  }

   /**
   * in: formData
   * @return grantType
  **/
  @ApiModelProperty(required = true, value = "in: formData")
  public String getGrantType() {
    return grantType;
  }

  public void setGrantType(String grantType) {
    this.grantType = grantType;
  }

  public Swaggeroauth2TokenParameters redirectUri(String redirectUri) {
    this.redirectUri = redirectUri;
    return this;
  }

   /**
   * in: formData
   * @return redirectUri
  **/
  @ApiModelProperty(value = "in: formData")
  public String getRedirectUri() {
    return redirectUri;
  }

  public void setRedirectUri(String redirectUri) {
    this.redirectUri = redirectUri;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Swaggeroauth2TokenParameters swaggeroauth2TokenParameters = (Swaggeroauth2TokenParameters) o;
    return Objects.equals(this.clientId, swaggeroauth2TokenParameters.clientId) &&
        Objects.equals(this.code, swaggeroauth2TokenParameters.code) &&
        Objects.equals(this.grantType, swaggeroauth2TokenParameters.grantType) &&
        Objects.equals(this.redirectUri, swaggeroauth2TokenParameters.redirectUri);
  }

  @Override
  public int hashCode() {
    return Objects.hash(clientId, code, grantType, redirectUri);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Swaggeroauth2TokenParameters {\n");
    
    sb.append("    clientId: ").append(toIndentedString(clientId)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    grantType: ").append(toIndentedString(grantType)).append("\n");
    sb.append("    redirectUri: ").append(toIndentedString(redirectUri)).append("\n");
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

