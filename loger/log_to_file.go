package loger

import (
	"fmt"
	"os"
)

/*
   func : 启动时,开启goroutine处理通道数据
   param : ch 只读通道  *LogToFileByGoChan类型
*/
//func getChanLogToFile(ch <-chan *LogToFileByGoChan) {
//	for {
//		select {
//		//当通道有值时
//		case logData := <-ch:
//			writeLogToFile(logData.HostName, logData.logMsg)
//		default:
//			time.Sleep(time.Millisecond * 500) //当通道无值时，交出cpu控制权
//		}
//	}
//}



/*
   func  写[通道]日志到文件
   param   errLevel int        错误级别
   param   logFilePath     string      日志文件存放路径
   param   errInfo     string      错误信息
*/
func writeLogToFile(msg LogMsg) {

	// 是否切割文件 获取日志文件地址
	//filePath := fileIsOpenCut(logFilePath)
	logName := ""
	if msg.Mode == 1 {
		logName = msg.LogName
	} else {
		logName = "access"
	}

	//fmt.Println("name:",logName)

	//打开文件 //判断文件是否存在
	f, err := CreatFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("creat file false,err %v \n", err)
		return
	}

	defer f.Close()

	//直接写入字符串数据
	//_, err = f.WriteString(logMsg)
	// []byte类型写入
	_, err = f.Write([]byte(msg.Msg))
	if err != nil {
		fmt.Println("log to file fail,err:", err)
	}

	// 确保写入到磁盘
	f.Sync()
}

//NewLogToFileByGoChan 初始化
//func NewLogToFileByGoChan(level int, hostname, logMsg string) *LogToFileByGoChan {
//	return &LogToFileByGoChan{
//		leveL:    level,
//		logMsg:   logMsg,
//		HostName: hostname, // 域名
//	}
//}
