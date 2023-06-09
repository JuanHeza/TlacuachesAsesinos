package constants

import (
	"io/ioutil"
	"os"

	"encoding/json"
	"fmt"
	"log"
)

type Inline string
type Estatus int

const (
	Const_estatus_por_entrar Estatus = 1
	Const_estatus_por_salir  Estatus = 2
	Const_estatus_completo   Estatus = 3

	Mongo_collection       string = "Registros"
	Mongo_collection_users string = "Usuarios"
	Mongo_uri              string = "mongodb+srv://JuanHeza:1hCYw6lH9fF26Prs@evilpanda.cgorqpw.mongodb.net/?retryWrites=true&w=majority"
	Mongo_database         string = "TlacuachesAsesinos"

	Const_file_url      string = "https://docs.google.com/spreadsheets/d/1ytQV2XIoxR2TdiBiBIARRx5mg1v48kk7VF74eZ4LanQ/edit#gid=0"
	Const_date_template string = "02/01/2006"
	Const_time_template string = "15:04:05"

	Const_langEn Inline = "en"
	Const_langEs Inline = "es"

	Const_qrCode Inline = "qrCode"
	Const_rgtSld Inline = "registerSalida"
	Const_rgtVst Inline = "registerVisitiante"
	Const_rgtDir Inline = "registerDirection"

	Const_saludoRegistro Inline = "Se registrara la siguiente Entrada"
	Const_rgtNom         Inline = "nombre"
	Const_rgtCom         Inline = "company"
	Const_rgtMot         Inline = "motvio"
	Const_rgtCll         Inline = "calle"
	Const_rgtExt         Inline = "exterior"

	Const_rgt Inline = ""

	Const_back   Inline = "back"
	Const_ok     Inline = "ok"
	Const_okEnt  Inline = "okSal"
	Const_okSal  Inline = "okEnt"
	Const_saludo Inline = "saludo"

	Const_saludoSalida Inline = "Ingrese la siguiente informacion salida"
	Const_sldFch       Inline = "sldFch"
	Const_sldHra       Inline = "sldHra"

	Const_saludoEntrada Inline = "Ingrese la siguiente informacion entrada"
	Const_entFch        Inline = "entFch"
	Const_entHra        Inline = "entHra"
	Const_entFab        Inline = "entFab"
	Const_entCol        Inline = "entCol"
	Const_entObs        Inline = "entObs"
	Const_entFvh        Inline = "entFvh"
	Const_entFid        Inline = "entFid"
	Const_verFvh        Inline = "verFvh"
	Const_verFid        Inline = "verFid"

	Const_saludoVisitante Inline = "Ingrese la siguiente informacion visitante"
	Const_error           Inline = "error"
	Const_solicitud       Inline = "solicitud"
	Const_home            Inline = "home"
	Const_entrada         Inline = "entrada"
	Const_visitante       Inline = "visitante"
	Const_salida          Inline = "salida"

	Const_delay        int16 = 30
	Const_step_nuevo   int   = 0
	Const_step_entrada int   = 1
	Const_step_salida  int   = 2
	Const_pages        int   = 5
)

var (
	Mongo_CollectionNames        = []string{Mongo_collection, Mongo_collection_users}
	Idioma                Inline = "es"
	Telegram_token               = os.Getenv("telegram_bot_token")
	Host_name                    = os.Getenv("HOST_NAME")
	Secure_key                   = os.Getenv("SECURE_STRING")
	Session_key                  = os.Getenv("SESSION_STRING")
	Puerto                       = os.Getenv("PORT")
	//textos[idioma][]
	meses = map[Inline]([12]string){
		Const_langEs: [12]string{
			"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre",
		},
		Const_langEn: [12]string{
			"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december",
		},
	}
	Mensaje = map[Inline][]string{
		Const_langEn: {
			"Se han guardado los datos con el siguiente folio\n",
			"Se ha registrado la entrada del siguiente folio\n",
			"Se ha registrado la salida del siguiente folio\n",
		},
		Const_langEs: {
			"Se han guardado los datos con el siguiente folio\n",
			"Se ha registrado la entrada del siguiente folio\n",
			"Se ha registrado la salida del siguiente folio\n",
		},
	}
	Textos = map[Inline](map[Inline]string){
		Const_langEn: map[Inline](string){
			Const_home:      "Hello and welcome to the service",
			Const_visitante: string(Const_saludoVisitante),
			Const_salida:    string(Const_saludoSalida),
			Const_back:      "Cancel",

			Const_qrCode: "Generate QR Code",
			Const_rgtSld: "Register Exit",
			Const_rgtVst: "Register Entrance",
			Const_rgtDir: "Register Visit",

			Const_saludo: "hello",
			Const_error:  "Unknow Error",

			Const_entrada:   string(Const_saludoRegistro),
			Const_solicitud: "Please input: ",
			Const_rgtNom:    "Name",
			Const_rgtCom:    "Company Name",
			Const_rgtMot:    "Motive of Visit",
			Const_rgtCll:    "Street Name",
			Const_rgtExt:    "Exterior Number",

			Const_sldHra: "Exit hour",
			Const_sldFch: "Exit date",

			Const_entHra: "Entrance hour",
			Const_entFch: "Entrance date",
			Const_entFab: "Fabricante",
			Const_entCol: "Color",
			Const_entObs: "Observaciones",
			Const_entFvh: "Vehicle Picture",
			Const_entFid: "Identification Picture",
		},
		Const_langEs: map[Inline](string){
			Const_home:      "hola y bienvenido al servicio",
			Const_visitante: string(Const_saludoVisitante),
			Const_salida:    string(Const_saludoSalida),
			Const_back:      "Cancelar",

			Const_qrCode: "Generar QR",
			Const_rgtSld: "Registrar Salida",
			Const_rgtVst: "Registrar Entrada",
			Const_rgtDir: "Registrar Visitante",

			Const_saludo: "hola",
			Const_error:  "Error Desconocido",

			Const_entrada:   string(Const_saludoRegistro),
			Const_solicitud: "Por favor ingrese ",
			Const_rgtNom:    "Nombre",
			Const_rgtCom:    "Compañia",
			Const_rgtMot:    "Motivo",
			Const_rgtCll:    "Calle",
			Const_rgtExt:    "Numero exterior",

			Const_sldHra: "Hora de salida",
			Const_sldFch: "Fecha de salida",

			Const_entHra: "Hora de entrada",
			Const_entFch: "Fecha de entrada",
			Const_entFab: "Fabricante/Marca del vehiculo",
			Const_entCol: "Color del vehiculo",
			Const_entObs: "Observaciones",
			Const_entFvh: "Foto Vehiculo",
			Const_entFid: "Foto Identificacion",
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

func GetValue(value Inline) string {
	return Textos[Idioma][value]
}

func GetMessage(value int) string {
	return Mensaje[Idioma][value]
}

func GetQuerry(step string) int {
	switch ToInline(step) {
	case Const_ok:
		return Const_step_nuevo
	case Const_okEnt:
		return Const_step_entrada
	case Const_okSal:
		return Const_step_salida
	}
	return -1
}

func ToInline(data string) Inline {
	return Inline(data)
}

func InputMessage(msg Inline) string {
	return Textos[Idioma][Const_solicitud] + Textos[Idioma][msg]
}

func PrintMes(mes int) string {
	return meses[Idioma][mes-1]
}

func GenerateCredentials() {
	credenciales := struct {
		Tipo         string `json:"type"`
		ProjectoId   string `json:"project_id"`
		PrivateKeyId string `json:"private_key_id"`
		PrivateKey   string `json:"private_key"`
		ClientEmail  string `json:"client_email"`
		ClientId     string `json:"client_id"`
		AuthUri      string `json:"auth_uri"`
		TokenUri     string `json:"token_uri"`
		AuthProvider string `json:"auth_provider_x509_cert_url"`
		ClientCert   string `json:"client_x509_cert_url"`
	}{
		Tipo:         "service_account",
		ProjectoId:   os.Getenv("sheet_project_id"),
		PrivateKeyId: os.Getenv("sheet_private_key_id"),
		PrivateKey:   os.Getenv("sheet_private_key"),
		ClientEmail:  os.Getenv("sheet_client_email"),
		ClientId:     os.Getenv("sheet_client_id"),
		AuthUri:      "https://accounts.google.com/o/oauth2/auth",
		TokenUri:     "https://oauth2.googleapis.com/token",
		AuthProvider: "https://www.googleapis.com/oauth2/v1/certs",
		ClientCert:   os.Getenv("sheet_client_x509_cert_url"),
	}
	jsonCoded, err := json.Marshal(credenciales)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("./database/token.json", jsonCoded, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
