package service

import (
	"context"
	"fmt"
	"job-portal/cmd/rediss"
	"job-portal/internal/models"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (r NewService) CreateUser(ctx context.Context, nu models.NewUser) (models.User, error) {
	user, err := r.rp.CreateU(ctx, nu)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (r NewService) Authenticate(ctx context.Context, email string, password string) (jwt.RegisteredClaims, error) {
	c, err := r.rp.AuthenticateUser(ctx, email, password)
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}
	return c, nil
}

func (r NewService) CheckEmail(e string) (bool, error) {
	b, err := r.rp.CheckUserEmail(e)
	if err != nil {
		return false, err
	}
	if b {
		otp := generateOTP()
		otpStr := strconv.Itoa(otp)

		from := "vishnuvirat693@gmail.com"
		password := "xjve lhkn iyhg mkco"

		// Recipient's email address
		to := e

		// SMTP server and port
		smtpServer := "smtp.gmail.com"
		smtpPort := 587

		// Message
		subject := "opt regarding job portal"
		body := fmt.Sprintf("one time password is:%s", otpStr)

		// Set up authentication information
		auth := smtp.PlainAuth("", from, password, smtpServer)

		// Compose the email
		message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

		// Connect to the SMTP server
		err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{to}, []byte(message))
		if err != nil {
			fmt.Println("Error sending email:", err)
			return false, err
		}

		fmt.Println("Email sent successfully.")
		rc := rediss.RedisClient()
		ctx := context.Background()
		err = rc.Set(ctx, e, otpStr, 5*time.Minute).Err()
		if err != nil {
			log.Error().Msgf("%v", err)
			return false, err
		}
		r, _ := rc.Get(ctx, e).Result()
		fmt.Println(r)
		return true, nil

	}
	return false, err
}
func (r NewService) UpdatePassword(np models.Reset) (bool, error) {
	rcx := rediss.RedisClient()
	ctx := context.Background()
	otp, err := rcx.Get(ctx, np.Email).Result()
	if err != nil || otp != np.Otp {
		return false, err
	}
	b, err := r.rp.UpdateUserPassword(np)
	if err != nil {
		return false, err
	}
	if b {
		from := "vishnuvirat693@gmail.com"
		password := "xjve lhkn iyhg mkco"

		// Recipient's email address
		to := np.Email

		// SMTP server and port
		smtpServer := "smtp.gmail.com"
		smtpPort := 587

		// Message
		subject := "Regarding recent password update"
		body := "Successfully reset your password for the job portal api"

		// Set up authentication information
		auth := smtp.PlainAuth("", from, password, smtpServer)

		// Compose the email
		message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

		// Connect to the SMTP server
		err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{to}, []byte(message))
		if err != nil {
			fmt.Println("Error sending email:", err)
			return false, err
		}

		fmt.Println("Successfully reset your password.")

		return true, nil

	}
	return false, err
}
func generateOTP() int {
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit number
	otp := rand.Intn(900000) + 100000
	return otp
}
