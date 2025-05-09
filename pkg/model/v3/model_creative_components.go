/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意组件
type CreativeComponents struct {
	Title              *[]TitleComponent              `json:"title,omitempty"`
	Description        *[]DescriptionComponent        `json:"description,omitempty"`
	Image              *[]ImageComponent              `json:"image,omitempty"`
	ImageList          *[]ImageListComponent          `json:"image_list,omitempty"`
	Video              *[]VideoComponent              `json:"video,omitempty"`
	Brand              *[]BrandComponent              `json:"brand,omitempty"`
	Consult            *[]ConsultComponent            `json:"consult,omitempty"`
	Phone              *[]PhoneComponent              `json:"phone,omitempty"`
	Form               *[]FormComponent               `json:"form,omitempty"`
	ActionButton       *[]ActionButtonComponent       `json:"action_button,omitempty"`
	ChosenButton       *[]ChosenButtonComponent       `json:"chosen_button,omitempty"`
	Label              *[]LabelComponent              `json:"label,omitempty"`
	ShowData           *[]ShowDataComponent           `json:"show_data,omitempty"`
	MarketingPendant   *[]MarketingPendantComponent   `json:"marketing_pendant,omitempty"`
	AppGiftPackCode    *[]AppGiftPackCodeComponent    `json:"app_gift_pack_code,omitempty"`
	ShopImage          *[]ShopImageComponent          `json:"shop_image,omitempty"`
	CountDown          *[]CountDownComponent          `json:"count_down,omitempty"`
	Barrage            *[]BarrageComponent            `json:"barrage,omitempty"`
	FloatingZone       *[]FloatingZoneComponent       `json:"floating_zone,omitempty"`
	TextLink           *[]TextLinkComponent           `json:"text_link,omitempty"`
	EndPage            *[]EndPageComponent            `json:"end_page,omitempty"`
	LivingDesc         *[]LivingDescComponent         `json:"living_desc,omitempty"`
	WechatChannels     *[]WechatChannelsComponent     `json:"wechat_channels,omitempty"`
	ShortVideo         *[]ShortVideoComponent         `json:"short_video,omitempty"`
	ElementStory       *[]ElementStoryComponent       `json:"element_story,omitempty"`
	WxgamePlayablePage *[]WxgamePlayablePageComponent `json:"wxgame_playable_page,omitempty"`
	MainJumpInfo       *[]JumpinfoComponent           `json:"main_jump_info,omitempty"`
	AppPromotionVideo  *[]AppPromotionVideoComponent  `json:"app_promotion_video,omitempty"`
	VideoShowcase      *[]VideoShowcaseComponent      `json:"video_showcase,omitempty"`
	ImageShowcase      *[]ImageShowcaseComponent      `json:"image_showcase,omitempty"`
	SocialSkill        *[]SocialSkillComponent        `json:"social_skill,omitempty"`
	MiniCardLink       *[]MiniCardLinkComponent       `json:"mini_card_link,omitempty"`
	FloatingZoneList   *[]FloatingZoneListComponent   `json:"floating_zone_list,omitempty"`
}
