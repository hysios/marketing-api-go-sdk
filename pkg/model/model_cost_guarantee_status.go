/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// CostGuaranteeStatus : 成本保障状态
type CostGuaranteeStatus string

// List of CostGuaranteeStatus
const (
	CostGuaranteeStatus_NONE       CostGuaranteeStatus = "COST_GUARANTEE_STATUS_NONE"
	CostGuaranteeStatus_EFFECTIVE  CostGuaranteeStatus = "COST_GUARANTEE_STATUS_EFFECTIVE"
	CostGuaranteeStatus_FAILED     CostGuaranteeStatus = "COST_GUARANTEE_STATUS_FAILED"
	CostGuaranteeStatus_FINISHED   CostGuaranteeStatus = "COST_GUARANTEE_STATUS_FINISHED"
	CostGuaranteeStatus_CONFIRMING CostGuaranteeStatus = "COST_GUARANTEE_STATUS_CONFIRMING"
	CostGuaranteeStatus_SUCCEEDED  CostGuaranteeStatus = "COST_GUARANTEE_STATUS_SUCCEEDED"
	CostGuaranteeStatus_FROZEN     CostGuaranteeStatus = "COST_GUARANTEE_STATUS_FROZEN"
)
