package model

import "github.com/minnowo/astoryofand/util"

type Use struct {
	Email    string `json:"email" form:"email"`
	FullName string `json:"fullname" form:"fullname"`
	Comment  string `json:"comment" form:"comment"`
}

func (o *Use) CheckValid() bool {

	if util.IsEmptyOrWhitespace(o.Email) ||
		util.IsEmptyOrWhitespace(o.FullName) ||
		util.IsEmptyOrWhitespace(o.Comment) {
		return false
	}

	return true
}
