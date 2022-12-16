package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"time"

	"wc_robot/common"
	"wc_robot/common/alapi"
	"wc_robot/handlers"
	"wc_robot/robot"
	"wc_robot/tasks"
)

// 日志设置初始化
func init() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	// 部署在 linux 上可直接通过 nohup ./wc_robot > robot.log & 运行并打印日志
	// 本机测试运行可取消下方注释，记录 log 便于观察

	// // 打印日志到本地 wc_robot.log
	// outputLogPath := "wc_robot.log"
	// f, err := os.Create(outputLogPath)
	// if err != nil {
	// 	log.Println("[WARN]创建日志文件失败, 日志仅输出在控制台")
	// }
	// w := io.MultiWriter(os.Stdout, f)
	// log.SetOutput(w)
}

func main() {
	begin := time.Now()
	defer func() {
		log.Printf("[INFO]本次机器人运行时间为: %s", time.Since(begin).String())
	}()
	r := robot.NewRobot()
	handlers.InitHandlers(r)
	if err := r.Login(); err != nil {
		log.Println(err)
	}
	tasks.InitTasks(common.GetConfig())
	go func() {
		//crontab := cron.New()  默认从分开始进行时间调度
		crontab := cron.New(cron.WithSeconds()) //精确到秒
		//定义定时器调用的任务函数
		task := func() {
			ms := robot.Storage.SearchMembersByNickName(1, "测试")
			if len(ms) < 1 {
				return
			}
			mingyan, err := alapi.GetMingYan()
			if err == nil {
				robot.Storage.Self.SendTextToUser(ms[0], mingyan)
			}
			soul, err := alapi.GetSoul()
			if err == nil {
				robot.Storage.Self.SendTextToUser(ms[0], soul)
			}

			hs, err := alapi.WeiboHotSearch()
			if err == nil {
				robot.Storage.Self.SendTextToUser(ms[0], hs)
			}
			gj, err := alapi.Gjmj()
			if err == nil {
				robot.Storage.Self.SendTextToUser(ms[0], gj)
			}
			caij, err := alapi.Caijing()
			if err == nil {
				robot.Storage.Self.SendTextToUser(ms[0], caij)
			}
			nethot, err := alapi.Networkhot()
			if err == nil {
				robot.Storage.Self.SendTextToUser(ms[0], nethot)
			}
			topnew, err := alapi.Topnews()
			if err == nil {
				robot.Storage.Self.SendTextToUser(ms[0], topnew)
			}

		}
		//定时任务
		spec := "0 0 * * * ?" //cron表达式，每五秒一次
		// 添加定时任务,
		crontab.AddFunc(spec, task)
		// 启动定时器
		crontab.Start()

	}()

	r.Block()
}
