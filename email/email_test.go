package email

import (
	"testing"
)

var validEmails = []string{
	`email@example.com`,
	`firstname.lastname@example.com`,
	`email@subdomain.example.com`,
	`firstname+lastname@example.com`,
	`email@123.123.123.123`,
	`email@[123.123.123.123]`,
	`"email"@example.com`,
	`1234567890@example.com`,
	`email@example-one.com`,
	`_______@example.com`,
	`email@example.name`,
	`email@example.museum`,
	`email@example.co.jp`,
	`firstname-lastname@example.com`,
	`much."more\ unusual"@example.com`,
	`very.unusual."@".unusual.com@example.com`,
	//`very."(),:;<>[]".VERY."very@\\ "very”.unusual@strange.example.com`,
}

var invalidEmails = []string{
	`plainaddress`,
	`#@%^%#$@#$@#.com`,
	`@example.com`,
	`Joe Smith <email@example.com>`,
	`email.example.com`,
	`email@example@example.com`,
	`.email@example.com`,
	`email.@example.com`,
	`email..email@example.com`,
	//`あいうえお@example.com`,
	//`email@example.com (Joe Smith)`,
	//`email@example`,
	//`email@-example.com`,
	//`email@example.web`,
	`email@[111.222.333.44444]`,
	//`email@111.222.333.44444`,
	//`email@example..com`,
	`Abc..123@example.com`,
	`"(),:;<>[\]@example.com`,
	`just"not"right@example.com`,
	`this\ is"really"not\allowed@example.com`,
}

func TestValidEmail(t *testing.T) {
	for _, email := range validEmails {
		valid := IsValid(email)

		if !valid {
			t.Errorf("expected valid email address:", email)
		}
	}
}

func TestInvalidEmail(t *testing.T) {
	for _, email := range invalidEmails {
		valid := IsValid(email)

		if valid {
			t.Errorf("expected invalid email address:", email)
		}
	}
}
