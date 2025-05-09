/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意组件数据
type CreativeComponentsPermit struct {
	Name                    *string                   `json:"name,omitempty"`
	Desc                    *string                   `json:"desc,omitempty"`
	Valid                   *CreativeComponentValid   `json:"valid,omitempty"`
	ComponentType           ComponentType             `json:"component_type,omitempty"`
	ComponentSubTypeOptions *[]ComponentSubTypeOption `json:"component_sub_type_options,omitempty"`
}
