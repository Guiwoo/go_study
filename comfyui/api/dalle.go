package api

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func CallDalle(prompt, clientID string) (string, error) {
	APIKEY := getEnv("dalle")
	c := openai.NewClient(APIKEY)
	ctx := context.Background()

	// Example image as base64
	reqBase64 := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize512x512,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		Model:          openai.CreateImageModelDallE2,
		N:              1,
	}

	respBase64, err := c.CreateImage(ctx, reqBase64)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return "", err
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		return "", err
	}

	fileName := "upload/" + clientID + ".png"
	reader := bytes.NewReader(imgBytes)

	path, err := UploadAWS_S3(fileName, reader)
	if err != nil {
		return "", err
	}
	return path, nil
}
