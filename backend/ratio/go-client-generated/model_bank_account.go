/*
 * Ratio API
 *
 * API endpoints and models for using the Ratio service
 *
 * API version: 1.0.0
 * Contact: support@ratio.me
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BankAccount struct {
	// The unique identifier of the bank account
	Id string `json:"id,omitempty"`
	// The time the bank account connection was created
	CreateTime string `json:"createTime,omitempty"`
	// The time the bank account connection was last updated
	UpdateTime string `json:"updateTime,omitempty"`
	// The name of the bank account
	Name string `json:"name,omitempty"`
	// The account number mask
	Mask string `json:"mask,omitempty"`
	// The status of the bank account link to the user
	LinkStatus string `json:"linkStatus,omitempty"`
	// The status of the bank account user identity verification
	VerificationStatus string `json:"verificationStatus,omitempty"`
}
