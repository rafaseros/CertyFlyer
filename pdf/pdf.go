package pdf

import (
	"log"
	"path/filepath"

	"github.com/signintech/gopdf"
)

// GenerarPDF crea un PDF para un certificado y lo guarda en el sistema de archivos.
// Recibe el nombre del participante, el nombre del taller, una cita y la ruta a la imagen de fondo.
// Retorna la ruta al archivo PDF generado o un error.
func GenerarPDF(nombre, rutaImagen string) (string, error) {

	anchoPagina := 812.0
	altoPagina := 602.0

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: anchoPagina, H: altoPagina}}) // Hoja carta horizontal W: 612, H: 792
	pdf.AddPage()

	// Añadir la imagen de fondo
	err := pdf.Image(rutaImagen, 0, 0, &gopdf.Rect{W: anchoPagina, H: altoPagina})
	if err != nil {
		return "", err
	}

	// Asegúrate de reemplazar "path/to/Roboto-Regular.ttf" con la ruta correcta a tu archivo de fuente TTF.
	err = pdf.AddTTFFont("Roboto-Regular", "./fonts/Roboto-Regular.ttf")
	if err != nil {
		log.Fatalf("Error al añadir fuente: %v", err)
	}

	err = pdf.SetFont("Roboto-Regular", "", 14)
	if err != nil {
		log.Fatalf("Error al configurar fuente: %v", err)
	}

	pdf.SetFontSize(40)

	// Ahora que la fuente ha sido cargada y configurada correctamente, puedes usarla para añadir o medir texto.
	anchoTexto, err := pdf.MeasureTextWidth(nombre)
	if err != nil {
		log.Fatalf("Error al medir ancho del texto: %v", err)
	}

	posX := (anchoPagina - anchoTexto) / 2

	// Nombre - Resaltar más

	pdf.SetX(posX)
	pdf.SetY(300) // Ajusta según sea necesario
	pdf.Text(nombre)

	// Guardar el PDF
	rutaArchivo := filepath.Join("out", nombre+".pdf")
	err = pdf.WritePdf(rutaArchivo)
	if err != nil {
		return "", err
	}

	return rutaArchivo, nil
}
