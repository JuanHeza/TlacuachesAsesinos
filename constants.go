package main

import (
	"os"

	"encoding/json"
	"fmt"
	"log"
)

type Inline string
const (
	const_qrCode Inline = "qrCode"
	const_langEn Inline = "lenguageEn"
	const_langEs Inline = "lenguageEs"
	const_rgtSld Inline = "registerSalida"
	const_rgtVst Inline = "registerVisitiante"
	const_rgtDir Inline = "registerDirection"
)

var (
	telegram_token = os.Getenv("telegram_bot_token")
	host_name      = os.Getenv("HOST_NAME")
	secure_key     = os.Getenv("SECURE_STRING")
	textos         = map[string](map[string]string){
		"en": map[string](string){
			"saludo": "hello",
			"error":  "Unknow Error",
		},
		"es": map[string](string){
			"saludo": "hola",
			"error":  "Error Desconocido",
		},
	}
)

func print(data interface{}) {
	//Converting to jsonn
	empJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(empJSON))
}

func toInline(data string)Inline{
	return Inline(data)
}