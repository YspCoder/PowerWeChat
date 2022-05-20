package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type ResponseDeviceMessage struct {
	*response.ResponseOfficialAccount

	Ret     int    `json:"ret"`
	RetInfo string `json:"ret_info"`
}
