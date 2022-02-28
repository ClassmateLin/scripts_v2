// @File:  jd_bean_home.go
// @Time:  2022/2/28 4:33 PM
// @Author: ClassmateLin
// @Email: classmatelin.site@gmail.com
// @Site: https://www.classmatelin.top
// @Description:
// @Cron: 45 0 * * *
package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"scripts/config/jd"
	"scripts/constracts"
	"scripts/global"
	"scripts/structs"
	"time"
)

type JdBeanHome struct {
	structs.JdBase
	client *resty.Request
}

// New
// @Description: 初始化
// @receiver JdBeanHome
// @param user
// @return JdBeanHome
func (JdBeanHome) New(user jd.User) constracts.Jd {
	obj := JdBeanHome{}
	obj.User = user
	obj.client = resty.New().R()
	return obj
}

func (j JdBeanHome) request(functionId string, body map[string]interface{}, args ...string) string {
	method := "POST"
	if len(args) >= 1 {
		method = args[0]
	}

	eu := "fafaf"
	fv := "fafasf"
	temp, _ := json.Marshal(body)
	params := map[string]string{
		"functionId":    functionId,
		"appid":         "ld",
		"clientVersion": "10.0.11",
		"client":        "apple",
		"eu":            eu,
		"fv":            fv,
		"osVersion":     "11",
		"uuid":          eu + fv,
		"openudid":      eu + fv,
		"body":          string(temp),
	}

	url := "https://api.m.jd.com/client.action?"

	if method == "GET" {
		resp, err := j.client.SetHeaders(map[string]string{
			"user-agent": global.GetJdUserAgent(),
			"cookie":     j.User.CookieStr,
		}).SetQueryParams(params).Get(url)
		if err != nil {
			return ""
		}
		return resp.String()
	}

	resp, err := j.client.SetQueryParams(params).
		SetHeaders(map[string]string{
			"user-agent": global.GetJdUserAgent(),
			"cookie":     j.User.CookieStr,
		}).Post(url)
	if err != nil {
		return ""
	}
	return resp.String()
}

// doTasks
// @Description: 领取额外京豆任务
// @receiver j
func (j JdBeanHome) doTasks() {
	taskData := j.request("findBeanHome", map[string]interface{}{
		"source": "wojing2", "orderId": "null",
		"rnVersion": "3.9", "rnClient": "1"})

	if code := gjson.Get(taskData, `code`).Int(); code != 0 {
		j.Println("获取任务列表失败...")
		return
	}

	taskProgress := gjson.Get(taskData, `data.taskProgress`).Int()
	taskThreshold := gjson.Get(taskData, `data.taskThreshold`).Int()

	if taskProgress >= taskThreshold {
		j.Println("今日已完成领额外京豆任务...")
		return
	}

	for i := 1; i < 6; i++ {
		taskRes := j.request("beanHomeTask", map[string]interface{}{
			"type":      string(rune(i)),
			"source":    "home",
			"awardFlag": false,
			"itemId":    "",
		})
		if code := gjson.Get(taskRes, `code`).Int(); code == 0 {
			taskProgress = gjson.Get(taskRes, `data.taskProgress`).Int()
			taskThreshold = gjson.Get(taskRes, `data.taskThreshold`).Int()
			j.Println(fmt.Sprintf("领额外京豆任务进度: %d/%d...", taskProgress, taskThreshold))
		}
		time.Sleep(time.Second * 2)
	}
}

// GetAward
// @Description: 领取京豆奖励
// @receiver j
func (j JdBeanHome) GetAward() {
	awardRes := j.request("beanHomeTask", map[string]interface{}{"source": "home", "awardFlag": true})

	if code := gjson.Get(awardRes, `code`).Int(); code == 0 {
		beanNum := gjson.Get(awardRes, `data.beanNum`).Int()
		totalBeanNum := gjson.Get(awardRes, `data.totalUserBean`).Int()
		if beanNum != 0 {
			j.Println(fmt.Sprintf("成功领取%d京豆, 当前京豆总数:%d...", beanNum, totalBeanNum))
		}
	}
}

// GetTitle
// @Description: 脚本名称
// @receiver j
// @return interface{}
func (j JdBeanHome) GetTitle() interface{} {
	return "领京豆"
}

// Exec
// @Description: 脚本入口
// @receiver j
func (j JdBeanHome) Exec() {
	j.doTasks()
	j.GetAward()
}

func main() {
	structs.RunJd(JdBeanHome{}, jd.UserList)
}
