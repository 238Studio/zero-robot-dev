package orangePi_test_subNodes

// PWMRange PWM周期
const PWMRange = 1000

// ChassisMotor 底盘电机
type ChassisMotor struct {
	// 使能引脚
	EnablePort int
	// 方向引脚
	DirectionPort int
	// PWM驱动引脚
	ActuatePort int
}

// Actuate 驱动电机转动
// 传入：(-1~1)电机转动转速
// 传出：无
func (motor *ChassisMotor) Actuate(aboutSpeed float64) {
	GPIO(motor.EnablePort, 1)
	if aboutSpeed >= 0 {
		GPIO(motor.DirectionPort, 1)
	} else {
		GPIO(motor.DirectionPort, -1)
	}
	PWMOut(motor.ActuatePort, aboutSpeed, PWMRange)
}
