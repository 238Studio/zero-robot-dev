package orangePi_test_subNodes

import "fmt"

// Init 初始化
// 传入：无
// 传出：无
func InitPins(pin int, r int) {

}

// PWMOut 根据占空比 输出PWM 周期默认是1000 (ms)
// 传入：引脚，占空比，持续时间(10ms)
func PWMOut(pin int, n float64, r int) {
	fmt.Printf("pin:%d n:%f", pin, n)
}

// GPIO GPIO输出设定
// 传入：电平
// 传出：无
func GPIO(pin int, io int) {
	fmt.Printf("pin:%d io:%d", pin, io)
}
