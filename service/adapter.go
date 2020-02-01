package service

import "encoding/json"

func GetAllAreaFromDXY() Response {
	//urlStr := "https://lab.isaaclin.cn/nCoV/api/area?latest=1"
	urlStr := "https://lab.ahusmart.com/nCoV/api/area?latest=1"

	resp := Get(urlStr)

	data := Response{}

	json.Unmarshal([]byte(resp), &data)

	return data
}

func GetHistoryAreaFromDXY() Response {
	urlStr := "https://lab.ahusmart.com/nCoV/api/area?latest=1"

	resp := Get(urlStr)

	data := Response{}

	json.Unmarshal([]byte(resp), &data)

	return data
}
