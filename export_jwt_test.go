package tools

import (
	"fmt"
	"testing"
	"time"
)

func TestJwtGenerate(t *testing.T) {
	res, err := JwtGenerate("123456", map[string]any{
		"A": 1,
		"t": time.Now().Unix(),
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestJwtParse(t *testing.T) {
	res, err := JwtParse[struct {
		A int   `json:"A"`
		T int64 `json:"t"`
	}]("123456", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2YWwiOnsiQSI6MSwidCI6MTY1ODQ4MjM1OX19.0G4APur9lgf86Svgdn_4j3fXxM4iWJvN8VLfCMFtlKo")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
