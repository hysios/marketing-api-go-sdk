/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type WalletCreateAddRequest struct {
	AccountId              *int64                  `json:"account_id,omitempty"`
	MdmId                  *int64                  `json:"mdm_id,omitempty"`
	WalletName             *string                 `json:"wallet_name,omitempty"`
	TagList                *[]string               `json:"tag_list,omitempty"`
	ContactInfoList        *[]ContactInfoStruct    `json:"contact_info_list,omitempty"`
	ContactNotifyCondition *ContactNotifyCondition `json:"contact_notify_condition,omitempty"`
}
