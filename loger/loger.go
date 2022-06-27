package loger

import (
	"fmt"
	"github.com/smallnest/chanx"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	LogDir = "logs"
	LogErr = "error"

	ERR_LEVEL_NO int = iota
	ERR_LEVEL_DEBUG
	ERR_LEVEL_TRACE
	ERR_LEVEL_INFO
	ERR_LEVEL_WARNING
	ERR_LEVEL_ERROR
	ERR_LEVEL_FATAL
	ERR_LEVEL_ACCESS
)

const (
	LogSyncMode = iota // 同步写入模式
	LogPoolMode        // 异步协程池模式
)

var content interface{}

type Log struct {
	Mode int
	Name string
}

// SetMode 配置log日志模式
func (l *Log) SetMode(m int) {
	// 默认模式为同步写入
	l.Mode = m
}

func (l *Log) SetName(name string) {
	l.Name = name
}

func (l *Log) GetName() string {

	if l.Name == "" {
		l.Name = "app"
	}

	return l.Name
}

var LogChan *chanx.UnboundedChan

var log Log

//初始化
func init() {

	LogChan = chanx.NewUnboundedChan(1000)

	log.SetMode(LogPoolMode)

	go getChanLogToFile(LogChan)

}

/*
   func : 启动时,开启goroutine处理通道数据
   param : ch 只读通道  *LogToFileByGoChan类型
*/
func getChanLogToFile(ch *chanx.UnboundedChan) {

	for {
		select {
		//当通道有值时
		case logData := <-ch.Out:
			log := logData.(LogMsg)

			writeLogToFile(log)
		default:
			time.Sleep(time.Millisecond * 500) //当通道无值时，交出cpu控制权
		}
	}
}

/*
func  组装日志的内容
param   errLevel int        错误级别
param   layer int       time.Caller()的层
param   errInfo     string      错误信息
return string
*/

//func LogContent(errLevel, layer int, LogName, Info string, arg ...interface{}) LogMsg {
func LogContent(errLevel, layer int, LogName, Info string, IsContent bool) LogMsg {
	timeString := time.Now().Format("2006-01-02 15:04:05.000")
	txtMap := map[int]string{
		ERR_LEVEL_NO:      "[N]", // normal
		ERR_LEVEL_DEBUG:   "[D]", // debug
		ERR_LEVEL_TRACE:   "[T]", // trace
		ERR_LEVEL_INFO:    "[I]", // info
		ERR_LEVEL_WARNING: "[W]", // warning
		ERR_LEVEL_ERROR:   "[E]", // error
		ERR_LEVEL_FATAL:   "[F]", // fatal
		ERR_LEVEL_ACCESS:  "[A]", // access
	}

	var logMsg LogMsg

	var msg = ""
	//fmt.Println("len: ",len(arg))

	if txtMap[errLevel] == "[A]" {
		// 访问请求不需要代码行数
		msg = timeString + " " + txtMap[errLevel] + " " + fmt.Sprintf("%v", Info) + "" + "\n"

	} else if IsContent {
		msg = timeString + " " + getLayerCode(layer) + " " + txtMap[errLevel] + " " + fmt.Sprintf("%v", Info) + " " + fmt.Sprintf("%v", content) + "\n"
	} else {
		msg = timeString + " " + getLayerCode(layer) + " " + txtMap[errLevel] + " " + fmt.Sprintf("%v", Info) + " " + "\n"
	}

	logMsg.LogName = LogName
	logMsg.Msg = msg

	//fmt.Println(msg)

	return logMsg

	//  fmt.Sprintf("%v",error)  //把错误信息转换成string

}

func getLayerCode(layer int) string {
	//传递参数，可以拿到当前执行程序执行隔了多少层
	pc, file, line, ok := runtime.Caller(layer)
	if !ok {
		return "use runtime.caller() failed," + "no find layer file name," + "no find layer func name"
	}

	//获得报名和文件名
	pkgNameAndFuncName := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	return "[" + pkgNameAndFuncName[0] + "/" + path.Base(file) + " -> " + pkgNameAndFuncName[1] + "()" + ":" + strconv.Itoa(line) + "]"
}

/*
   func : 把日志写入到场景[通道]
   param : errLevel int 错误级别
   param : logFilePath string 日志存放位置
   param : logMsg string 日志信息
   param : arg ...any
*/
func logWriteToScene(errLevel int, logName, logMsg string, content bool) {

	//是否开启日志
	logmsg := LogContent(errLevel, 4, logName, logMsg, content)

	//开启后，日志写到哪的场景
	switch log.Mode {

	case 0: // 文件 no goroutine
		logmsg.Mode = 0
		writeLogToFile(logmsg)

		fmt.Println("写入文件")

	case 1: // 发送到通道 ToChan
		logmsg.Mode = 1
		select {

		case LogChan.In <- logmsg:
		default:
			//当通道满了，走这，丢弃日志，保证不出现阻塞
			fmt.Println("通道满了")

		}

	default: //终端
		fmt.Println("终端...")
		//outputToTerminal(logmsg)
	}

	//close(LogChan.In)
}

func Access(hostname, logInfo string, arg ...interface{}) {
	logWriteToScene(ERR_LEVEL_ACCESS, hostname, logInfo, false)
}

func Trace(logInfo string, arg ...interface{}) {
	if IsInterfaceNil(arg) {
		logWriteToScene(ERR_LEVEL_TRACE, log.GetName(), logInfo, false)
	} else {
		content = arg
		logWriteToScene(ERR_LEVEL_TRACE, log.GetName(), logInfo, true)
	}

}

func Debug(logInfo string, arg ...interface{}) {
	if IsInterfaceNil(arg) {
		logWriteToScene(ERR_LEVEL_DEBUG, log.GetName(), logInfo, false)
	} else {
		content = arg
		logWriteToScene(ERR_LEVEL_DEBUG, log.GetName(), logInfo, true)
	}
}

func Warning(logInfo string, arg ...interface{}) {
	if IsInterfaceNil(arg) {
		logWriteToScene(ERR_LEVEL_WARNING, log.GetName(), logInfo, true)
	} else {
		content = arg
		logWriteToScene(ERR_LEVEL_WARNING, log.GetName(), logInfo, false)
	}

}

func Info(logInfo string, arg ...interface{}) {
	if IsInterfaceNil(arg) {
		logWriteToScene(ERR_LEVEL_INFO, log.GetName(), logInfo, false)
	} else {
		content = arg
		logWriteToScene(ERR_LEVEL_INFO, log.GetName(), logInfo, true)
	}

}

func Error(logInfo string, arg ...interface{}) {

	if IsInterfaceNil(arg) {
		logWriteToScene(ERR_LEVEL_ERROR, LogErr, logInfo, false)
	} else {
		content = arg
		logWriteToScene(ERR_LEVEL_ERROR, LogErr, logInfo, true)
	}

}

func Fatal(logInfo string, arg ...interface{}) {
	if IsInterfaceNil(arg) {
		logWriteToScene(ERR_LEVEL_FATAL, log.GetName(), logInfo, false)
	} else {
		content = arg
		logWriteToScene(ERR_LEVEL_FATAL, log.GetName(), logInfo, true)
	}

}

type LogMsg struct {
	LogName string // 日志文件名称
	Msg     string // 日志内容
	Mode    int    // 日志是否单独记录
}

func Test(name string) {

	var msg = LogMsg{
		LogName: "www.baidu.com",
		Msg:     "Hello world...",
	}

	//for i := 0; i < 500; i++ {
	//	msg := fmt.Sprintf("abcd+%d", i)
	//	LogChan.In <- msg
	//}

	fmt.Println(msg)

	LogChan.In <- msg

	fmt.Println("Hello World!!!", name)
}
