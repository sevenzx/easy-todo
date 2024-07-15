package resp

import "easytodo/model"

type UserResp struct {
	User      *model.User `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt string      `json:"expires_at"`
}
