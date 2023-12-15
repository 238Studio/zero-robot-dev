package orangePi_test_subNodes

const baseURL = "localhost/ws"

// Robot 机器人构成
type Robot struct {
	// 底盘
	chassis *Chassis
	// 摄像头0
	camera0 *Camera
	// 机械臂
	// 指令传输网络通道
	commandConn *Net
	// 状态返回网络通道
	conditionConn *Net
	// 视频返回网络通道
	videoConn *Net
	// 关机通道
	stopChannel chan struct{}
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
		camera0:       nil,
		commandConn:   InitNet(baseURL + "command:6666"),
		conditionConn: InitNet(baseURL + "condition:6666"),
		videoConn:     InitNet(baseURL + "video:6666"),
		stopChannel:   make(chan struct{}),
	}
}

func (robot *Robot) StartRobot() {
	err := robot.conditionConn.StartWebsocket()
	if err != nil {
		return
	}
	err = robot.commandConn.StartWebsocket()
	if err != nil {
		return
	}
	err = robot.videoConn.StartWebsocket()
	if err != nil {
		return
	}
	for {
		select {
		case <-robot.stopChannel:
			break
		default:
			continue
		}
	}
}
