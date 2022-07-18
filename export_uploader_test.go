package tools

import (
	"fmt"
	"os"
	"testing"
)

func TestFileUploadIO(t *testing.T) {

	f, err := os.Open("C:\\Users\\Administrator\\Desktop\\图片s\\img\\f.png")
	if err != nil {
		t.Fatal(err)
	}
	stat, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	id, err := FileUploadIO(f, stat.Name())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(id)
}

func TestFileGet(t *testing.T) {
	id := "722d7648047-00257f-2242ef-22b92e-007d2d8c1c743da.png"
	fmt.Println(FileInfo(id))
}

func TestName(t *testing.T) {
	v := fmt.Sprintf("%08d", 1231)
	fmt.Println(v, len([]byte(v)))
}
