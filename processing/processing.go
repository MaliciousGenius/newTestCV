package processing

import (
	"github.com/lazywei/go-opencv/opencv"
	"image"
)

func StartProcessingSource(chFrame chan *opencv.IplImage, chImages chan image.Image) {
	go func() {
		for {
			frame := <- chFrame
			sourceImage := frame.ToImage()
			chImages <- sourceImage
		}
	}()
}

func StartProcessingFinished(frameToCannyChanel chan *opencv.IplImage, imgChanel chan image.Image) {
	go func() {
		for {
			frame := <- frameToCannyChanel
			w := frame.Width()
			h := frame.Height()

			cedge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 3)

			gray := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
			edge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
			defer gray.Release()
			defer edge.Release()

			opencv.CvtColor(frame, gray, opencv.CV_BGR2GRAY)

			opencv.Smooth(gray, edge, opencv.CV_BLUR, 3, 3, 0, 0)
			opencv.Not(gray, edge)

			opencv.Canny(gray, edge, 50, 200, 3)

			opencv.Zero(cedge)
			opencv.Copy(frame, cedge, edge)

			newSourceImage := cedge.ToImage()
			imgChanel <- newSourceImage
		}
	}()
}
