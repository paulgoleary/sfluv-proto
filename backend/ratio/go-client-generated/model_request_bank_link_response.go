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

type RequestBankLinkResponse struct {
	// The link token for the Plaid integration
	LinkToken string `json:"linkToken"`
}
