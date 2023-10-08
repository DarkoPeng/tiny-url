package main

import (
	"fmt"
)

type Timestamp int64

// 短链接域名
const Base_Domain = "https://example.com"

type Link struct {
	SourceLink    string
	ShortLink     string
	GeneratedTime Timestamp
	VisitCount    int
	MaxVisitCount int
	Expired       Timestamp
}

var links []Link

// 生成一个唯一 id 用于短链接匹配
func GenerateUUID() {

}

func main() {
	// fmt.Println("Hello Tiny-URL")
	fmt.Printf(" short link %v\n", links)
}
