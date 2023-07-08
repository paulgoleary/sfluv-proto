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

type SecondFactorAuthResponse struct {
	// The authenticated Bearer token
	SessionJwt string `json:"sessionJwt"`
	User *User `json:"user,omitempty"`
}