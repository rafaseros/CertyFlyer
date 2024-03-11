package files

import (
	"certificados_backend/validations"
	"encoding/csv"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Validar y procesar el archivo CSV
func ProcesarCSV(fileCSV multipart.FileHeader) ([][]string, error) {
	// Procesar el CSV
	srcCSV, err := fileCSV.Open()
	if err != nil {
		return nil, err
	}
	defer srcCSV.Close()

	csvReader := csv.NewReader(srcCSV)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	err = validations.ValidarArchivoCSV(records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// Guardar la imagen de fondo temporalmente
func GuardarImagenTemporal(fileImagen multipart.FileHeader) (string, error) {
	// Guardar la imagen temporalmente para usarla en la generación de PDF
	srcImagen, err := fileImagen.Open()
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return "", err
	}
	defer srcImagen.Close()

	// Leer el contenido del archivo imagen en un slice de bytes
	contenidoImagen, err := io.ReadAll(srcImagen)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return "", err
	}

	rutaImagen := filepath.Join("temp", fileImagen.Filename)
	err = os.WriteFile(rutaImagen, contenidoImagen, os.ModePerm)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return "", err
	}
	return rutaImagen, nil
}
