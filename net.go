package orangePi_test_subNodes

import "github.com/gorilla/websocket"

type Net struct {
	// 网络连接
	conn *websocket.Conn
	// 停止监听通道
	stopChannel *chan struct{}
	// url
	url string
}

// StartWebsocket 开启websocket
// 传入：地址
// 传出：错误
func (net *Net) StartWebsocket() error {
	conn, _, err := websocket.DefaultDialer.Dial("ws://"+net.url, nil)
	net.conn = conn
	return err
}

// WriteBinMessage 写入二进制消息
// 传入：消息
// 传出：错误
func (net *Net) WriteBinMessage(data *[]byte) error {
	return net.conn.WriteMessage(websocket.BinaryMessage, *data)
}

// WriteTextMessage 写入文字消息
// 传入：消息
// 传出：错误
func (net *Net) WriteTextMessage(data *string) error {
	return net.conn.WriteMessage(websocket.TextMessage, []byte(*data))
}

// GetMessageChannel 获取消息通道
// 传入：无
// 传出：二进制通道，文本消息通道
func (net *Net) GetMessageChannel() (*chan *[]byte, *chan *string) {
	binChannel := make(chan *[]byte)
	textChannel := make(chan *string)
	go func() {
		for {
			select {
			case <-*net.stopChannel:
				break
			default:
				messageType, data, err := net.conn.ReadMessage()
				if err != nil {
					break
				}
				if messageType == websocket.BinaryMessage {
					binChannel <- &data
				} else if messageType == websocket.TextMessage {
					str := string(data)
					textChannel <- &str
				}
			}
		}
	}()
	return &binChannel, &textChannel
}

// StopListen 终止监听
// 传入：无
// 传出：无
func (net *Net) StopListen() {
	*net.stopChannel <- struct{}{}
}

// StopConnection 终止网络连接
// 传入：无
// 传出：错误
func (net *Net) StopConnection() error {
	return net.conn.Close()
}

// InitNet 初始化网络包
// 传入：url
// 传出：无
func InitNet(url string) *Net {
	return &Net{
		conn:        nil,
		stopChannel: nil,
		url:         url,
	}
}
