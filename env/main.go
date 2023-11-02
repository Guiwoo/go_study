package main

import (
	mCfg "env/config"
)

var config CustomConfig

type CustomConfig struct {
	Name  string `envconfig:"NAME" required:"true"`
	Email string `envconfig:EMAIL required:"true"`
	mCfg.Secret
}

//func init() {
//	cfg := CustomConfig{}
//	prefix := "G"
//	if err := envconfig.Process(prefix, &cfg); err != nil {
//		panic(err)
//	}
//	fmt.Printf("%+v", cfg)
//}

func main() {
}
