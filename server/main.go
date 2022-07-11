package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/captchas"
	"strconv"
	"time"
)

//定义临时存放验证码坐标轴的字典
var CaptchaData map[string]int

func main() {
	CaptchaData = make(map[string]int, 0)
	captchas.SlideCaptchaData = make(map[string]int, 0)
	// Example: Get captchas data
	http.HandleFunc("/api/go_captcha_data", captchas.GetCaptchaData)
	// Example: Post check data
	http.HandleFunc("/api/go_captcha_check_data", captchas.CheckCaptcha)
	// Example: demo、JS原生实现
	http.HandleFunc("/go_captcha_example", captchas.Demo)

	// 移动滑块验证码
	// 获取数据
	http.HandleFunc("/api/slide_captcha_data", GetImgTest)
	http.HandleFunc("/api/slide_captcha_check_data", CheckSlideCaptcha)


	http.HandleFunc("/api/slide_captcha2", captchas.NewSlideCaptcha)
	http.HandleFunc("/api/slide_captcha2_check", captchas.CheckSlideCaptcha)


	// 这是静态资源----Vue版本
	static := captchas.GetPWD() + "/static/vue/"
	fsh := http.FileServer(http.Dir(static))
	http.Handle("/go_captcha_demo/", http.StripPrefix("/go_captcha_demo/",fsh))

	// 临时定时清空缓存，由于是demo即在程序内部实现
	captchas.RunTimedTask()

	log.Println("ListenAndServe 0.0.0.0:9001")
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}

func CheckSlideCaptcha(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	id := r.Form.Get("id")
	left := r.Form.Get("left")
	leftInt, err := strconv.Atoi(left)
	if nil != err {
		bt, _ := json.Marshal(map[string]interface{}{
			"code": 500,
			"info": "传输参数错误",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}
	if _, ok := CaptchaData[id]; !ok {
		bt, _ := json.Marshal(map[string]interface{}{
			"code": 500,
			"info": "id不存在",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}

	if (CaptchaData[id] - leftInt) >= 10 {
		bt, _ := json.Marshal(map[string]interface{}{
			"code": 504,
			"info": "验证失败",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}
	bt, _ := json.Marshal(map[string]interface{}{
		"code": 200,
		"info": "",
	})
	_, _ = fmt.Fprintf(w, string(bt))
	return
}

func GetImgTest(w http.ResponseWriter, r *http.Request) {
	cap := captchas.New()
	//n := strconv.Itoa(rand.Intn(7) + 10)
	if err := cap.SetBgImg("./examples/1.png"); err != nil {
		fmt.Println(err)
	}
	if err := cap.SetBgImgLayer("./examples/6.png"); err != nil {
		fmt.Println(err)
	}
	_, im, imSlide := cap.OutImgEncodeString()
	x, y := cap.GetXY()
	id := time.Now().Format("150405") + strconv.Itoa(y)
	CaptchaData[id] = x
	bt, _ := json.Marshal(map[string]interface{}{
		"id": id,
		"y": y,
		"im": im,
		"imSlide": imSlide,
	})
	_, _ = fmt.Fprintf(w, string(bt))
}

