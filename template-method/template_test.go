package template_method

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.getAndSendOTP(4)

	fmt.Println()

	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.getAndSendOTP(4)
}
