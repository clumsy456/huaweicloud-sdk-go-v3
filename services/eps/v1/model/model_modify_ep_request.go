/*
    * EPS
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// Request Object
type ModifyEpRequest struct {
	EnterpriseProjectId string `json:"enterprise_project_id"`
	Body *EnterpriseProject `json:"body,omitempty"`
}