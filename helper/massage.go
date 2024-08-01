package helper

const (
	DB_CONN_FAIL         = "Error connecting to MongoDB: %v"
	DB_CONN_SUCCESSFULLY = "*******************D A T A B A S E C O N N E C T E D********"
	DB_FAIL_CREATE_CLINT = "error creating MongoDB client: %v"
	DB_STRING_NOT_FOUND  = "failed to access db connection string"
	EmailValidationError = "erong Email passed"
	AlreadyVerifiedError = "already verified"
	OtpVerifiedError     = "otp already sent to email"
	OtpNotFound 		 = "otp not found"
	OtpNotMatch          = "otp not match"
	OtpExpire			 = "otp expire"
)
