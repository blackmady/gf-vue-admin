package response

import "server/app/model/apis"

type ApiListResponse struct {
	Apis []*apis.Entity `json:"apis"`
}

type ApiResponse struct {
	Api *apis.Entity `json:"api"`
}