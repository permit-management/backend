package constants

type IDRequest struct {
	ID uint `json:"id" binding:"required"`
}
