# TlacuachesAsesinos

https://media.wizards.com/2018/dnd/downloads/DnD_BasicRules_2018.pdf

https://wepik.com/es/inteligencia-artificial

https://www.yunzii.com/products/yunzii-actto-b703-retro-typewriter-mechanical-keyboard?variant=44359017988338

https://www.amazon.com.mx/EPOMAKER-Intercambiable-Bluetooth-retroiluminaci%C3%B3n-interruptor/dp/B0BS18JCZ6/ref=d_pd_day0_sccl_3_3/134-5768482-1243942?pd_rd_w=Iyh1v&content-id=amzn1.sym.1c8f9346-a87e-48c7-994a-82c93e96af91&pf_rd_p=1c8f9346-a87e-48c7-994a-82c93e96af91&pf_rd_r=FANY55HT746VF90TRE1F&pd_rd_wg=X4fiM&pd_rd_r=fbd8a621-d364-4ace-bdc6-a412abbd7316&pd_rd_i=B0BS18KHNP&th=1

https://www.amazon.com.mx/dp/B09WJN83ZS/ref=sspa_dk_detail_2?pf_rd_p=f4b065a7-c255-43f0-a2c0-b69847ce06d2&pf_rd_r=Z1GAWCRYDNFEGRM0PXZJ&pd_rd_wg=gw64w&pd_rd_w=p8UeR&content-id=amzn1.sym.f4b065a7-c255-43f0-a2c0-b69847ce06d2&pd_rd_r=8c7c7dc8-6c0c-4ef4-b173-bddd388b57fd&s=electronics&sp_csd=d2lkZ2V0TmFtZT1zcF9kZXRhaWw&spLa=ZW5jcnlwdGVkUXVhbGlmaWVyPUEzOVBaNEJGNTkyMzhTJmVuY3J5cHRlZElkPUEwOTE5NDkwMjY3SUVURUhNRlhLUyZlbmNyeXB0ZWRBZElkPUEwMDA1OTY2M1I0OFg1OFlDWFg0NSZ3aWRnZXROYW1lPXNwX2RldGFpbCZhY3Rpb249Y2xpY2tSZWRpcmVjdCZkb05vdExvZ0NsaWNrPXRydWU&th=1

https://www.amazon.com.mx/dp/B09Q5MCFW3/ref=sspa_dk_detail_6?psc=1&pf_rd_p=f4b065a7-c255-43f0-a2c0-b69847ce06d2&pf_rd_r=Z1GAWCRYDNFEGRM0PXZJ&pd_rd_wg=gw64w&pd_rd_w=p8UeR&content-id=amzn1.sym.f4b065a7-c255-43f0-a2c0-b69847ce06d2&pd_rd_r=8c7c7dc8-6c0c-4ef4-b173-bddd388b57fd&s=electronics&sp_csd=d2lkZ2V0TmFtZT1zcF9kZXRhaWw&smid=A34XVXTBCHMT4Z&spLa=ZW5jcnlwdGVkUXVhbGlmaWVyPUFBSjdLOTJQUFlXVEMmZW5jcnlwdGVkSWQ9QTA1MzU3MTAzQTZVRlVFMTNQNjJDJmVuY3J5cHRlZEFkSWQ9QTA2MjAyNTMxM0U3RUkwUE1JQTJWJndpZGdldE5hbWU9c3BfZGV0YWlsJmFjdGlvbj1jbGlja1JlZGlyZWN0JmRvTm90TG9nQ2xpY2s9dHJ1ZQ==


// https://github.com/meneer-code/Connect-Telegram-Bot-to-Google-Sheets-ChatGPT-OpenAI
// https://developers.google.com/sheets/api/quickstart/go
// https://developers.google.com/workspace/guides/create-project?authuser=1



Dudas

Formato del folio
    - El formato es estrictamente necesario asi?
        Fecha Juliana XXXX
        Clave del domicilio en 3 digitos (único) XXX
        Clave de horario en 3 digitos XXX
    
    - La fecha juliana son 7 digitos
    - Los domicilios tendriasn que estar listados para poderse generar una clave homogenea
        - Un usuario puede escribir "Prolongacion ildefonso fuentes #531" mientras que otro "Prol. ildefonso fuentes 531"
        - Se usaria el numero de domicilio ya que ese seria una entrada igual en cada usuario
    - Como se van a seccionar los horarios? por 3 turnos "mañana, tarde y noche"? por horas "1 ,2, 3, ..., 23"

    - Actualmente se esta usando el formato unix para el folio, un numero de 10 digitos que que codifican la fecha y dia 

El QR
    - El codigo QR servira como preRegistro? el usuario llena el formulario y le genera un QR que al ser leido generara los datos de entrada 
        o sera un pase directo donde el portador del QR pasara sin necesidad de llenar el formulario de registro
    - Habra necesidad de QR de salida? que al ser leido llene automaticamente los datos de salida 

Usuarios
    - Por seguridad del bot se tiene contemplado implementar 3 niveles "residente - caseta - administrador"
    - Al iniciar el bot, se revisara si existe en la base de datos, de no existir se le respondera con "Su solicitud para usar el servicio 
        esta siendo procesada, espere a que un administrador la responda" y no podra hacer nada
    - Al administrador le llegara un mensaje "El Usuario @ quiere hacer uso del sistema" con 4 opciones "dar permiso de residente", 
        "dar permiso de vigilante", "dar permiso de administrador" y "rechazar", el usuario recibira una actualizacion de su estatus, los 
        primeros 3 le dara acceso a ciertas areas del bot, la ultima no le permitira hacer nada y futuras solicitudes seran omitidas
    - Administrador contara con la opcion de revocar permisos a usuarios asi como tambien cambiar los tipo de usuario "vigilante A pasa a ser 
        administrador"

Envio masivo 
    - Tener una lista de usuarios permitiria hacer un envio masivo de mensajes desde un administrador ya sea hacia todos los usuarios o solo
        a los usuarios de un nivel especifico "enviar solo a residentes"

Interfaz Web
    - Se descarto el excel ya que buscar un folio seria una tarea exhaustiva segun vayan creciendo los registros, una base de datos haria la
        busqueda mas agil
    - Para sustituir la representacion visual de los datos en el excel se propone hacer una interfaz web con un inicio de sesion para que solo 
        administradores puedan acceder y ver en una tabla todos los datos del registro y las funciones de aceptar/rechazar usuarios y envio     
        masivo
    - El servidor web es necesario, independientemente de la interfaz ya que se necesita para la generacion, procesamiento y validacion del 
        codigo QR, si se requiere que el software tenga conexion con dispositivos fisicos como un porton tendria que revisarse como seria la 
            conexion y si se requiere el desarrollo de un dispositivo que sirva de interfaz entre el servidor y el hardware

Duda de juan para Iza
    - Se puede hacer el pago en 2 partes, una a mitad del desarrollo y una al entregar
    - Esto lo digo por lo del juego, si lo quiero tendria hasta el 20 para comprarlo con las figuras de vinil XD