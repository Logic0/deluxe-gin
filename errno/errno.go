package errno

const (
	// 兼容逻辑原来的代码，否则改动点太多
	SUCCESS                 = 200
	SUCCESS_MSG             = "ok"
	FORMAT_ERROR            = 3000
	FORMAT_ERROR_MSG        = "request body must be valid JSON"
	PARAM_ERROR             = 3001
	PARAM_ERROR_MSG         = "param error"
	SYSTEM_ERROR            = 3002
	SYSTEM_ERROR_MSG        = "system error, wait and retry"
	NO_PERMISSION           = 3003
	NO_PERMISSION_MSG       = "no auth, check your permission"
	NOT_EXIST               = 3004
	NOT_EXIST_MSG           = "resource required not exists"
	ALREADY_EXIST           = 3005
	ALREADY_EXIST_MSG       = "item already exists"
	ALREADY_PROCESSED       = 3006
	ALREADY_PROCESSED_MSG   = "item had been processed"
	SESSION_ID_CONFILCT     = 3007
	SESSION_ID_CONFLICT_MSG = "session id conflict"
	VEHICLE_ID_INVALID      = 3008 // 车辆不存在

	ON_GOING     = 3100 // 车控命令执行中
	ON_GOING_MSG = "command is running, wait result callback"

	PIN_ERROR                                = 3101 // PIN 码错误
	PIN_NOT_SET                              = 3102 // PIN 码未设置
	EXECUTE_FAILED                           = 3103 // 车控命令执行失败。在所有校验通过之后, 走到实际下发到车执行时才会返回
	PIN_ERROR_TOO_MANY_TIMES                 = 3104 // PIN 码错误太多次
	TSP_RET_USER_NO_PERMISSION               = 3200 // TSP 返回指定的 tsp_userid 没有权限
	TSP_RET_USER_NOT_EXIST                   = 3201 // TSP 返回指定的 tsp_userid 不存在
	TSP_UNKNOWN_FAILED                       = 3300 // 其他原因执行车控失败
	SPEED_NOT_ZERO                           = 3301 // 车速不为0
	TURN_ENGINE_OFF_FAILED                   = 3302 // 引擎熄火失败(关闭空调时有可能失败)
	WAKE_UP_FAILED                           = 3303 // 车辆唤醒失败
	GET_STATUS_TIMEOUT                       = 3304 // 车况查询超时
	TBOX_EXECUTE_FAILED                      = 3305 // TBOX 返回失败
	COMMAND_NOT_SATISFIED                    = 3306 // 不支持当前远控命令、或执行条件不满足
	VEHICLE_IS_BUSY                          = 3307 // 其他命令正在执行中, 车辆忙
	TSP_RET_START_ENGINE_FAILED              = 3308 // 启动发动机失败
	COMMAND_ISSUING_TIMEOUT                  = 3309 // 车控命令下发超时, 包含访问 tsp-access 超时, tsp-access 访问
	COMMAND_OK_BUT_RELATION_OPERATION_FAILED = 3310 // 指令执行成功，但是关联设备没有关闭（比如空调关闭成功，引擎关闭失败）
	ELECTRIC_ERR                             = 3311 // 电车上电错误
	REMOTE_CONTROLE_MODE_CLOSED              = 3312 // 车辆远程控制已关闭，请通过车机打开后重试
	BRAKE_CHARGED_ERR                        = 3313 // 车辆处于本地模式，请松开刹车踏板后重试
	COCKPIT_CLEAN_ERR                        = 3314 // 远程开启座舱清洁失败

)

var errnoMsgMap = map[int]string{
	SUCCESS:                                  "ok",
	FORMAT_ERROR:                             "request body must be valid JSON",
	PARAM_ERROR:                              "param error",
	SYSTEM_ERROR:                             "system error, wait and retry",
	NO_PERMISSION:                            "no auth, check your permission",
	NOT_EXIST:                                "resource required not exists",
	ALREADY_EXIST:                            "item already exists",
	ALREADY_PROCESSED:                        "item had been processed",
	SESSION_ID_CONFILCT:                      "session id conflict",
	VEHICLE_ID_INVALID:                       "invalid vehicle id(maybe not exist)",                            // 车辆不存在
	ON_GOING:                                 "command is running, wait result callback",                       // 车控命令执行中
	PIN_ERROR:                                "pin is wrong",                                                   // PIN 码错误
	PIN_NOT_SET:                              "pin is not set",                                                 // PIN 码未设置
	EXECUTE_FAILED:                           "command execute failed",                                         // 车控命令执行失败。在所有校验通过之后, 走到实际下发到车执行时才会返回
	PIN_ERROR_TOO_MANY_TIMES:                 "invalid pin too many times, operation limited",                  // PIN 码错误太多次
	TSP_RET_USER_NO_PERMISSION:               "TSP returned tsp_userid has no permission",                      // TSP 返回指定的 tsp_userid 没有权限
	TSP_RET_USER_NOT_EXIST:                   "TSP returned tsp_userid not exist",                              // TSP 返回指定的 tsp_userid 不存在
	TSP_UNKNOWN_FAILED:                       "tsp return unknown failed",                                      // 其他原因执行车控失败
	SPEED_NOT_ZERO:                           "speed not zero",                                                 // 车速不为0
	TURN_ENGINE_OFF_FAILED:                   "turn off engine failed",                                         // 引擎熄火失败(关闭空调时有可能失败)
	WAKE_UP_FAILED:                           "wake up vehicle failed",                                         // 车辆唤醒失败
	GET_STATUS_TIMEOUT:                       "get vehicle status failed",                                      // 车况查询超时
	TBOX_EXECUTE_FAILED:                      "tbox execute command failed",                                    // TBOX 返回失败
	COMMAND_NOT_SATISFIED:                    "condition to run command is not satisfied",                      // 不支持当前远控命令、或执行条件不满足
	VEHICLE_IS_BUSY:                          "vehicle is busy, try again later",                               // 其他命令正在执行中, 车辆忙
	TSP_RET_START_ENGINE_FAILED:              "engine start failed",                                            // 发动机启动失败
	COMMAND_ISSUING_TIMEOUT:                  "command issuing timeout, try to get vehicle status to check it", // 车控命令下发超时, 包含访问 tsp-access 超时, tsp-access 访问
	COMMAND_OK_BUT_RELATION_OPERATION_FAILED: "command executed succeed but relation command failed",           // 指令执行成功，但是关联设备没有关闭（比如空调关闭成功，引擎关闭失败）
	ELECTRIC_ERR:                             "failed to power on the vehicle",                                 // 电车上电错误
	REMOTE_CONTROLE_MODE_CLOSED:              "remote control mode closed, please turn it on vehicle",          // 车辆远程控制已关闭，请通过车机打开后重试
	BRAKE_CHARGED_ERR:                        "command executed failed cause brake has been charged",           // 车辆处于本地模式，轻松开刹车踏板后重试
	COCKPIT_CLEAN_ERR:                        "open cockpit clean error,please check vehicle status",           // 远程开座舱清洁执行错误
}

func GetMsg(code int) string {
	msg, ok := errnoMsgMap[code]
	if ok {
		return msg
	}

	return errnoMsgMap[SYSTEM_ERROR]
}
