package util

const (
	ErrReadFileBox    = "Read file box fail, filePath: %s"
	ErrMkdirFail      = "Mkdir fail, path: %s"
	ErrReadProfile    = "Read profile fail, file: %s"
	ErrSaveProfile    = "Save profile fail, file: %s"
	ErrSetScavProfile = "Set scav profile fail, sessID: %s"
	ErrOpenFile       = "Open file fail"
	ErrReadFile       = "Read file fail"
	ErrReadDirToJson  = "Read dir to json fail"
	ErrParseJson      = "Parse json fail"

	ErrRouterNotFound = "404"
	ErrHandleReq      = "Handle req fail"
	ErrSendResponse   = "Send response fail"
	ErrReadBody       = "Read body fail"
	ErrIllegalArg     = "Illegal argument"
	ErrUserNotExist   = "User not exist"
	ErrUserExist      = "User exist"

	ErrInvalidRequest    = "Invalid request"
	ErrLoginFail         = "Login fail"
	ErrRegisterFail      = "Register fail"
	ErrDatabaseFileCrash = "Database file crash"
)
