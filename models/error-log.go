package models

type ErrorLog struct {
	StorageBase
	Caller   string `json:"caller"`
	ErrorMsg string `json:"error_msg"`
}
