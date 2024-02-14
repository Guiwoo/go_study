package api

import (
	"fmt"
	"testing"
)

func TestUploadImage(t *testing.T) {
	err := UploadImage("/Users/guiwoopark/Desktop/personal/study/comfyui/assets/example2.png")
	fmt.Println(err)
}
