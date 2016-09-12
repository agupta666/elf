package actions

import (
	"fmt"
	"net/http"

	"github.com/agupta666/hash/utils"
)

func writeFileName(name string, ext string, w http.ResponseWriter) {
	fileName := fmt.Sprintf("%s%s", name, ext)
	value := fmt.Sprintf("attachment; filename=%s", fileName)
	w.Header().Set("Content-Disposition", value)
}

func writeMimeType(ext string, w http.ResponseWriter) {
	mimeType := utils.TypeByExtension(ext)
	w.Header().Set("Content-Type", mimeType)
}
