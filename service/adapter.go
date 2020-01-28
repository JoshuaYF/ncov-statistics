package service

import "encoding/json"

func GetAllAreaFromDXY() Response {
	urlStr := "http://lab.isaaclin.cn/nCoV/api/area?latest=0"

	resp := Get(urlStr)

	data := Response{}

	json.Unmarshal([]byte(resp), &data)

	return data
}
