package main

import (
	"./grabber"
	"./processing"
	"./handlers"
	"github.com/lazywei/go-opencv/opencv"
	"image"
	"net/http"
)

var frameSourceChanel chan *opencv.IplImage
var frameFinishedChanel chan *opencv.IplImage
var imgSourceChanel chan image.Image
var imgFinishedChanel chan image.Image

var router *http.ServeMux

func main() {
	frameSourceChanel = make(chan *opencv.IplImage, 10)
	frameFinishedChanel = make(chan *opencv.IplImage, 10)
	imgSourceChanel = make(chan image.Image, 5)
	imgFinishedChanel = make(chan image.Image, 5)

	sourceCapture := opencv.NewFileCapture("http://71.2.17.14/mjpg/video.mjpg")
	if sourceCapture == nil {
		panic("cannot open video streaming")
	}
	defer sourceCapture.Release()

	grabber.StartGrabb(sourceCapture, frameSourceChanel, frameFinishedChanel)
	processing.StartProcessingSource(frameSourceChanel, imgSourceChanel)
	processing.StartProcessingFinished(frameFinishedChanel, imgFinishedChanel)

	router = http.NewServeMux()

	router.HandleFunc("/", handlers.Homepage)

	router.HandleFunc("/source.mpng", func(w http.ResponseWriter, r *http.Request) {
		handlers.Source(w, r, imgSourceChanel)
	})

	router.HandleFunc("/finished.mpng", func(w http.ResponseWriter, r *http.Request) {
		handlers.Finished(w, r, imgFinishedChanel)
	})

	http.ListenAndServe(":8000", router)
}
