package orangePi_test_subNodes

/*
#include <stdio.h>
#include <stdlib.h>
#include <wiringPi.h>
#include <softPwm.h>

void init(int pin,int r){
softPwmCreate(pin, 0, r);
}

void pwm(int pin,int range)
{
	softPwmWrite(PWM_PIN, range);
    return 0;
}

*/
import "C"
import "fmt"

// InitPins 初始化
// 传入：需要初始化的Pin,PWM上限
// 传出：无
func InitPins(pin int, r int) {
	C.init(pin, r)
}

// PWMOut 根据占空比 输出PWM 上限
// 传入：引脚，占空比
func PWMOut(pin int, n float64, r int) {
	C.pwm(pin, (int)(n*(float64(r))))
}

// GPIO GPIO输出设定
// 传入：电平
// 传出：无
func GPIO(pin int, io int) {
	fmt.Printf("pin:%d io:%d", pin, io)
}
