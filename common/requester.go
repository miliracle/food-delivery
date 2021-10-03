package common

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
