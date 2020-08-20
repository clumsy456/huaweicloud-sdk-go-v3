/*
 * Cbr
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AssociateVaultPolicyRequest struct {
	VaultId string          `json:"vault_id"`
	Body    *VaultAssociate `json:"body,omitempty"`
}

func (o AssociateVaultPolicyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AssociateVaultPolicyRequest", string(data)}, " ")
}