package models

type Body struct {
	Ar1 []string `json:"ar1" binding:"required"`
	Ar2 []string `json:"ar2"  binding:"required"`
}
