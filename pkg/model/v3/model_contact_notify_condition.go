/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 监控预警信息
type ContactNotifyCondition struct {
	StatusDesc    *string                `json:"status_desc,omitempty"`
	Status        *int64                 `json:"status,omitempty"`
	ConditionList *[]ConditionInfoStruct `json:"condition_list,omitempty"`
}
