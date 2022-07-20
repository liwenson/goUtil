package loger

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"time"
)

var DefaultPerm = os.ModeDir | os.ModePerm

// CutFileByDateTimeFormat 按日期切割文件的名称时间格式
//const CutFileByDateTimeFormat = "200601021504"
const CutFileByDateTimeFormat = "20060102"

// CutFileBySizeTimeFormat 按文件大小切割文件的名称时间格式
//const CutFileBySizeTimeFormat = "20060102150405"
const CutFileBySizeTimeFormat = "200601021504"

// LogFileMaxSize 文件超过多大时，切割文件
const LogFileMaxSize = 1 * 1024 * 1024 // 1 * 1024 * 1024  //1m

//------------------------创建文件------------------------//

/*
   func    创建目录
   param   dirPath string      文件路径
   param   perm    uint32      文件权限 0644
   return  bool    返回信息
*/

func MkdirFile(dirPath string, perm os.FileMode) (bool bool, errInfo error, ) {
	if perm == 0 {
		perm = DefaultPerm
	}
	err := os.Mkdir(dirPath, perm)
	if err != nil {
		return false, err
	}
	return true, errInfo
}

/*
   func  创建完整目录路径，即中间目录不存在的话也一起创建
   param   dirPath string      文件路径
   param   perm    uint32      文件权限 0644
   return  bool    返回信息
*/

func MkdirAll(dirPath string, perm os.FileMode) (bool bool, errInfo error, ) {
	if perm == 0 {
		perm = DefaultPerm
	}
	err := os.MkdirAll(dirPath, perm)
	if err != nil {
		return false, err
	}
	return true, errInfo
}

// CreatFile 创建文件
func CreatFile(LogName string, flag int, perm os.FileMode) (*os.File, error) {

	if find := strings.Contains(LogName, ":"); find {
		LogName = strings.Replace(LogName, ":", "-", -1)
	}

	filePath := fmt.Sprintf("%s/%s.log", LogDir, LogName)

	_, err := MkdirAll(LogDir, 0666)
	if err != nil {
		return nil, err
	}

	if flag == 0 {
		flag = os.O_CREATE | os.O_APPEND | os.O_TRUNC | os.O_WRONLY
	}
	if perm == 0 {
		flag = 0666
	}

	file, err := os.OpenFile(filePath, flag, perm)

	if err != nil {
		return file, err
	}
	return file, nil
}

// GetFileNameByPath 获取文件的大小 //单位，字节
func GetFileNameByPath(path string) string {

	filetemp := strings.Split(path, "/")
	filename := filetemp[len(filetemp)-1]

	return filename
}

// GetFileSizeByPath 获取文件的大小 单位，字节
func GetFileSizeByPath(path string) int {
	fo, err := os.Stat(path)
	if err != nil {
		//log.Error("open file fail ,file path:" + path,err)
		return 0
	}
	return int(fo.Size())
}

//------------------------bufio读取文件-----------------------------//

// BufioReadFile bufio读取文件  bufio是在file的基础上封装了一层API，支持更多的功能。
//zydh：可自定义读行，分段读
//BufioReadFile 缓冲区读取文件
func BufioReadFile(path string, splitChar byte) {
	if !FileOrDirIsExist(path) {
		fmt.Println("path is not exist")
		return
	}

	if splitChar == 0 {
		splitChar = '\n'
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file) //创建读对象 bufio.NewReader()(带缓冲区的方式打开,适合打开较大的文件)
	for {
		line, err := reader.ReadString(splitChar) //注意是字符，单引号  //\n 一行一行的读取 reader.ReadString()(读取文件)
		if err == io.EOF {                        //io.EOF表示读到文件末尾
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line) //这里不用fmt.Println()
	}
}

//------------------------写文件------------------------//

func WriteFile(txt, path string, flag int, perm os.FileMode) (bool bool, errInfo error) {
	if flag == 0 {
		flag = os.O_CREATE | os.O_APPEND | os.O_TRUNC | os.O_WRONLY
	}
	if perm == 0 {
		flag = 0666
	}
	file, err := os.OpenFile(path, flag, perm)

	if err != nil {
		fmt.Println("open file failed, err:", err)
		return false, err
	}
	defer file.Close()
	//str := "hello 沙河"
	//file.Write([]byte(str))       //写入字节切片数据
	//file.WriteString("hello 小王子") //直接写入字符串数据
	_, err = file.WriteString(txt) //直接写入字符串数据
	if err != nil {
		return false, err
	}
	return true, errInfo
}

//------------------------判断文件是否存在------------------------//

//判断文件是否存在
//go中判断一个文件或者文件夹是否存在方法为：os.Stat() ，通过对返回的错误值进行判断
//1.如果err的值为nil，说明文件或文件夹存在
//2.如果返回的错误类型 使用 os.IsNotExist() 判断为true，说明文件或文件夹不存在
//3.如果返回的错误为其他类型，则不确定是否存在
//所以封装一个函数，用来判断 文件或文件夹是否存在
/*
   func    判断文件或目录是否存在
   link 原文链接：https://blog.csdn.net/leo_jk/article/details/118255913
   author  zydh[朝游东海]
   param   path string         文件路径
   return  bool    返回信息
*/

func FileOrDirIsExist(path string) (bool bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}
	return false

	//通过stat的isDir还可以判断一个路径是文件夹还是文件
	//stat.IsDir()
}

//------------------------删除文件------------------------//

/*
   func    根据文件路径删除文件
   author  zydh[朝游东海]
   param   path string         文件路径
   return  bool    返回信息
*/

func DelFileByPath(path string) (bool bool, errInfo error) {
	if FileOrDirIsExist(path) {
		err := os.Remove(path)
		if err != nil {
			return false, err
		}
		return true, errInfo
	}
	return true, errInfo
}

// FileRename 文件重命名 - ? 毫无用处
func FileRename(oldName, newName string) (bool, error) {
	err := os.Rename(oldName, newName)
	if err != nil {
		return false, err
	}
	return true, nil
}

//GetFileLastPath 获取文件路径 ./zydh/log/log.txt --> ./zydh/log/
func GetFileLastPath(path string) string {
	//分隔
	splice := strings.Split(path, "/")
	return strings.Join(splice[:len(splice)-1], "/")
}

//------------------------切割文件------------------------//
/*
   1.按文件大小切割
   2.按日期切割
*/

/*
   func:FileCutBySize 按文件大小切割  每次记录前判断这个文件的大小
   param : path string 文件路径
   analysis：
       ////os.Rename("./aa/bb/c1/file.go", "./aa/bb/c2/file.go")
       //err := os.Rename(pName,logName)  //会把文件移动到/根目录下
*/

func FileCutBySize(path string) string {
	fmt.Println("FileCutBySize:", path)
	//判断文件是否超出大小
	if GetFileSizeByPath(path) >= LogFileMaxSize {
		//2.备份  xx.log -> xx.log.back202205071155  // warning.txt.bak20220507122541.973
		//在原目录备份                             /      warning.txt             warning.txt.bak20220507122541.973
		err := os.Rename(GetFileLastPath(path)+"/"+GetFileNameByPath(path), getLogFilePathByTime(path, 1))
		if err != nil {
			fmt.Println("zydhfile/file.go FileCutBySize  Backup failed")
		}
	}
	return path
}

//FileCutByTime 按日期切割
/*
   方法一如下
   方法二，在LogToFile{}结构体中新增字段，oldPathFileName,把最新的文件名存入字段，然后在程序中和现在的时间比较，不用查
*/
func FileCutByTime(path string) string {
	//返回路径
	return getLogFilePathByTime(path, 2)
}

/*
   func : 获取新的日志文件路径
   param : path string 原路径
   param : scene int 场景 1 按文件大小 2 按日期
   return : string 根据时间拼接的新路径
*/
func getLogFilePathByTime(path string, scene int) string {
	//文件路径和名词
	pName := GetFileNameByPath(path) //warning.txt
	var nowStr, logName string
	if scene == 1 {
		nowStr = time.Now().Format(CutFileBySizeTimeFormat) //20220507122541.973
		logName = fmt.Sprintf("%s.bak%s", pName, nowStr)    //拼接一个备份
	} else if scene == 2 {
		nowStr = time.Now().Format(CutFileByDateTimeFormat) //20220507122541 //按秒分隔
		//logName = fmt.Sprintf("[%s]%s", nowStr, pName)           //拼接一个备份
		logName = fmt.Sprintf("%s-%s", pName, nowStr) //拼接一个备份
	}

	//[2022-05-07-12:25:41]warning.txt
	return GetFileLastPath(path) + "/" + logName
	// ./zydhlog/[2022-05-07-12:25:41]warning.txt
}

// IsInterfaceNil 判断 interface 是否为空
func IsInterfaceNil(i interface{}) bool {
	vi := reflect.ValueOf(i)

	if vi.Kind() == reflect.Slice {
		return vi.IsNil()
	}

	return false
}

func printTypeValue(slist ...interface{}) string {
	// 字节缓冲作为快速字符串连接
	var b bytes.Buffer
	// 遍历参数
	for _, s := range slist {
		// 将interface{}类型格式化为字符串
		str := fmt.Sprintf("%v", s)
		// 类型的字符串描述
		var typeString string
		// 对s进行类型断言
		switch s.(type) {
		case bool: // 当s为布尔类型时
			typeString = "bool"
		case string: // 当s为字符串类型时
			typeString = "string"
		case int: // 当s为整型类型时
			typeString = "int"
		}
		// 写字符串前缀
		b.WriteString("value: ")
		// 写入值
		b.WriteString(str)
		// 写类型前缀
		b.WriteString(" type: ")
		// 写类型字符串
		b.WriteString(typeString)
		// 写入换行符
		b.WriteString("\n")
	}

	return b.String()
}
