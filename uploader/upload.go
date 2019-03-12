package uploader

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error in reading from request file: ", err)
		return
	}
	defer file.Close()
	io.Copy(&buf, file)
	//fmt.Println("File content: ", buf.String())
	fmt.Println("File header: ", header.Header)
	return
}
