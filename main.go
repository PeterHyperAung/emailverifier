package emailverifier

import (
	"fmt"
	"log"
	"net"
	"net/smtp"
	"regexp"
	"strings"
)

func ValidateEmailFormat(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func checkAndGetMXRecords(domain string) (bool, []*net.MX) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error while looking up MX records: %v\n", err)
	}

	if len(mxRecords) > 0 {
		return true, mxRecords
	}

	return false, mxRecords
}

func checkSPFRecords(domain string) bool {
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error while looking up SPF records failed: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			return true
		}
	}

	return false
}

func checkDMARCRecord(domain string) bool {
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error while checking DMARC records: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			return true
		}
	}

	return false
}

func verifyEmail(email string, mx *net.MX) bool {
	client, err := smtp.Dial(mx.Host + ":25")
	if err != nil {
		log.Printf("Error while connecting to %s: %v", mx.Host, err)
		return false
	}
	defer client.Close()

	client.Mail("verifier@verifier.com")
	if err := client.Rcpt(email); err != nil {
		log.Printf("Error. Failed to verify email %s: %v", email, err)
		return false
	}

	return true
}

func CheckEmail(email string) bool {
	if !ValidateEmailFormat(email) {
		fmt.Printf("Invalid email format")
		return false
	}

	domain := strings.Split(email, "@")[1]

	hasMX, mxRecords := checkAndGetMXRecords(domain)

	isEmailValid := false
	for _, mx := range mxRecords {
		if verifyEmail(email, mx) {
			isEmailValid = true
			break
		}
	}

	return hasMX && checkSPFRecords(domain) && checkDMARCRecord(domain) && isEmailValid
}
