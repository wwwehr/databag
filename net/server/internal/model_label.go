/*
 * DataBag
 *
 * DataBag provides storage for decentralized identity based self-hosting apps. It is intended to support sharing of personal data and hosting group conversations. 
 *
 * API version: 0.0.1
 * Contact: roland.osborne@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package databag

type Label struct {

	LabelId string `json:"labelId"`

	LabelRevision int64 `json:"labelRevision"`

	Type_ string `json:"type"`

	Data string `json:"data"`

	Created int32 `json:"created"`

	Groups []string `json:"groups,omitempty"`
}
