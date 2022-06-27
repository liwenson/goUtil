package ua

type UserAgent struct{}

func (ua *UserAgent) Random() string {
	return UA.GetAllRandom()
}

func (ua *UserAgent) Chrome() string {
	return UA.GetRandom(CHROME)
}

func (ua *UserAgent) InternetExplorer() string {
	return UA.GetRandom(INTERNET_EXPLORER)
}

func (ua *UserAgent) Firefox() string {
	return UA.GetRandom(FIREFOX)
}

func (ua *UserAgent) Safari() string {
	return UA.GetRandom(SAFARI)
}

func (ua *UserAgent) Android() string {
	return UA.GetRandom(ANDROID)
}

func (ua *UserAgent) MacOSX() string {
	return UA.GetRandom(MAC_OS_X)
}

func (ua *UserAgent) IOS() string {
	return UA.GetRandom(IOS)
}

func (ua *UserAgent) Linux() string {
	return UA.GetRandom(LINUX)
}

func (ua *UserAgent) IPhone() string {
	return UA.GetRandom(IPHONE)
}

func (ua *UserAgent) IPad() string {
	return UA.GetRandom(IPAD)
}

func (ua *UserAgent) Computer() string {
	return UA.GetRandom(COMPUTER)
}

func (ua *UserAgent) Mobile() string {
	return UA.GetRandom(MOBILE)
}
