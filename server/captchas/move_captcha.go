package captchas

import (
	"crypto/rand"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
	"image/png"
	"math/big"
	"os"
	"server/utils"
	"strconv"
)

func GetRandInt(max int) int {
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(max-1)))
	return int(num.Int64())
}

func CreateCode() (string, int, int) {
	//生成随机数
	nums := GetRandInt(10)
	imageId := utils.RandStringBytesMaskImpr(16)
	f, _ := os.Open("./captcha/bg/" + strconv.Itoa(nums) + ".png")
	//获取随机x坐标
	imageRandX := GetRandInt(480 - 100)
	if imageRandX < 200 {
		imageRandX += 200
	}
	//获取随机y坐标
	imageRandY := GetRandInt(240 - 100)
	if imageRandY < 100 {
		imageRandY += 100
	}
	//转化为image对象
	m, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	//设置截取的最大坐标值和最小坐标值
	maxPotion := image.Point{
		X: imageRandX,
		Y: imageRandY,
	}
	minPotion := image.Point{
		X: imageRandX - 100,
		Y: imageRandY - 100,
	}
	subimg := image.Rectangle{
		Max: maxPotion,
		Min: minPotion,
	}
	f, err = os.Create("./captcha/code/" + imageId + "screenshot.jpeg")
	defer f.Close()
	//截取图像
	data := imaging.Crop(m, subimg)
	jpeg.Encode(f, data, nil)
	//设置遮罩
	createCodeImg("./captcha/bg/"+strconv.Itoa(nums)+".png", minPotion, imageId)
	return imageId, imageRandX, imageRandY
}
func createCodeImg(path string, minPotion image.Point, imageId string) {
	bg, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	maskFile, err := os.Open("./captcha/mask.png")
	if err != nil {
		panic(err)
	}
	bgimg, err := png.Decode(bg)
	maskimg, err := png.Decode(maskFile)
	data := imaging.Overlay(bgimg, maskimg, minPotion, 1.0)
	f, err := os.Create("./captcha/code/" + imageId + ".jpeg")
	defer f.Close()
	jpeg.Encode(f, data, nil)
}