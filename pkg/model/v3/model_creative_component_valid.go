/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意组件字段结构
type CreativeComponentValid struct {
	Required  *bool  `json:"required,omitempty"`
	MinOccurs *int64 `json:"min_occurs,omitempty"`
	MaxOccurs *int64 `json:"max_occurs,omitempty"`
	HasDepend *bool  `json:"has_depend,omitempty"`
}
