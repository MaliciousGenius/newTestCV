package handlers

import (
	"bytes"
	"image/png"
	"image"
	"net/http"
	"fmt"
	"io"
	"log"
)

func Homepage (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
	<html>
	<head>
		<title>GoCV</title>
	</head>
	<body style="background:#000000">
		<table width="100%" height="100%">
			<tr>
				<td align=center>
					<img style="image-rendering:-moz-crisp-edges; image-rendering:pixelated" src="source.mpng" />
				</td>
				<td align=center>
					<img style="image-rendering:-moz-crisp-edges; image-rendering:pixelated" src="finished.mpng" />
				</td>
			</tr>
		</table>
	</body>
	</html>
	`))
}

func Source (w http.ResponseWriter, r *http.Request, ch chan image.Image)  {
	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=--pngboundary")
	w.WriteHeader(http.StatusOK)

	for {
		var buf bytes.Buffer
		img := <- ch
		enc := png.Encoder{CompressionLevel: png.BestSpeed}
		enc.Encode(&buf, img)
		png := buf.Bytes()
		w.Write([]byte("--pngboundary"))
		w.Write([]byte("Content-Type: image/png\n"))
		w.Write([]byte(fmt.Sprintf("Content-Length: %d\n\n", len(png))))
		if _, err := io.Copy(w, bytes.NewReader(png)); err != nil {
			return
		}
		log.Println("<--- frame")
	}
}

func Finished (w http.ResponseWriter, r *http.Request, ch chan image.Image)  {
	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=--pngboundary")
	w.WriteHeader(http.StatusOK)

	for {
		var buf bytes.Buffer
		img := <- ch
		enc := png.Encoder{CompressionLevel: png.BestSpeed}
		enc.Encode(&buf, img)
		png := buf.Bytes()
		w.Write([]byte("--pngboundary"))
		w.Write([]byte("Content-Type: image/png\n"))
		w.Write([]byte(fmt.Sprintf("Content-Length: %d\n\n", len(png))))
		if _, err := io.Copy(w, bytes.NewReader(png)); err != nil {
			return
		}
		log.Println("<--- frame")
	}
}


	//router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	//	res.Header().Set("Content-Type", "text/html")
	//	res.Write([]byte(`
	//		<html>
	//		<head>
	//			<title>Bitvis</title>
	//		</head>
	//		<body style="margin:0; background:#000">
	//				<table width="100%" height="100%">
	//						<tr>
	//								<td valign=middle align=center>
	//										<a href="frame.png" target="_blank">
	//											<img style="image-rendering:-moz-crisp-edges; image-rendering:pixelated" src="stream.mpng" />
	//										</a>
	//								</td>
	//						</tr>
	//				</table>
	//		</body>
	//		</html>
	//	`))
	//})