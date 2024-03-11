package mail

import (
	"gopkg.in/gomail.v2"
)

// Configuración del servidor SMTP
const (
	SMTPHost = "smtp.gmail.com"
	SMTPPort = 587
)

// EnviarCorreo envía un correo electrónico al destinatario con el certificado adjunto.
// Recibe la dirección de correo del destinatario, la ruta al archivo PDF a adjuntar,
// el nombre del taller y una cita personalizada para incluir en el cuerpo del correo.
func EnviarCorreo(email, password, destinatario, rutaCertificado, title, msj, nombre string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", destinatario)
	m.SetHeader("Subject", "¡Tu certificado del "+title+" está aquí!")
	m.SetBody("text/html", `

    <p>¡Hola `+nombre+`!</p>

    <p>Espero que este mensaje te encuentre bien.</p>

    <p>Nos complace informarte que tu certificado del `+title+` ya está listo. Adjunto encontrarás tu certificado personalizado.</p>

    <p>`+msj+`</p>

    <p>Si tienes alguna pregunta o necesitas más información, no dudes en ponerte en contacto con nosotros.</p>

    <p>¡Felicidades de nuevo por completar `+title+`! Esperamos que este certificado te sea útil en tu trayectoria profesional.</p>

    <p>Saludos cordiales,</p>`)
	m.Attach(rutaCertificado)

	d := gomail.NewDialer(SMTPHost, SMTPPort, email, password)

	// Enviar el correo
	return d.DialAndSend(m)
}
