package orangePi_test_subNodes

// Robot 机器人构成
type Robot struct {
	// 底盘
	chassis *Chassis
	// 摄像头0
	camera0 *Camera
}

// InitRobot 初始化机器人
func InitRobot() *Robot {
	leftMotor := &ChassisMotor{
		EnablePort:    0,
		DirectionPort: 1,
		ActuatePort:   2,
	}
	rightMotor := &ChassisMotor{
		EnablePort:    3,
		DirectionPort: 4,
		ActuatePort:   5,
	}
	GPIO(leftMotor.DirectionPort, 0)
	GPIO(rightMotor.DirectionPort, 0)
	GPIO(leftMotor.EnablePort, 0)
	GPIO(rightMotor.DirectionPort, 0)
	InitPins(leftMotor.ActuatePort, PWMRange)
	InitPins(rightMotor.ActuatePort, PWMRange)
	return &Robot{
		chassis: &Chassis{
			leftMotor:  leftMotor,
			rightMotor: rightMotor,
		},
		camera0: nil,
	}
}
