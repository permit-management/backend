package constants

type IDRequest struct {
	ID string `json:"id" binding:"required"`
}
