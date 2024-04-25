package utils

import (
	"encoding/json"
)

type CodeRep struct {
	Code string `json:"code"`
}

func ParseCode(code string) string {
	c := CodeRep{
		Code: code,
	}

	jsonDatam, _ := json.Marshal(c)
	return string(jsonDatam)

}

func ParsePort(port string) string {
	return ":" + port
}
