package helper

const (
	API_VERSION = "v1"
	APIGroup    = "ecommerce"

	BadRequestMassage = "reqest not fulfilled"

	//schedular constants

	HealthCheckRoute = "/health"
	MDBUri           = "localhost:27017"
	Database         = "ecommerce"
	Sender           = ""

	//AuthRoutes
	VerifyEmailRoute = "/verify-email"
	VerifyOtpRoute   = "/verify-otp"
	ResendEmailRoute = "/resend-email"
)

// Time slot for Otp vailidation
const (
	OtpValidation = 60
)

//collections

const (
	VerificationsCollection = "verifications"
)
