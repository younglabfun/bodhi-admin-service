package model

import "gorm.io/gorm"

var ErrNotFound = gorm.ErrRecordNotFound

type PageReq struct {
	Page  int64  `json:"page"`
	Size  int64  `json:"size"`
	Sort  string `json:"sort,optional"`
	Order string `json:"order,optional"`
	Field string `json:"field,optional"`
	Value string `json:"value,optional"`
}
