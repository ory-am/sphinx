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
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * HealthNotReadyStatus
 */
@javax.annotation.Generated(value = "io.swagger.codegen.languages.JavaClientCodegen", date = "2019-06-11T14:10:17.875+02:00")
public class HealthNotReadyStatus {
  @JsonProperty("errors")
  private Map<String, String> errors = null;

  public HealthNotReadyStatus errors(Map<String, String> errors) {
    this.errors = errors;
    return this;
  }

  public HealthNotReadyStatus putErrorsItem(String key, String errorsItem) {
    if (this.errors == null) {
      this.errors = new HashMap<String, String>();
    }
    this.errors.put(key, errorsItem);
    return this;
  }

   /**
   * Errors contains a list of errors that caused the not ready status.
   * @return errors
  **/
  @ApiModelProperty(value = "Errors contains a list of errors that caused the not ready status.")
  public Map<String, String> getErrors() {
    return errors;
  }

  public void setErrors(Map<String, String> errors) {
    this.errors = errors;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    HealthNotReadyStatus healthNotReadyStatus = (HealthNotReadyStatus) o;
    return Objects.equals(this.errors, healthNotReadyStatus.errors);
  }

  @Override
  public int hashCode() {
    return Objects.hash(errors);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HealthNotReadyStatus {\n");
    
    sb.append("    errors: ").append(toIndentedString(errors)).append("\n");
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

