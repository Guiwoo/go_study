package api

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
)

const (
	prompt = "prompt"
	inputs = "inputs"

	ksampler      = "3"
	ksamplerSeed  = "seed"
	ksamplerCfg   = "cfg"
	ksamplerSteps = "steps"

	model    = "4"
	modelKey = "ckpt_name"

	positive    = "6"
	positiveKey = "text"

	negative    = "7"
	negativeKey = "text"

	output         = "9"
	fileNamePrefix = "filename_prefix"

	image          = "33"
	imageWidth     = "width"
	imageHeight    = "height"
	imageBatchSize = "batch_size"
)

func getEnv(env string) string {
	value, ok := os.LookupEnv(env)
	if ok == false {
		log.Fatalf("fail to get env %+v", env)
		return ""
	}
	return value
}

func SetClientID(value string, m map[string]interface{}) {
	m["ClientID"] = value
}

func SetStringJson(target, key, value string, m map[string]interface{}) {
	m[prompt].(map[string]interface{})[target].(map[string]interface{})[inputs].(map[string]interface{})[key] = value
}

func SetBigInt(target, key, value string, m map[string]interface{}) {
	intValue := new(big.Int)
	intValue, ok := intValue.SetString(value, 10)
	if !ok {
		log.Println("Failed to parse string as big.Int")
		return
	}
	m[prompt].(map[string]interface{})[target].(map[string]interface{})[inputs].(map[string]interface{})[key] = value
}

func SetIntJson(target, key, value string, m map[string]interface{}) {
	v, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("fail to conver seed to int %+v", err)
		return
	}
	m[prompt].(map[string]interface{})[target].(map[string]interface{})[inputs].(map[string]interface{})[key] = v
}

func SetFloatJson(target, key, value string, m map[string]interface{}) {
	v, err := strconv.ParseFloat(value, 32)
	if err != nil {
		log.Printf("fail to convert float %+v", err)
	}
	m[prompt].(map[string]interface{})[target].(map[string]interface{})[inputs].(map[string]interface{})[key] = v
}

type S3Manager struct {
	Bucket     string
	Region     string
	Cfg        *aws.Config
	client     *s3.Client
	uploader   *manager.Uploader
	downloader *manager.Downloader
}

var s3Manager *S3Manager

func loadS3Manager() error {
	accessKey := getEnv("aws_key")
	secretKey := getEnv("aws_secret_key")
	cred := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("ap-northeast-2"),
		config.WithLogConfigurationWarnings(true),
		config.WithCredentialsProvider(cred),
	)

	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	downloader := manager.NewDownloader(client)

	s3Manager = &S3Manager{
		Bucket:     "test-guiwoo",
		Region:     "ap-northeast-2",
		Cfg:        &cfg,
		client:     client,
		uploader:   uploader,
		downloader: downloader,
	}
	return nil
}

func UploadAWS_S3(filename string, data io.Reader) (string, error) {
	if s3Manager == nil {
		if err := loadS3Manager(); err != nil {
			return "", err
		}
	}

	multi := "multipart/form-data"

	input := &s3.PutObjectInput{
		Bucket:      &s3Manager.Bucket,
		Key:         &filename,
		Body:        data,
		ContentType: &multi,
	}

	output, err := s3Manager.uploader.Upload(context.TODO(), input)
	if err != nil {
		return "", err
	}
	return output.Location, nil
}
