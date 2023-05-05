package main

import (
	"os"

	"encoding/json"
	"fmt"
	"log"
)

type Inline string

const (
	const_langEn Inline = "en"
	const_langEs Inline = "es"

	const_qrCode Inline = "qrCode"
	const_rgtSld Inline = "registerSalida"
	const_rgtVst Inline = "registerVisitiante"
	const_rgtDir Inline = "registerDirection"

	const_saludoEntrada Inline = "Se registrara la siguiente Entrada"
	const_rgtNom        Inline = "nombre"
	const_rgtCom        Inline = "company"
	const_rgtMot        Inline = "motvio"
	const_rgtCll        Inline = "calle"
	const_rgtExt        Inline = "exterior"

	const_rgt Inline = ""

	const_back   Inline = "back"
	const_ok     Inline = "ok"
	const_saludo Inline = "saludo"

	const_saludoSalida Inline = "Ingrese la siguiente informacion salida"
	const_sldFch       Inline = "sldFch"
	const_sldHra       Inline = "sldHra"

	const_saludoVisitante Inline = "Ingrese la siguiente informacion visitante"
	const_error           Inline = "error"
	const_solicitud       Inline = "solicitud"
	const_home            Inline = "home"
	const_entrada         Inline = "entrada"
	const_visitante       Inline = "visitante"
	const_salida          Inline = "salida"
	const_delay           int16  = 5
)

var (
	telegram_token = os.Getenv("telegram_bot_token")
	host_name      = os.Getenv("HOST_NAME")
	secure_key     = os.Getenv("SECURE_STRING")
	session_key    = os.Getenv("SESSION_STRING")
	//textos[idioma][]
	meses = map[Inline]([12]string){
		const_langEs: [12]string{
			"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre",
		},
		const_langEn: [12]string{
			"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december",
		},
	}
	textos = map[Inline](map[Inline]string){
		const_langEn: map[Inline](string){
			const_home:      "Hello and welcome to the service",
			const_visitante: string(const_saludoVisitante),
			const_salida:    string(const_saludoSalida),

			const_qrCode: "Generate QR Code",
			const_rgtSld: "Register Exit",
			const_rgtVst: "Register Visit",
			const_rgtDir: "Register Entrance",

			const_saludo: "hello",
			const_error:  "Unknow Error",

			const_entrada:   string(const_saludoEntrada),
			const_solicitud: "Please write the ",
			const_rgtNom:    "Name",
			const_rgtCom:    "Company Name",
			const_rgtMot:    "Motive of Visit",
			const_rgtCll:    "Street Name",
			const_rgtExt:    "Exterior Number",

			const_sldHra: "Register Exit hour",
			const_sldFch: "Register Exit date",
		},
		const_langEs: map[Inline](string){
			const_home:      "hola y bienvenido al servicio",
			const_visitante: string(const_saludoVisitante),
			const_salida:    string(const_saludoSalida),

			const_qrCode: "Generar QR",
			const_rgtSld: "Registrar Salida",
			const_rgtVst: "Registrar Visitante",
			const_rgtDir: "Registrar Entrada",

			const_saludo: "hola",
			const_error:  "Error Desconocido",

			const_entrada:   string(const_saludoEntrada),
			const_solicitud: "Por favor ingrese el ",
			const_rgtNom:    "Nombre",
			const_rgtCom:    "Compañia",
			const_rgtMot:    "Motivo",
			const_rgtCll:    "nombre de la Calle",
			const_rgtExt:    "Numero exterior",

			const_sldHra: "Registrar Hora de salida",
			const_sldFch: "Registrar Fecha de salida",
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

func toInline(data string) Inline {
	return Inline(data)
}

func inputMessage(msg Inline) string {
	return textos[idioma][const_solicitud] + textos[idioma][msg]
}

func printMes(mes int) string {
	return meses[idioma][mes-1]
}
