package services

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/akashriva/gin_framework/helper"
	"github.com/akashriva/gin_framework/models"
)

// Struct to hold the OTP data
type OTPData struct {
	OTP string
}

// Function to check if a file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Function to find the absolute path of a file
func FindFilePath(relativePath string) (string, error) {
	if FileExists(relativePath) {
		return filepath.Abs(relativePath)
	}
	return "", fmt.Errorf("file %s does not exist", relativePath)
}

/*
 *Send Email By SendGrid
 */

func SendEmailSendGrid(req models.Verification) (models.Verification, error) {
	seed := time.Now().UnixNano()
	otp := RandomnumString(seed)
	// Find the absolute path of the HTML template
	templatePath, err := FindFilePath("public/Email/otp.html")
	if err != nil {
		log.Println("error finding template: %v", err)
		return req, err
	}

	// Parse the HTML template from the correct path
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("error parsing template:-->", err)
		return req, err
	}

	// Create a buffer to hold the executed template
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, OTPData{OTP: otp}); err != nil {
		log.Println("error executing template:-->", err)
		return req, err
	}

	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return req, errors.New("SENDGRI_API_KEY environment variable is not found")
	}
	//create sendGrid Client
	client := sendgrid.NewSendClient(apiKey)

	//set up email message message
	from := mail.NewEmail("Send name ", helper.Sender)
	to := mail.NewEmail("Recipient Name", req.Email)
	subject := "OTP verification mail"
	req.Otp = otp
	plainTextContent := fmt.Sprintf("Your OTP is: %s", otp)
    htmlContent := tpl.String()
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	//send the email message
	_ ,errResp := client.Send(message)
	if err != nil {
		return req, errResp
	}
	return req, nil
}

/*
 * Generate six digit OTP
 */
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomnumString(seed int64) string {
	rand.Seed(seed)

	digits := make([]rune, 4)
	letters := make([]rune, 2)

	for i := 0; i < 4; i++ {
		digits[i] = rune(rand.Intn(10) + '0')
	}
	for i := 0; i < 2; i++ {
		letters[i] = rune(letterBytes[rand.Intn(len(letterBytes))])
	}

	combined := append(digits, letters...)
	rand.Shuffle(len(combined), func(i, j int) { combined[i], combined[j] = combined[j], combined[i] })

	return string(combined)
}
