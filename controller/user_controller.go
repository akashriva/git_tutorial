package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/akashriva/gin_framework/config"
	"github.com/akashriva/gin_framework/helper"
	"github.com/akashriva/gin_framework/models"
	"github.com/akashriva/gin_framework/services"
)

/*
 * Email verify process
 */

func VerifyEmail(ctx *gin.Context) {
	var req models.Verification
	var emailFormateCheck = models.ValidateEmail(req.Email)
	if emailFormateCheck{log.Println(helper.EmailValidationError);return}
	postBodyErr := ctx.BindJSON(&req)
	if postBodyErr != nil {
		log.Println(postBodyErr)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": true, "message": postBodyErr})
		return
	}
	if req.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": true, "message": helper.EmailValidationError})
		return

	}
	resp := config.Mgr.GetSingleRecordByEmail(req.Email, helper.VerificationsCollection)
	//checking if otp expire
	if resp.Otp != "" {
		sec := resp.CreatedAt.Unix() + helper.OtpValidation
		if sec < time.Now().Unix() {
			req, checkEmail := services.SendEmailSendGrid(req)
			if checkEmail != nil {
				log.Println(postBodyErr)
				ctx.JSON(http.StatusBadRequest, gin.H{"error": true, "message": helper.EmailValidationError})
				return
			}

			req.CreatedAt = time.Now()
			//Update opration
			config.Mgr.UpdateVerification(req, helper.VerificationsCollection)
			ctx.JSON(http.StatusOK, gin.H{"error": false, "message": "success"})
			return
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": true, "message":helper.AlreadyVerifiedError})
		}
	}
	req, checkEmail := services.SendEmailSendGrid(req)
	if checkEmail != nil {
		log.Println(postBodyErr)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": true, "message": helper.OtpVerifiedError})
	}

	req.CreatedAt = time.Now()
	//Insertion of Record
	config.Mgr.Insert(req, helper.VerificationsCollection)
	ctx.JSON(http.StatusOK, gin.H{"error": false, "message": "succes"})
}

/*
 * Verify Otp
*/ 
func VerifyOtp(ctx *gin.Context){
	var req models.Verification
	postBodyErr := ctx.BindJSON(&req)
	if postBodyErr != nil {
		log.Println(postBodyErr)
		ctx.JSON(http.StatusBadRequest,gin.H{"error":true,"message":postBodyErr})
		return
	}
	// checking email field Not be Empty
	if req.Email == ""{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":true, "message":helper.EmailValidationError})
		return
	}

	fmt.Printf("********Email****** %s *********OTP******** %s ",req.Email,req.Otp )
	//checking OTP field Not be Empty
	if req.Otp == ""{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":true,"message":helper.OtpNotFound})
		return
	}
	//checking the recode in verification collection
	resp := config.Mgr.GetSingleRecordByEmail(req.Email,helper.VerificationsCollection)
	// if email and otp is already verified
	if resp.Status{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":true, "message":helper.AlreadyVerifiedError})
		return
	}

	sec := resp.CreatedAt.Unix() + helper.OtpValidation
	// Otp not Match
	if req.Otp != resp.Otp{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":true,"message":helper.OtpNotMatch})
	}
	//Otp Expired
	if sec < time.Now().Unix(){
		ctx.JSON(http.StatusBadRequest, gin.H{"error":true,"message":helper.OtpExpire})
		return
	}
	//verified Email
	req.Status = true
	req.CreatedAt = time.Now()
	err := config.Mgr.UpdateEmailVerifiedStatus(req,helper.VerificationsCollection)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":true, "message":helper.AlreadyVerifiedError})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"error":false,"message":"success"})

}

/*
 *Resend the Otp if Email exists
 *
*/ 
func ResendOTPEmail(ctx *gin.Context){ 
	var req models.Verification
	var emailFormateCheck = models.ValidateEmail(req.Email)
	if emailFormateCheck{log.Println(helper.EmailValidationError);return}
	postBodyErr := ctx.BindJSON(&req)
	if postBodyErr != nil {
		log.Println(postBodyErr)
		ctx.JSON(http.StatusBadRequest,gin.H{"error":true,"message":postBodyErr})
		return
	}
	if req.Email == ""{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":true,"message":helper.EmailValidationError})
		return
	}
	resp:= config.Mgr.GetSingleRecordByEmail(req.Email, helper.VerificationsCollection)
	if resp.Email == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":true,"message":helper.EmailValidationError})
		return
	}
	req ,checkEmail := services.SendEmailSendGrid(req)
	if checkEmail != nil{
		log.Println(postBodyErr)
		ctx.JSON(http.StatusBadRequest, gin.H{"error":true, "message":helper.EmailValidationError})
		return
	}
	req.CreatedAt = time.Now()
	config.Mgr.UpdateEmailVerifiedStatus(req, helper.VerificationsCollection)
	ctx.JSON(http.StatusBadRequest, gin.H{"error":false,"message":"success"})
}