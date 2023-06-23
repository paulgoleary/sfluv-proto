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

type UserMask struct {
	// The unique identifier of the user
	Id string `json:"id,omitempty"`
	// The time the user was created
	CreateTime string `json:"createTime,omitempty"`
	// The time the user was last updated
	UpdateTime string `json:"updateTime,omitempty"`
	// The last 4 digits of the user's phone number
	PhoneMask string `json:"phoneMask,omitempty"`
	// The user's masked email address
	EmailMask string `json:"emailMask,omitempty"`
}
