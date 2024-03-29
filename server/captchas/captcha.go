package captchas

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
)

func Run() (*CutoutRet, error) {
	// 加载图片到缓存
	LoadBackgroudImages("./examples/b")
	// 加载背景图到缓存
	LoadBlockImages("./examples/block")
	bgImage, err := randBackgroudImage()
	if err != nil {
		return nil, err
	}
	bkImage, err := randBlockImage()
	if err != nil {
		return nil, err
	}
	return run(bgImage, bkImage)
}

func run(bgImage, bkImage *ImageBuf) (*CutoutRet, error) {
	ret := new(CutoutRet)
	bgWidth := bgImage.getWidth()
	bgHeight := bgImage.getHeight()
	bkWidth := bkImage.getWidth()
	bkHeight := bkImage.getHeight()
	ret.Point = randPoint(bgWidth, bgHeight, bkWidth)
	newBkImage := &ImageBuf{
		w: bkWidth,
		h: bkHeight,
		i: image.NewNRGBA(image.Rect(0, 0, bkWidth, bkHeight)),
	}
	x := ret.Point.X

	// 抠图
	cutOut(bgImage, bkImage, newBkImage, x)

	var err error
	ret.BackgroudImg, err = img2Base64(bgImage)
	if err != nil {
		return nil, err
	}
	ret.BlockImg, err = img2Base64(newBkImage)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// randPoint 随机生成抠图位置
func randPoint(bgWidth, bgHeight, bkWidth int) *Point {
	wDiff := bgWidth - bkWidth
	//hDiff := bgHeight - bkWidth
	var x, y int
	if wDiff <= 0 {
		x = 5
	} else {
		x = r.Intn(wDiff-100) + 100
	}
	y = 0
	return &Point{x, y}
}

// cutOut 抠图
func cutOut(bgImage, bkImage, newBkImage *ImageBuf, x int) {
	var values [9]color.RGBA64
	bkWidth := bkImage.getWidth()
	bkHeight := bkImage.getHeight()
	for i := 0; i < bkWidth; i++ {
		for j := 0; j < bkHeight; j++ {
			pixel := bkImage.getRGBA(i, j)
			// 滑块图片非透明像素点，从背景图偏移x 像素拷贝到新图层
			if pixel.A > 0 {
				newBkImage.setRGBA(i, j, bgImage.getRGBA(x+i, j))
				readNeighborPixel(bgImage, x+i, j, &values)
				bgImage.setRGBA(x+i, j, gaussianBlur(&values))
			}
			if i == (bkWidth-1) || j == (bkHeight-1) {
				continue
			}
			rightPixel := bkImage.getRGBA(i+1, j)
			bottomPixel := bkImage.getRGBA(i, j+1)
			// 用白色给底图和新图层描边
			if (pixel.A > 0 && rightPixel.A == 0) ||
				(pixel.A == 0 && rightPixel.A > 0) ||
				(pixel.A > 0 && bottomPixel.A == 0) ||
				(pixel.A == 0 && bottomPixel.A > 0) {
				white := color.White
				newBkImage.setRGBA(i, j, white)
				bgImage.setRGBA(x+i, j, white)
			}
		}
	}
}

//readNeighborPixel 读取邻近9个点像素，后面最类似高斯模糊计算
//（并非严格的高斯模糊，高斯模糊算法效率太低，本例不需要严格的高斯模糊算法）
// |2|3|4|
// |5|1|6|
// |7|8|9|
// 中心点为1
func readNeighborPixel(img *ImageBuf, x, y int, pixels *[9]color.RGBA64) {
	xStart := x - 1
	yStart := y - 1
	current := 0
	for i := xStart; i < 3+xStart; i++ {
		for j := yStart; j < 3+yStart; j++ {
			tx := i
			if tx < 0 {
				tx = -tx
			} else if tx >= img.getWidth() {
				tx = x
			}
			ty := j
			if ty < 0 {
				ty = -ty
			} else if ty >= img.getHeight() {
				ty = y
			}
			pixels[current] = img.getRGBA(tx, ty)
			current++
		}
	}
}

// gaussianBlur 类高斯模糊算法
func gaussianBlur(values *[9]color.RGBA64) color.RGBA64 {
	//这边需要 uint32 防止多个uint16相加后溢出
	var r uint32
	var g uint32
	var b uint32
	var a uint32
	for i := 0; i < len(values); i++ {
		if i == 4 { //去掉中间原像素点
			continue
		}
		x := values[i]
		r += uint32(x.R)
		g += uint32(x.G)
		b += uint32(x.B)
		a += uint32(x.A)
	}
	return color.RGBA64{
		uint16(r / 8),
		uint16(g / 8),
		uint16(b / 8),
		uint16(a / 8)}
}

// img2Base64 图片base64
func img2Base64(image *ImageBuf) (string, error) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, image.i); err != nil {
		return "", fmt.Errorf("unable to encode png: %w", err)
	}
	data := buf.Bytes()
	baseImage := fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(data))
	return baseImage, nil
}
