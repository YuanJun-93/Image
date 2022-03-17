package model

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamRegisterUser struct {
	UserName        string `json:"username" binding:"required"`
	PassWord        string `'json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成(1) 不投(0) 反对(-1)
}

type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`
	Page        int64  `json:"page" form:"page" example:"1"`
	Size        int64  `json:"size" form:"size" example:"10"`
	Order       string `json:"order" form:"order" example:"score"`
}

type ParamCouponInfo struct {
	CouponID string `json:"coupon_id" binding:"required"`
}

type ParamImagesUpload struct {
	ImageName string `json:"image_name" binding:"required"`
	ImageUrl  string `json:"image_url" binding:"required"`
	ImageCost int64  `json:"image_cost" binding:"required"`
}
