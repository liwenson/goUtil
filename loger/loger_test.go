package loger

import "testing"

func TestLoger(t *testing.T) {

	var log Log
	log.SetName("test")
	for i := 0; i < 100; i++ {
		Nginx("nginx","sdfasdfasdfasdf")
	}


}
