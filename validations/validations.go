package validations

import (
	"errors"
	"mime/multipart"
	"net/mail"
	"regexp"
	"strings"
)

// ValidarArchivoCSV verecordsCSV tenga el formato correcto
func ValidarArchivoCSV(records [][]string) error {

	for _, record := range records {
		if len(record) != 2 {
			return errors.New("el archivo CSV debe tener exactamente dos columnas")
		}

		// Validar el nombre
		if !ValidarNombre(record[0]) {
			return errors.New("el nombre contiene caracteres inválidos o números")
		}

		// Validar el correo electrónico
		if !ValidarEmail(record[1]) {
			return errors.New("uno de los correos electrónicos no tiene un formato válido")
		}
	}

	return nil
}

// ValidarImagen verifica que el archivo sea de tipo imagen
func ValidarImagen(fileImagen multipart.FileHeader) error {
	contentType := fileImagen.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		return errors.New("el archivo proporcionado no es una imagen")
	}

	return nil
}

// ValidarCredenciales verifica que el correo electrónico y la contraseña sean válidos
func ValidarCredenciales(email, password string) error {
	if !ValidarEmail(email) {
		return errors.New("el correo electrónico no tiene un formato válido")
	}

	if len(password) < 6 {
		return errors.New("la contraseña debe tener más de 6 caracteres")
	}

	return nil
}

// ValidarEmail verifica que la cadena tenga formato de correo electrónico válido
func ValidarEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// validarNombre verifica que el nombre no contenga números ni caracteres especiales
func ValidarNombre(nombre string) bool {
	re := regexp.MustCompile(`^[a-zA-ZáéíóúÁÉÍÓÚñÑ\s]+$`)
	return re.MatchString(nombre)
}
