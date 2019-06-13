/*
 *Gawain Open Source Project
 *Author: Gawain Antarx
 *Create Date: 2018-Jun-03
 *
 */

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/urfave/cli"
)

const (
	CU   = "10010"
	CMCC = "10086"
	BUPT = "bupt"
)

func main() {
	var user string
	var pass string
	var logout bool
	app := cli.NewApp()
	app.Version = "v1.0-beta"
	app.Usage = "北邮校园网登录软件"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "u",
			Value:       "user",
			Usage:       "请输入学号",
			Destination: &user,
		},
		cli.StringFlag{
			Name:        "p",
			Value:       "password",
			Usage:       "请输入密码",
			Destination: &pass,
		},
		cli.StringFlag{
			Name:  "l",
			Value: "bupt",
			Usage: "请选择运营商：10010 for 中国联通，10086 for 中国移动, bupt for 校园网.",
		},
		cli.BoolFlag{
			Name:        "o",
			Usage:       "Log out.",
			Destination: &logout,
		},
	}
	app.Action = func(c *cli.Context) {
		if c.Bool("o") {
			req := gorequest.New()
			_, _, e := req.Get("http://ngw.bupt.edu.cn/logout").End()
			if e == nil {
				fmt.Println("Logout!")
			}
		} else {
			if user == "user" && pass == "password" {
				fmt.Println("信息有误，请检查输入。")
				os.Exit(-1)
			} else {
				switch c.String("l") {
				case CU:
					fmt.Println(login(user, pass, "CUC-BRAS"))
				case CMCC:
					fmt.Println(login(user, pass, "CMCC-BRAS"))
				case BUPT:
					fmt.Println(login(user, pass, ""))
				}
			}
		}
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func login(u string, p string, l string) string {
	tmpUser := fmt.Sprintf("user=%s", u)
	tmpPass := fmt.Sprintf("pass=%s", p)
	tmpLine := fmt.Sprintf("line=%s", l)

	req := gorequest.New()
	_, body, err := req.Post("http://ngw.bupt.edu.cn/login").Send(tmpUser).Send(tmpPass).Send(tmpLine).End()
	if err != nil {
		return "Network Error."
	} else {
		if strings.Contains(body, "You have successfully logged in") {
			return "Log in successfully"
		} else {
			fmt.Println(body)
			return "Failed."
		}
	}

}
