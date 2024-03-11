package routes

import (
	"certificados_backend/files"
	"certificados_backend/mail"
	"certificados_backend/pdf"
	"certificados_backend/validations"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadEndpoint(c *gin.Context) {
	fileCSV, err := c.FormFile("csvfile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	records, err := files.ProcesarCSV(*fileCSV)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileImagen, err := c.FormFile("imagen")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validations.ValidarImagen(*fileImagen)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rutaImagen, err := files.GuardarImagenTemporal(*fileImagen)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Leer los valores del formulario para el nombre del taller y en mensaje del correo, en correo de envio y la contraseña
	email := c.PostForm("email")
	password := c.PostForm("password")
	title := c.PostForm("title")
	msj := c.PostForm("msj")

	err = validations.ValidarCredenciales(email, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resultados []map[string]string

	// Generar certificados para cada registro
	for _, record := range records {
		nombre, destinatario := record[0], record[1]

		resultado := map[string]string{"nombre": nombre, "email": destinatario, "estado": "", "error": ""}

		// Generar el PDF
		rutaCertificado, err := pdf.GenerarPDF(nombre, rutaImagen)

		if err != nil {
			log.Printf("Error al generar el certificado para %s: %v\n", nombre, err)
			resultado["estado"] = "Error en la generación del PDF"
			resultado["error"] = err.Error()
			resultados = append(resultados, resultado)
			continue
		}

		err = mail.EnviarCorreo(email, password, destinatario, rutaCertificado, title, msj, nombre)
		if err != nil {
			log.Printf("Error al enviar el correo a %s: %v\n", destinatario, err)
			resultado["estado"] = "Error en el envío del correo"
			resultado["error"] = err.Error()
			resultados = append(resultados, resultado)
			continue
		}

		resultado["estado"] = "Éxito"
		resultados = append(resultados, resultado)
	}

	// Limpiar la imagen temporal
	os.Remove(rutaImagen)

	c.JSON(http.StatusOK, gin.H{"resultados": resultados})
}
