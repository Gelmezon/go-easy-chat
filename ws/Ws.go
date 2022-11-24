package ws

var ChanelMap map[string]chan []byte

func ChanelMapInit() {
	ChanelMap = make(map[string]chan []byte, 100)
}
