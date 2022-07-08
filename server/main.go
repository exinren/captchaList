package main

import (
	"log"
	"net/http"
	"server/captchas"
)

func main() {
	// Example: Get captchas data
	http.HandleFunc("/api/go_captcha_data", captchas.GetCaptchaData)
	// Example: Post check data
	http.HandleFunc("/api/go_captcha_check_data", captchas.CheckCaptcha)
	// Example: demo、JS原生实现
	http.HandleFunc("/go_captcha_example", captchas.Demo)

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
