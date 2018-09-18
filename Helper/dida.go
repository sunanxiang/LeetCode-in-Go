package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	gomail "gopkg.in/gomail.v2"
)

const (
	didaTaskFile = "dida.task.txt"
)

func dida(prefix string, p problem) {
	task := p.didaTask(prefix)
	mailToDida(task)
}

func mailToDida(task string) {
	cfg := getConfig()

	if cfg.SMTP == "" || cfg.Port == 0 || cfg.EmailPasswd == "" ||
		cfg.From == "" || cfg.To == "" {
		log.Println("没有配置 Email，无法发送任务")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", cfg.To)
	task += " ^LeetCode "
	task = delay(task)
	m.SetHeader("Subject", task)
	m.SetBody("text/plain", fmt.Sprintf("添加日期 %s", time.Now()))
	d := gomail.NewDialer(cfg.SMTP, cfg.Port, cfg.From, cfg.EmailPasswd)

	if err := d.DialAndSend(m); err != nil {
		log.Println("无法发送任务到 滴答清单：", err)
		saveLocal(task)
		return
	}

	log.Printf("已经在滴答清单中添加任务： %s", task)
}

func saveLocal(task string) {
	ts, err := ioutil.ReadFile(didaTaskFile)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("无法读取 %s：%s\n", didaTaskFile, err)
		}
		f, _ := os.Create(didaTaskFile)
		f.Close()
	}

	ts = append(ts, []byte(task+"\n")...)

	err = ioutil.WriteFile(didaTaskFile, ts, 0755)
	if err != nil {
		log.Fatalf("无法写入 %s: %s\n", didaTaskFile, err)
	}

	log.Printf("新建任务已经写入 %s，请手动添加到滴答清单", didaTaskFile)
}

var m = map[string]time.Duration{
	"#do": 15,
	"#re": 90,
	"#fa": 30,
}

func delay(task string) string {
	key := task[:3]
	if day, ok := m[key]; ok {
		task += time.Now().Add(time.Hour * 24 * day).Format("2006-01-02")
		m[key] += 2
	}
	return task
}
