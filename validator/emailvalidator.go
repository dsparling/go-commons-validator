package validator

import (
	//"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	// TODO: error parsing regexp: invalid character class range: `\p{Cntrl}`
	//SPECIAL_CHARS = "\\p{Cntrl}\\(\\)<>@,;:'\\\\\\\"\\.\\[\\]"
	SPECIAL_CHARS = "\\(\\)<>@,;:'\\\\\\\"\\.\\[\\]"
	VALID_CHARS   = "[^\\s" + SPECIAL_CHARS + "]"
	QUOTED_USER   = "(\"[^\"]*\")"
	WORD          = "((" + VALID_CHARS + "|')+|" + QUOTED_USER + ")"

	LEGAL_ASCII_REGEX = `^\\p{ASCII}+$`
	EMAIL_REGEX       = `^(.+)@(.+?)$`
	USER_REGEX        = "^" + WORD + "(\\." + WORD + ")*$"
	IP_DOMAIN_REGEX   = "^\\[(.*)\\]$"

	IPV4_REGEX = "^(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})$"
)

func IsValid(emailAddress string) bool {
	emailAddress = strings.TrimSpace(emailAddress)

	if emailAddress == "" {
		return false
	}

	r, _ := regexp.Compile(LEGAL_ASCII_REGEX)
	match := r.MatchString(emailAddress)
	if match {
		return false
	}

	// Check the whole email address structure
	r2, _ := regexp.Compile(EMAIL_REGEX)
	match2 := r2.MatchString(emailAddress)
	if !match2 {
		return false
	}

	if strings.HasSuffix(emailAddress, ".") {
		return false
	}

	result := r2.FindStringSubmatch(emailAddress)
	if len(result) < 3 {
		return false
	}
	user := result[1]
	domain := result[2]

	if !isValidUser(user) {
		return false
	}

	if !isValidDomain(domain) {
		return false
	}

	return true
}

func isValidUser(user string) bool {
	r, _ := regexp.Compile(USER_REGEX)
	return r.MatchString(user)
}

func isValidDomain(domain string) bool {
	// see if domain is an IP address in brackets
	//	Matcher ipDomainMatcher = IP_DOMAIN_PATTERN.matcher(domain);
	r, _ := regexp.Compile(IP_DOMAIN_REGEX)

	match := r.MatchString(domain)
	if match {
		// Domain is IP address in brackets
		//fmt.Println("Domain is IP address in brackets")
		//fmt.Println(domain)
		groups := r.FindStringSubmatch(domain)
		if len(groups) < 2 {
			return false
		}
		inet4Address := groups[1]
		//fmt.Println("inet4Address")
		//fmt.Println(inet4Address)

		r2, _ := regexp.Compile(IPV4_REGEX)
		match2 := r2.MatchString(inet4Address)
		if match2 {
			// Check if it's an Inet4 Address
			groups2 := r2.FindStringSubmatch(inet4Address)
			if len(groups2) < 5 {
				return false
			}
			// TODO: IP6
			//fmt.Println("IP4")

			// verify that address subgroups are legal
			for i := 1; i <= len(groups2)-1; i++ {
				ipSegment := groups2[i]
				//fmt.Println(ipSegment)

				if ipSegment == "" || len(ipSegment) <= 0 {
					return false
				}

				// Make sure it's a number
				iIpSegment, err := strconv.Atoi(ipSegment)
				if err != nil {
					return false
				}

				if iIpSegment > 255 {
					return false
				}
			}
			return true
		} else {
			// More than four segments, or a segment has more than three digits
			return false
		}
	} else {
		// TODO: Domain is symbolic name
		//fmt.Println("Domain is symbolic name")
		//fmt.Println(domain)
	}
	// TODO
	//	if (ipDomainMatcher.matches()) {
	//	InetAddressValidator inetAddressValidator =
	//	InetAddressValidator.getInstance();
	//	return inetAddressValidator.isValid(ipDomainMatcher.group(1));
	//	} else {
	//	// Domain is symbolic name
	//	DomainValidator domainValidator =
	//	DomainValidator.getInstance(allowLocal);
	//	return domainValidator.isValid(domain);
	//	}

	return true
}
