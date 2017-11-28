package grabber

import (
	"github.com/lazywei/go-opencv/opencv"
	"log"
)

func StartGrabb(capture *opencv.Capture, chOne chan *opencv.IplImage, chTwo chan *opencv.IplImage) {
	go func() {
		for {
			if !capture.GrabFrame() {
				log.Println("Not grabe source capture")
			}

			frame := capture.RetrieveFrame(1)
			if frame == nil {
				log.Println("Frame is nil")
			}

			chOne <- frame.Clone()
			chTwo <- frame.Clone()
		}
	}()
}