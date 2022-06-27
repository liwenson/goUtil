# ua

UserAgent

随机返回一个UserAgent

## 支持

    All User-Agent Random
    Chrome
    InternetExplorer (IE)
    Firefox
    Safari
    Android
    MacOSX
    IOS
    Linux
    IPhone
    IPad
    Computer
    Mobile

## 用法

```go
package main

import (
	"log"

	"github.com/liwenson/goUtil/useragent"
)

func main() {
	// 推荐使用
	ua := useragent.UserAgent{}
	random := ua.Random()
	log.Printf("Random: %s", random)

	chrome := ua.Chrome()
	log.Printf("Chrome: %s", chrome)

	internetExplorer := ua.InternetExplorer()
	log.Printf("IE: %s", internetExplorer)

	firefox := ua.Firefox()
	log.Printf("Firefox: %s", firefox)

	safari := ua.Safari()
	log.Printf("Safari: %s", safari)

	android := ua.Android()
	log.Printf("Android: %s", android)

	macOSX := ua.MacOSX()
	log.Printf("MacOSX: %s", macOSX)

	ios := ua.IOS()
	log.Printf("IOS: %s", ios)

	linux := ua.Linux()
	log.Printf("Linux: %s", linux)

	iphone := ua.IPhone()
	log.Printf("IPhone: %s", iphone)

	ipad := ua.IPad()
	log.Printf("IPad: %s", ipad)

	computer := ua.Computer()
	log.Printf("Computer: %s", computer)

	mobile := ua.Mobile()
	log.Printf("Mobile: %s", mobile)
}

```


