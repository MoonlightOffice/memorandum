package main

import (
	"encoding/json"
	"fmt"
)

type SSEData interface {
	Marshal() []byte
}

func Format(data SSEData, keepAlive ...bool) string {
	return fmt.Sprintf("data: %s\n", string(data.Marshal()))
}

func FormatKeepAlive() string {
	return ": keep-alive\n"
}

type UserData struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func (u UserData) Marshal() []byte {
	b, _ := json.Marshal(u)
	return b
}
