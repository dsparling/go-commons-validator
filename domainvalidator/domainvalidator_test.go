/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * http://svn.apache.org/viewvc/commons/proper/validator/trunk/src/test/java/org/apache/commons/validator/routines/DomainValidatorTest.java?view=log
 */
package domainvalidator

import (
	"testing"
)

func TestValidDomains(t *testing.T) {
	validDomains := []string{
		`apache.org`,         // apache.org should validate
		`www.google.com`,     // www.google.com should validate
		`test-domain.com`,    // test-domain.com should validate
		`test---domain.com`,  // test---domain.com should validate
		`test-d-o-m-ain.com`, // test-d-o-m-ain.com should validate
		`as.uk`,              // two-letter domain label should validate
		`ApAchE.Org`,         // case-insensitive ApAchE.Org should validate
		`z.com`,              // single-character domain label should validate
		`i.have.an-example.domain.name`, // i.have.an-example.domain.name should validate
	}
	for _, domain := range validDomains {
		valid := IsValid(domain)

		if !valid {
			t.Errorf("expected valid domain:", domain)
		}
	}
}

func TestInvalidDomains(t *testing.T) {
	invalidDomains := []string{
		`.org`,             //bare TLD .org shouldn't validate
		` apache.org `,     // domain name with spaces shouldn't validate
		`apa che.org`,      // domain name containing spaces shouldn't validate
		`-testdomain.name`, // domain name starting with dash shouldn't validate
		`testdomain-.name`, // domain name ending with dash shouldn't validate
		`---c.com`,         // domain name starting with multiple dashes shouldn't validate
		`c--.com`,          // domain name ending with multiple dashes shouldn't validate
		//`apache.rog`,            // domain name with invalid TLD shouldn't validate
		`http://www.apache.org`, // URL shouldn't validate
		` `, // Empty string shouldn't validate as domain name
		//``,    assertFalse("Null shouldn't validate as domain name", validator.isValid(null));
	}
	for _, domain := range invalidDomains {
		valid := IsValid(domain)

		if valid {
			t.Errorf("expected invalid domain:", domain)
		}
	}
}

//func TestValidTopLevelDomains(t *testing.T) {
//	// infrastructure TLDs
//	validInfrastructureTopLevelDomains := []string{
//		`.arpa`, // .arpa should validate as iTLD
//	}
//	for _, domain := range validInfrastructureTopLevelDomains {
//		valid := IsInfrastructureTld(domain)
//
//		if !valid {
//			t.Errorf("expected valid instrastructure tld:", domain)
//		}
//	}
//
//	// generic TLDs
//	validGenericTopLevelDomains := []string{
//		`.name`, // .name should validate as gTLD
//	}
//	for _, domain := range validGenericTopLevelDomains {
//		valid := IsGenericTld(domain)
//
//		if !valid {
//			t.Errorf("expected valid generic tld:", domain)
//		}
//	}
//
//	// country code TLDs
//	validCountryCodeTopLevelDomains := []string{
//		`.uk`, // .uk should validate as ccTLD
//	}
//	for _, domain := range validCountryCodeTopLevelDomains {
//		valid := IsValidCountryCodeTld(domain)
//
//		if !valid {
//			t.Errorf("expected valid country code tld:", domain)
//		}
//	}
//
//	// case-insensitive
//	validTopLevelDomains := []string{
//		`.COM`, // .COM should validate as TLD
//		`.BiZ`, // .BiZ should validate as TLD"
//	}
//	for _, domain := range validTopLevelDomains {
//		valid := IsValidTld(domain)
//
//		if !valid {
//			t.Errorf("expected valid tld:", domain)
//		}
//	}
//}

//func TestInvalidTopLevelDomains(t *testing.T) {
//	// infrastructure TLDs
//	invalidInfrastructureTopLevelDomains := []string{
//		`.com`, // .com shouldn't validate as iTLD
//	}
//	for _, domain := range invalidInfrastructureTopLevelDomains {
//		valid := IsInfrastructureTld(domain)
//
//		if valid {
//			t.Errorf("expected invalid instrastructure tld:", domain)
//		}
//	}
//
//	// generic TLDs
//	invalidGenericTopLevelDomains := []string{
//		`.us`, // .us shouldn't validate as gTLD
//	}
//	for _, domain := range invalidGenericTopLevelDomains {
//		valid := IsGenericTld(domain)
//
//		if valid {
//			t.Errorf("expected invalid generic tld:", domain)
//		}
//	}
//
//	// country code TLDs
//	invalidCountryCodeTopLevelDomains := []string{
//		`.org`, // .org shouldn't validate as ccTLD
//	}
//	for _, domain := range invalidCountryCodeTopLevelDomains {
//		valid := IsValidCountryCodeTld(domain)
//
//		if valid {
//			t.Errorf("expected invalid country code tld:", domain)
//		}
//	}
//
//	// corner cases
//	invalidTopLevelDomains := []string{
//		`.nope`, // invalid TLD shouldn't validate
//		``, // empty string shouldn't validate as TLD
//assertFalse("null shouldn't validate as TLD", validator.isValid(null));
//	}
//	for _, domain := range invalidTopLevelDomains {
//		valid := IsValidTld(domain)
//
//		if valid {
//			t.Errorf("expected invalid tld:", domain)
//		}
//	}
//}

// TODO
//public void testAllowLocal() {
//   DomainValidator noLocal = DomainValidator.getInstance(false);
//   DomainValidator allowLocal = DomainValidator.getInstance(true);
//
//   // Default is false, and should use singletons
//   assertEquals(noLocal, validator);
//
//   // Default won't allow local
//   assertFalse("localhost.localdomain should validate", noLocal.isValid("localhost.localdomain"));
//   assertFalse("localhost should validate", noLocal.isValid("localhost"));
//
//   // But it may be requested
//   assertTrue("localhost.localdomain should validate", allowLocal.isValid("localhost.localdomain"));
//   assertTrue("localhost should validate", allowLocal.isValid("localhost"));
//   assertTrue("hostname should validate", allowLocal.isValid("hostname"));
//   assertTrue("machinename should validate", allowLocal.isValid("machinename"));
//
//   // Check the localhost one with a few others
//   assertTrue("apache.org should validate", allowLocal.isValid("apache.org"));
//   assertFalse("domain name with spaces shouldn't validate", allowLocal.isValid(" apache.org "));
//}

//func TestIDN(t *testing.T) {
//	// infrastructure TLDs
//	validIDNs := []string{
//		`www.xn--bcher-kva.ch`, // b\u00fccher.ch in IDN should validate
//	}
//	for _, domain := range validIDNs {
//		valid := IsValid(domain)
//
//		if !valid {
//			t.Errorf("expected valid IDN", domain)
//		}
//	}
//}
