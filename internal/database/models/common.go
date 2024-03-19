package models

type BasicValidator interface {
	CheckValid() bool
}

type UserDataType int32

const (
	OrderType UserDataType = 1 << iota
	UsecaseType
)

type UserData struct {
	Type UserDataType `json:"type"`
}
