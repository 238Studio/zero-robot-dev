package orangePi_test_subNodes

import "math"

// Chassis 底盘
type Chassis struct {
	// 左侧电机
	leftMotor *ChassisMotor
	// 右侧电机
	rightMotor *ChassisMotor
}

// Actuate 控制底盘运动
// 传入：运动速度(-1~1) 向前为正,转动速度(-1~1) 向右为正
// 传出：无
func (chassis *Chassis) Actuate(v float64, rotate float64) {
	if v == 0 && rotate == 0 {
		chassis.rightMotor.Actuate(0)
		chassis.leftMotor.Actuate(0)
		return
	}
	speed := math.Sqrt(v*v + rotate*rotate)
	r := rotate / speed
	leftSpeed := (1 - r) * speed
	rightSpeed := (1 - r) * speed
	rightSpeed -= r * speed
	leftSpeed += r * speed
	chassis.rightMotor.Actuate(rightSpeed)
	chassis.leftMotor.Actuate(leftSpeed)
}
