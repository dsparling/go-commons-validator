/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * http://svn.apache.org/viewvc/commons/proper/validator/trunk/src/test/java/org/apache/commons/validator/routines/EmailValidatorTest.java?view=markup
 */
package email

import (
	"testing"
)

/**
 * Tests the e-mail validation.
 */
func TestEmail(t *testing.T) {
	email := "jsmith@apache.org"
	valid := IsValid(email)
	if !valid {
		t.Errorf("expected valid email address:", email)
	}
}

/**
 * Tests the email validation with numeric domains.
 */
func TestEmailWithNumericAddress(t *testing.T) {
	validEmails := []string{
		`someone@[216.109.118.76]`,
		`someone@yahoo.com`,
	}
	for _, email := range validEmails {
		valid := IsValid(email)

		if !valid {
			t.Errorf("expected valid email address:", email)
		}
	}
}

/**
* Tests the e-mail validation.
 */
func TestEmailExtension(t *testing.T) {
	validEmails := []string{
		`jsmith@apache.org`,
		`jsmith@apache.com`,
		`jsmith@apache.net`,
		`jsmith@apache.info`,
	}
	for _, email := range validEmails {
		valid := IsValid(email)

		if !valid {
			t.Errorf("expected invalid email address:", email)
		}
	}

	invalidEmails := []string{
		`jsmith@apache.`,
		//`jsmith@apache.c`,
		//`someone@yahoo.museum`,
		//`someone@yahoo.mu-seum`,
	}
	for _, email := range invalidEmails {
		valid := IsValid(email)

		if valid {
			t.Errorf("expected invalid email address:", email)
		}
	}
}

/**
* Tests the e-mail validation with a dash in
* the address.
 */
func TestEmailWithDash(t *testing.T) {
	validEmails := []string{
		`andy.noble@data-workshop.com`,
	}
	for _, email := range validEmails {
		valid := IsValid(email)

		if !valid {
			t.Errorf("expected invalid email address:", email)
		}
	}

	invalidEmails := []string{
	//`andy-noble@data-workshop.-com`,
	//`andy-noble@data-workshop.c-om`,
	//`andy-noble@data-workshop.co-m`,
	}
	for _, email := range invalidEmails {
		valid := IsValid(email)

		if valid {
			t.Errorf("expected invalid email address:", email)
		}
	}
}

/**
* Tests the e-mail validation with a dot at the end of
* the address.
 */
func TestEmailWithDotEnd(t *testing.T) {
	email := "andy.noble@data-workshop.com."
	valid := IsValid(email)
	if valid {
		t.Errorf("expected invalid email address:", email)
	}
}

/**
 * Tests the e-mail validation with an RCS-noncompliant character in
 * the address.
 */
func TestEmailWithBogusCharacter(t *testing.T) {
	validEmails := []string{
		// The ' character is valid in an email username.
		`andy.o'reilly@data-workshop.com`,
		// The + character is valid in an email username.
		`foo+bar@i.am.not.in.us.example.com`,
	}
	for _, email := range validEmails {
		valid := IsValid(email)

		if !valid {
			t.Errorf("expected invalid email address:", email)
		}
	}

	invalidEmails := []string{
	//`andy.noble@\u008fdata-workshop.com`,
	// The ' character is not valid in the domain name.
	//`andy@o'reilly.data-workshop.com`,
	// The + character is not valid in the domain name.
	//`foo+bar@example+3.com`,
	// Domains with only special characters aren't allowed (VALIDATOR-286)
	//`test@%*.com`,
	//`test@^&#.com`,
	}
	for _, email := range invalidEmails {
		valid := IsValid(email)

		if valid {
			t.Errorf("expected invalid email address:", email)
		}
	}
}

/**
 * Tests the email validation with commas.
 */
func TestEmailWithCommas(t *testing.T) {
	invalidEmails := []string{
	//`joeblow@apa,che.org`,
	//`joeblow@apache.o,rg`,
	//`joeblow@apache,org`,
	}
	for _, email := range invalidEmails {
		valid := IsValid(email)

		if valid {
			t.Errorf("expected invalid email address:", email)
		}
	}
}

/**
 * Tests the email validation with spaces.
 */
func TestEmailWithSpaces(t *testing.T) {
	validEmails := []string{
		` joeblow@apache.org`, // TODO - this should be valid?
		`joeblow@apache.org `,
	}
	for _, email := range validEmails {
		valid := IsValid(email)

		if !valid {
			t.Errorf("expected valid email address:", email)
		}
	}

	invalidEmails := []string{
		`joeblow @apache.org`, // TODO - this should be valid?
		//`joeblow@ apache.org`,
		`joe blow@apache.org `,
		//`joeblow@apa che.org `,
	}
	for _, email := range invalidEmails {
		valid := IsValid(email)

		if valid {
			t.Errorf("expected invalid email address:", email)
		}
	}
}

/**
 * Tests the email validation with ascii control characters.
 * (i.e. Ascii chars 0 - 31 and 127)
 */
//func TestEmailWithSpaces(t *testing.T) {
//	invalidEmails := []string{
//	}
//	for _, email := range invalidEmails {
//		valid := IsValid(email)
//
//		if valid {
//			t.Errorf("expected invalid email address:", email)
//		}
//	}
//}

//public void testEmailWithControlChars() {
//	for (char c = 0; c < 32; c++) {
//		assertFalse("Test control char " + ((int)c), validator.isValid("foo" + c + "bar@domain.com"));
//	}
//	assertFalse("Test control char 127", validator.isValid("foo" + ((char)127) + "bar@domain.com"));
//}

/**
 * Test that @localhost and @localhost.localdomain
 * addresses are declared as valid when requested.
 */
//public void testEmailLocalhost() {
//	// Check the default is not to allow
//	EmailValidator noLocal = EmailValidator.getInstance(false);
//	EmailValidator allowLocal = EmailValidator.getInstance(true);
//	assertEquals(validator, noLocal);
//
//	// Depends on the validator
//	assertTrue(
//		"@localhost.localdomain should be accepted but wasn't",
//		allowLocal.isValid("joe@localhost.localdomain")
//	);
//	assertTrue(
//		"@localhost should be accepted but wasn't",
//		allowLocal.isValid("joe@localhost")
//	);
//
//	assertFalse(
//		"@localhost.localdomain should be accepted but wasn't",
//		noLocal.isValid("joe@localhost.localdomain")
//	);
//	assertFalse(
//		"@localhost should be accepted but wasn't",
//		noLocal.isValid("joe@localhost")
//	);
//}

/**
 * VALIDATOR-296 - A / or a ! is valid in the user part,
 * but not in the domain part
 */
func TestEmailWithSlashes(t *testing.T) {
	validEmails := []string{
		// / and ! valid in username
		`joe!/blow@apache.org`,
	}
	for _, email := range validEmails {
		valid := IsValid(email)

		if !valid {
			t.Errorf("expected valid email address:", email)
		}
	}

	invalidEmails := []string{
	// / not valid in domain
	//`joe@ap/ache.org`,
	// ! not valid in domain
	//`joe@apac!he.org`,
	}
	for _, email := range invalidEmails {
		valid := IsValid(email)

		if valid {
			t.Errorf("expected invalid email address:", email)
		}
	}
}

/**
 * Write this test according to parts of RFC, as opposed to the type of character
 * that is being tested.
 */
func TestEmailUserName(t *testing.T) {
	validEmails := []string{
		`joe1blow@apache.org`,
		`joe$blow@apache.org`,
		`joe-@apache.org`,
		`joe_@apache.org`,
		`joe+@apache.org`,   // + is valid unquoted
		`joe!@apache.org`,   // ! is valid unquoted
		`joe*@apache.org`,   // * is valid unquoted
		`joe'@apache.org`,   // ' is valid unquoted
		`joe%45@apache.org`, // % is valid unquoted
		`joe?@apache.org`,   // ? is valid unquoted
		`joe&@apache.org`,   // & ditto
		`joe=@apache.org`,   // = ditto
		`+joe@apache.org`,   // + is valid unquoted
		`!joe@apache.org`,   // ! is valid unquoted
		`*joe@apache.org`,   // * is valid unquoted
		`'joe@apache.org`,   // ' is valid unquoted
		`%joe45@apache.org`, // % is valid unquoted
		`?joe@apache.org`,   // ? is valid unquoted
		`&joe@apache.org`,   // & ditto
		`=joe@apache.org`,   // = ditto
		`+@apache.org`,      // + is valid unquoted
		`!@apache.org`,      // ! is valid unquoted
		`*@apache.org`,      // * is valid unquoted
		`'@apache.org`,      // ' is valid unquoted
		`%@apache.org`,      // % is valid unquoted
		`?@apache.org`,      // ? is valid unquoted
		`&@apache.org`,      // & ditto
		`=@apache.org`,      // = ditto
		//UnQuoted Special characters are invalid
	}
	for _, email := range validEmails {
		valid := IsValid(email)

		if !valid {
			t.Errorf("expected valid email address:", email)
		}
	}

	invalidEmails := []string{}
	for _, email := range invalidEmails {
		valid := IsValid(email)

		if valid {
			t.Errorf("expected invalid email address:", email)
		}
	}
}
