package constants

import (
	"os"

	"encoding/json"
	"fmt"
	"log"
)

type Inline string

const (
	Const_langEn Inline = "en"
	Const_langEs Inline = "es"

	Const_qrCode Inline = "qrCode"
	Const_rgtSld Inline = "registerSalida"
	Const_rgtVst Inline = "registerVisitiante"
	Const_rgtDir Inline = "registerDirection"

	Const_saludoEntrada Inline = "Se registrara la siguiente Entrada"
	Const_rgtNom        Inline = "nombre"
	Const_rgtCom        Inline = "company"
	Const_rgtMot        Inline = "motvio"
	Const_rgtCll        Inline = "calle"
	Const_rgtExt        Inline = "exterior"

	Const_rgt Inline = ""

	Const_back   Inline = "back"
	Const_ok     Inline = "ok"
	Const_saludo Inline = "saludo"

	Const_saludoSalida Inline = "Ingrese la siguiente informacion salida"
	Const_sldFch       Inline = "sldFch"
	Const_sldHra       Inline = "sldHra"

	Const_saludoVisitante Inline = "Ingrese la siguiente informacion visitante"
	Const_error           Inline = "error"
	Const_solicitud       Inline = "solicitud"
	Const_home            Inline = "home"
	Const_entrada         Inline = "entrada"
	Const_visitante       Inline = "visitante"
	Const_salida          Inline = "salida"
	Const_delay           int16  = 5
)

var (
	Idioma         Inline = "es"
	Telegram_token        = os.Getenv("telegram_bot_token")
	Host_name             = os.Getenv("HOST_NAME")
	Secure_key            = os.Getenv("SECURE_STRING")
	Session_key           = os.Getenv("SESSION_STRING")
	//textos[idioma][]
	meses = map[Inline]([12]string){
		Const_langEs: [12]string{
			"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre",
		},
		Const_langEn: [12]string{
			"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december",
		},
	}
	Textos = map[Inline](map[Inline]string){
		Const_langEn: map[Inline](string){
			Const_home:      "Hello and welcome to the service",
			Const_visitante: string(Const_saludoVisitante),
			Const_salida:    string(Const_saludoSalida),

			Const_qrCode: "Generate QR Code",
			Const_rgtSld: "Register Exit",
			Const_rgtVst: "Register Visit",
			Const_rgtDir: "Register Entrance",

			Const_saludo: "hello",
			Const_error:  "Unknow Error",

			Const_entrada:   string(Const_saludoEntrada),
			Const_solicitud: "Please write the ",
			Const_rgtNom:    "Name",
			Const_rgtCom:    "Company Name",
			Const_rgtMot:    "Motive of Visit",
			Const_rgtCll:    "Street Name",
			Const_rgtExt:    "Exterior Number",

			Const_sldHra: "Register Exit hour",
			Const_sldFch: "Register Exit date",
		},
		Const_langEs: map[Inline](string){
			Const_home:      "hola y bienvenido al servicio",
			Const_visitante: string(Const_saludoVisitante),
			Const_salida:    string(Const_saludoSalida),

			Const_qrCode: "Generar QR",
			Const_rgtSld: "Registrar Salida",
			Const_rgtVst: "Registrar Visitante",
			Const_rgtDir: "Registrar Entrada",

			Const_saludo: "hola",
			Const_error:  "Error Desconocido",

			Const_entrada:   string(Const_saludoEntrada),
			Const_solicitud: "Por favor ingrese el ",
			Const_rgtNom:    "Nombre",
			Const_rgtCom:    "Compañia",
			Const_rgtMot:    "Motivo",
			Const_rgtCll:    "nombre de la Calle",
			Const_rgtExt:    "Numero exterior",

			Const_sldHra: "Registrar Hora de salida",
			Const_sldFch: "Registrar Fecha de salida",
		},
	}
)

func Print(data interface{}) {
	//Converting to jsonn
	empJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(empJSON))
}

func ToInline(data string) Inline {
	return Inline(data)
}

func InputMessage(msg Inline) string {
	return Textos[Idioma][Const_saludo] + Textos[Idioma][msg]
}

func PrintMes(mes int) string {
	return meses[Idioma][mes-1]
}

func generateCredentials(){
    credenciales := struct {
        Make    string `json:"make"`
        Model   string `json:"model"`
        Mileage string `json:"mileage"`
    }{
        Make:    os.Getenv("nombre"),
        Model:   os.Getenv("modelo"),
        Mileage: os.Getenv("mileage"),
    }
    
    jsonCoded, err := json.Marshal(credenciales)
    fmt.Println(jsonCoded, err)
    
    err = ioutil.WriteFile("token.json", jsonCoded, 0644)
    if err != nil {
        log.Fatal(err)
    }
}
