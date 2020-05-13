package utils

import (
	"fmt"
	//"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/png"
	"os"
	"os/exec"
)

func ParserPng(imagePath string) (image.Image, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}
func MergePngVertical(srcImagePaths []string, dstPath string) error {
	////纵向合并image
	width := 0
	heigth := 0
	if len(srcImagePaths) <= 0 {
		return nil
	}
	srcImages := make([]image.Image, len(srcImagePaths))
	for index, path := range srcImagePaths {
		img, err := ParserPng(path)
		if err != nil {
			return err
		}
		w := int(img.Bounds().Max.X - img.Bounds().Min.X)
		if w > width {
			width = w
		}
		heigth += int(img.Bounds().Max.Y - img.Bounds().Min.Y)
		srcImages[index] = img
		//fmt.Printf("width:%v,heigth:%v", width, heigth)
	}
	//fmt.Printf("len(srcImagePaths):%v", len(srcImagePaths))
	m := image.NewRGBA(image.Rect(0, 0, width, heigth))
	imgHeigth := 0
	for index, _ := range srcImages {
		//fmt.Printf("imgHeigth:%v", imgHeigth)
		//fmt.Printf("m.Bounds():%v", m.Bounds())
		draw.Draw(m, m.Bounds().Add(image.Point{0, imgHeigth}), srcImages[index], srcImages[index].Bounds().Min, draw.Src)
		imgHeigth += int(srcImages[index].Bounds().Max.Y - srcImages[index].Bounds().Min.Y)
	}
	//resize
	//newImg := resize.Resize(uint(width*3/4), uint(heigth*3/4), m, resize.Lanczos3)
	file, err := os.Create(dstPath)
	if err != nil {
		return err
	}

	if err := png.Encode(file, m); err != nil {
		file.Close()
		return err
	}
	file.Close()

	if err := exec.Command("/bin/sh", "-c", fmt.Sprintf(`optipng %s`, dstPath)).Run(); err != nil {
		return err
	}

	return nil
}
