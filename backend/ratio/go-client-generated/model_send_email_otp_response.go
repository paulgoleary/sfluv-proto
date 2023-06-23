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

type SendEmailOtpResponse struct {
	// The ID of the email address
	EmailId string `json:"emailId"`
	// The masked email address
	EmailMask string `json:"emailMask"`
}
