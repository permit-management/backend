package errcode

import "net/http"

var (
	// Success         = NewError(0, "Success", http.StatusOK)
	ServerError     = NewError(100000, "Server Error", http.StatusInternalServerError)
	BadRequest      = NewError(100001, "Bad Request", http.StatusBadRequest)
	TooManyRequests = NewError(100002, "Too Many Requests", http.StatusTooManyRequests)

	// 2001xx for authentication related error, TODO specify more data
	InvalidToken       = NewError(200100, "Invalid Token", http.StatusUnauthorized)
	UnauthorizedToken  = NewError(200101, "Token Not Provided", http.StatusUnauthorized)
	UnauthorizedAccess = NewError(200102, "Access Not Allowed", http.StatusForbidden)
	InvalidUser        = NewError(200103, "Invalid User", http.StatusUnauthorized)

	// 2002xx for data related error, TODO specify more data
	InvalidRequest      = NewError(200200, "Invalid Request", http.StatusBadRequest)
	ErrorUploadFileFail = NewError(200202, "File Upload Failed", http.StatusBadRequest)
	FileNotFound        = NewError(200203, "File Not Found", http.StatusNotFound)
	InvalidFileFormat   = NewError(200204, "Invalid File Format", http.StatusBadRequest)
	InvalidFileSize     = NewError(200205, "Invalid File Size", http.StatusBadRequest)
)
