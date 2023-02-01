package dto

type ReqByIDs struct {
	IDS []uint64 `json:"ids" form:"ids" binding:"required,gte=1"`
}
