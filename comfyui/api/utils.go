package api

import (
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
