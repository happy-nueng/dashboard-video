package main

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"net/http"
	"time"

	"github.com/fogleman/gg"
)

func main() {
	http.HandleFunc("/stream", streamHandler)
	http.ListenAndServe(":8080", nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")

	for {
		timeStr := time.Now().Format("15:04:05")
		img := createImageWithText(timeStr)

		buf := new(bytes.Buffer)
		jpeg.Encode(buf, img, nil)

		w.Write([]byte("--frame\r\n"))
		w.Write([]byte("Content-Type: image/jpeg\r\n\r\n"))
		w.Write(buf.Bytes())
		w.Write([]byte("\r\n"))
		time.Sleep(1 * time.Second)
	}
}

func createImageWithText(text string) image.Image {
	const width, height = 640, 480
	dc := gg.NewContext(width, height)

	dc.SetColor(color.Black)
	dc.Clear()
	dc.SetColor(color.White)
	// dc.LoadFontFace("/font.ttf", 96)
	dc.DrawStringAnchored(text, width/2, height/2, 0.5, 0.5)

	return dc.Image()
}
