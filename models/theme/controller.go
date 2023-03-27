package theme

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	qa "github.com/TykTechnologies/portal-admin"
)

func Download(context *qa.Context) {
	var (
		result  interface{}
		err     error
		zipPath string
	)

	if result, err = context.FindOne(); err != nil {
		context.AddError(err)
		return
	}

	theme := result.(*Config)
	zipPath = theme.Path + ".zip"

	// Open the zip file
	zipFile, err := os.Open(zipPath)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer zipFile.Close()

	// Set the Content-Disposition header to specify the filename and instruct the browser to download the file
	filename := filepath.Base(zipPath)
	context.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// Copy the contents of the zip file to the response body
	_, err = io.Copy(context.Writer, zipFile)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Upload(context *qa.Context) {
	// Parse the multipart form containing the zip file upload
	err := context.Request.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the file from the form data
	file, header, err := context.Request.FormFile("file")
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a new zip file to write the uploaded file to
	zipPath := "themes/" + header.Filename
	zipFile, err := os.Create(zipPath)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer zipFile.Close()

	// Copy the uploaded file to the zip file
	_, err = io.Copy(zipFile, file)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	context.Writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(context.Writer, "File uploaded successfully")
}
