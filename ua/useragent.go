package ua

import (
	"encoding/json"
	"math/rand"
	"sync"
	"time"
)

const (
	CHROME            = "chrome"
	INTERNET_EXPLORER = "internet-explorer"
	FIREFOX           = "firefox"
	SAFARI            = "safari"

	ANDROID  = "android"
	MAC_OS_X = "mac-os-x"
	IOS      = "ios"
	LINUX    = "linux"

	IPHONE = "iphone"
	IPAD   = "ipad"

	COMPUTER = "computer"
	MOBILE   = "mobile"
)

var (
	M                    map[string][]string
	BrowserUserAgentMaps = map[string][]string{
		"software_name": {
			CHROME,
			INTERNET_EXPLORER,
			FIREFOX,
			SAFARI,
		},
		"operating_system_name": {
			ANDROID,
			MAC_OS_X,
			IOS,
			LINUX,
		},
		"operating_platform": {
			IPHONE,
			IPAD,
		},
		"hardware_type_specific": {
			COMPUTER,
			MOBILE,
		},
	}
)

func init() {
	f := NewFileCache("./", "fake_useragent.json")

	cacheContent, _ := f.Read()

	json.Unmarshal(cacheContent, &M)

	UA.SetData(M)
}

type useragent struct {
	data map[string][]string
	lock sync.Mutex
}

var (
	UA = useragent{data: make(map[string][]string)}
	r  = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func (u *useragent) Get(key string) []string {
	return u.data[key]
}

func (u *useragent) GetAll() map[string][]string {
	return u.data
}

func (u *useragent) GetRandom(key string) string {
	browser := u.Get(key)
	len := len(browser)
	if len < 1 {
		return ""
	}

	n := r.Intn(len)
	return browser[n]
}

func (u *useragent) GetAllRandom() string {
	browsers := u.GetAll()
	datas := []string{}
	for _, uas := range browsers {
		datas = append(datas, uas...)
	}

	len := len(datas)
	if len < 1 {
		return ""
	}

	n := r.Intn(len)
	return datas[n]
}

func (u *useragent) Set(key, value string) {
	u.lock.Lock()
	defer u.lock.Unlock()
	u.data[key] = append(u.data[key], value)
}

func (u *useragent) SetData(data map[string][]string) {
	u.data = data
}
