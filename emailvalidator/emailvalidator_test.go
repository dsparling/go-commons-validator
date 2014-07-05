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
 * http://svn.apache.org/viewvc/commons/proper/validator/trunk/src/test/java/org/apache/commons/validator/routines/EmailValidatorTest.java?view=log
 */
package emailvalidator

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
		// UnQuoted Special characters are invalid, except...
		`joe.ok@apache.org`, // . allowed embedded
		// Quoted Special characters are valid
		//`\"joe.\"@apache.org`,
		//`\".joe\"@apache.org`,
		//`\"joe+\"@apache.org`,
		//`\"joe!\"@apache.org`,
		//`\"joe*\"@apache.org`,
		//`\"joe'\"@apache.org`,
		//`\"joe(\"@apache.org`,
		//`\"joe)\"@apache.org`,
		//`\"joe,\"@apache.org`,
		//`\"joe%45\"@apache.org`,
		//`\"joe;\"@apache.org`,
		//`\"joe?\"@apache.org`,
		//`\"joe&\"@apache.org`,
		//`\"joe=\"@apache.org`,
		//`\"..\"@apache.org`,
	}
	for _, email := range validEmails {
		valid := IsValid(email)

		if !valid {
			t.Errorf("expected valid email address:", email)
		}
	}

	invalidEmails := []string{
		// UnQuoted Special characters are invalid
		`joe.@apache.org`,    // . not allowed at end of local part
		`.joe@apache.org`,    // . not allowed at start of local part
		`.@apache.org`,       // . not allowed alone
		`joe..ok@apache.org`, // .. not allowed embedded
		`..@apache.org`,      // .. not allowed alone
		`joe(@apache.org`,
		`joe)@apache.org`,
		`joe,@apache.org"`,
		`joe;@apache.org`,
	}
	for _, email := range invalidEmails {
		valid := IsValid(email)

		if valid {
			t.Errorf("expected invalid email address:", email)
		}
	}
}

/** TODO
 * These test values derive directly from RFC 822 &
 * Mail::RFC822::Address & RFC::RFC822::Address perl test.pl
 * For traceability don't combine these test values with other tests.
 */
//private static final ResultPair[] testEmailFromPerl = {
//	new ResultPair("abigail@example.com", true),
//	new ResultPair("abigail@example.com ", true),
//	new ResultPair(" abigail@example.com", true),
//	new ResultPair("abigail @example.com ", true),
//	new ResultPair("*@example.net", true),
//	new ResultPair("\"\\\"\"@foo.bar", true),
//	new ResultPair("fred&barny@example.com", true),
//	new ResultPair("---@example.com", true),
//	new ResultPair("foo-bar@example.net", true),
//	new ResultPair("\"127.0.0.1\"@[127.0.0.1]", true),
//	new ResultPair("Abigail <abigail@example.com>", true),
//	new ResultPair("Abigail<abigail@example.com>", true),
//	new ResultPair("Abigail<@a,@b,@c:abigail@example.com>", true),
//	new ResultPair("\"This is a phrase\"<abigail@example.com>", true),
//	new ResultPair("\"Abigail \"<abigail@example.com>", true),
//	new ResultPair("\"Joe & J. Harvey\" <example @Org>", true),
//	new ResultPair("Abigail <abigail @ example.com>", true),
//	new ResultPair("Abigail made this < abigail @ example . com >", true),
//	new ResultPair("Abigail(the bitch)@example.com", true),
//	new ResultPair("Abigail <abigail @ example . (bar) com >", true),
//	new ResultPair("Abigail < (one) abigail (two) @(three)example . (bar) com (quz) >", true),
//	new ResultPair("Abigail (foo) (((baz)(nested) (comment)) ! ) < (one) abigail (two) @(three)example . (bar) com (quz) >", true),
//	new ResultPair("Abigail <abigail(fo\\(o)@example.com>", true),
//	new ResultPair("Abigail <abigail(fo\\)o)@example.com> ", true),
//	new ResultPair("(foo) abigail@example.com", true),
//	new ResultPair("abigail@example.com (foo)", true),
//	new ResultPair("\"Abi\\\"gail\" <abigail@example.com>", true),
//	new ResultPair("abigail@[example.com]", true),
//	new ResultPair("abigail@[exa\\[ple.com]", true),
//	new ResultPair("abigail@[exa\\]ple.com]", true),
//	new ResultPair("\":sysmail\"@ Some-Group. Some-Org", true),
//	new ResultPair("Muhammed.(I am the greatest) Ali @(the)Vegas.WBA", true),
//	new ResultPair("mailbox.sub1.sub2@this-domain", true),
//	new ResultPair("sub-net.mailbox@sub-domain.domain", true),
//	new ResultPair("name:;", true),
//	new ResultPair("':;", true),
//	new ResultPair("name: ;", true),
//	new ResultPair("Alfred Neuman <Neuman@BBN-TENEXA>", true),
//	new ResultPair("Neuman@BBN-TENEXA", true),
//	new ResultPair("\"George, Ted\" <Shared@Group.Arpanet>", true),
//	new ResultPair("Wilt . (the Stilt) Chamberlain@NBA.US", true),
//	new ResultPair("Cruisers: Port@Portugal, Jones@SEA;", true),
//	new ResultPair("$@[]", true),
//	new ResultPair("*()@[]", true),
//	new ResultPair("\"quoted ( brackets\" ( a comment )@example.com", true),
//	new ResultPair("\"Joe & J. Harvey\"\\x0D\\x0A <ddd\\@ Org>", true),
//	new ResultPair("\"Joe &\\x0D\\x0A J. Harvey\" <ddd \\@ Org>", true),
//	new ResultPair("Gourmets: Pompous Person <WhoZiWhatZit\\@Cordon-Bleu>,\\x0D\\x0A" +
//	" Childs\\@WGBH.Boston, \"Galloping Gourmet\"\\@\\x0D\\x0A" +
//	" ANT.Down-Under (Australian National Television),\\x0D\\x0A" +
//	" Cheapie\\@Discount-Liquors;", true),
//	new ResultPair(" Just a string", false),
//	new ResultPair("string", false),
//	new ResultPair("(comment)", false),
//	new ResultPair("()@example.com", false),
//	new ResultPair("fred(&)barny@example.com", false),
//	new ResultPair("fred\\ barny@example.com", false),
//	new ResultPair("Abigail <abi gail @ example.com>", false),
//	new ResultPair("Abigail <abigail(fo(o)@example.com>", false),
//	new ResultPair("Abigail <abigail(fo)o)@example.com>", false),
//	new ResultPair("\"Abi\"gail\" <abigail@example.com>", false),
//	new ResultPair("abigail@[exa]ple.com]", false),
//	new ResultPair("abigail@[exa[ple.com]", false),
//	new ResultPair("abigail@[exaple].com]", false),
//	new ResultPair("abigail@", false),
//	new ResultPair("@example.com", false),
//	new ResultPair("phrase: abigail@example.com abigail@example.com ;", false),
//	new ResultPair("invalid�char@example.com", false)
//};

/** TODO
 * Write this test based on perl Mail::RFC822::Address
 * which takes its example email address directly from RFC822
 *
 * FIXME This test fails so disable it with a leading _ for 1.1.4 release.
 * The real solution is to fix the email parsing.
 */
//public void _testEmailFromPerl() {
//	for (int index = 0; index < testEmailFromPerl.length; index++) {
//		String item = testEmailFromPerl[index].item;
//		if (testEmailFromPerl[index].valid) {
//			assertTrue("Should be OK: "+item, validator.isValid(item));
//		} else {
//			assertFalse("Should fail: "+item, validator.isValid(item));
//		}
//	}
//}

func TestValidator293(t *testing.T) {
	validEmails := []string{
		`abc-@abc.com`,
		`abc_@abc.com`,
		`abc-def@abc.com`,
		`abc_def@abc.com`,
	}
	for _, email := range validEmails {
		valid := IsValid(email)

		if !valid {
			t.Errorf("expected invalid email address:", email)
		}
	}

	invalidEmails := []string{
	//	`abc@abc_def.com`,
	}
	for _, email := range invalidEmails {
		valid := IsValid(email)

		if valid {
			t.Errorf("expected invalid email address:", email)
		}
	}
}
