/*
Package sanitize is a custom library of various sanitation methods to transform data
*/
package sanitize

import (
	"testing"
)

//TestAlpha tests the alpha sanitize method
func TestAlpha(t *testing.T) {
	var originalString = "Test This String-!123"
	var expectedOutput = "Test This String"

	result := Alpha(originalString)

	if result != expectedOutput {
		t.Fatal("Alpha Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestAlphaNumeric tests the alphanumeric sanitize method
func TestAlphaNumeric(t *testing.T) {
	var originalString = "Test This String-!123"
	var expectedOutput = "Test This String123"

	result := AlphaNumeric(originalString)

	if result != expectedOutput {
		t.Fatal("AlphaNumeric Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestDecimal tests the decimal sanitize method
func TestDecimal(t *testing.T) {
	var originalString = "String1.23"
	var expectedOutput = "1.23"

	result := Decimal(originalString)

	if result != expectedOutput {
		t.Fatal("Decimal Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestDomain tests the domain sanitize method
func TestDomain(t *testing.T) {
	var originalString = "http://www.I am a domain.com"
	var expectedOutput = "iamadomain.com"

	result := Domain(originalString)

	if result != expectedOutput {
		t.Fatal("Domain Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestDriversLicense tests the drivers license sanitize method
func TestDriversLicense(t *testing.T) {
	var originalString = "F-1234-5678-!-9"
	var expectedOutput = "F123456789"

	result := DriversLicense(originalString)

	if result != expectedOutput {
		t.Fatal("Drivers License Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestEmail tests the email sanitize method
func TestEmail(t *testing.T) {
	var originalString = "mailto:testME@GmAil.com"
	var expectedOutput = "testme@gmail.com"

	result := Email(originalString)

	if result != expectedOutput {
		t.Fatal("Email Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "test_ME@GmAil.com"
	expectedOutput = "test_me@gmail.com"

	result = Email(originalString)

	if result != expectedOutput {
		t.Fatal("Email Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "test-ME@GmAil.com"
	expectedOutput = "test-me@gmail.com"

	result = Email(originalString)

	if result != expectedOutput {
		t.Fatal("Email Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "test.ME@GmAil.com"
	expectedOutput = "test.me@gmail.com"

	result = Email(originalString)

	if result != expectedOutput {
		t.Fatal("Email Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "test_ME @GmAil.com"
	expectedOutput = "test_me@gmail.com"

	result = Email(originalString)

	if result != expectedOutput {
		t.Fatal("Email Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

}

//TestFirstToUpper tests the first to upper method
func TestFirstToUpper(t *testing.T) {
	originalString := "thisworks"
	expectedOutput := "Thisworks"

	result := FirstToUpper(originalString)

	if result != expectedOutput {
		t.Fatal("FirstToUpper did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "Thisworks"
	expectedOutput = "Thisworks"

	result = FirstToUpper(originalString)

	if result != expectedOutput {
		t.Fatal("FirstToUpper did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestHtml tests the html sanitize method
func TestHtml(t *testing.T) {
	originalString := "<b>This works?</b>"
	expectedOutput := "This works?"

	result := Html(originalString)

	if result != expectedOutput {
		t.Fatal("Html Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "<b>This works?</b><i></i></br><html></html>"
	expectedOutput = "This works?"

	result = Html(originalString)

	if result != expectedOutput {
		t.Fatal("Html Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestIpAddress tests the ip address sanitize method
func TestIpAddress(t *testing.T) {
	var originalString = "192.168.3.6"
	var expectedOutput = "192.168.3.6"

	result := IpAddress(originalString)

	if result != expectedOutput {
		t.Fatal("IpAddress Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "fail"
	expectedOutput = ""

	result = IpAddress(originalString)

	if result != expectedOutput {
		t.Fatal("IpAddress Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "2602:306:bceb:1bd0:44ef:fedb:4f8f:da4f"
	expectedOutput = "2602:306:bceb:1bd0:44ef:fedb:4f8f:da4f"

	result = IpAddress(originalString)

	if result != expectedOutput {
		t.Fatal("IpAddress Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestNameFormal tests the name formal method
func TestNameFormal(t *testing.T) {
	var originalString = "Mark Mc'Cuban-Host"
	var expectedOutput = "Mark Mc'Cuban-Host"

	result := NameFormal(originalString)

	if result != expectedOutput {
		t.Fatal("Numeric Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "Johnny Apple.Seed, Martin"
	expectedOutput = "Johnny Apple.Seed, Martin"

	result = NameFormal(originalString)

	if result != expectedOutput {
		t.Fatal("Numeric Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "Does #Not Work!"
	expectedOutput = "Does Not Work"

	result = NameFormal(originalString)

	if result != expectedOutput {
		t.Fatal("Numeric Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestNumeric tests the numeric sanitize method
func TestNumeric(t *testing.T) {
	var originalString = "Test This String-!1234"
	var expectedOutput = "1234"

	result := Numeric(originalString)

	if result != expectedOutput {
		t.Fatal("Numeric Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestPunctuation tests the name formal method
func TestPunctuation(t *testing.T) {
	var originalString = "Mark Mc'Cuban-Host"
	var expectedOutput = "Mark Mc'Cuban-Host"

	result := Punctuation(originalString)

	if result != expectedOutput {
		t.Fatal("Numeric Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "Johnny Apple.Seed, Martin"
	expectedOutput = "Johnny Apple.Seed, Martin"

	result = Punctuation(originalString)

	if result != expectedOutput {
		t.Fatal("Numeric Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "Does #Not Work!"
	expectedOutput = "Does #Not Work!"

	result = Punctuation(originalString)

	if result != expectedOutput {
		t.Fatal("Numeric Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "Does, 123^* Not & Work!?"
	expectedOutput = "Does, 123 Not & Work!?"

	result = Punctuation(originalString)

	if result != expectedOutput {
		t.Fatal("Numeric Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestSeo tests the seo sanitize method
func TestSeo(t *testing.T) {
	var originalString = "Test This String!"
	var expectedOutput = "TestThisString"

	result := Seo(originalString)

	if result != expectedOutput {
		t.Fatal("SEO Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestSingleLine test the single line sanitize method
func TestSingleLine(t *testing.T) {
	var originalString = `Mark
Mc'Cuban-Host`
	var expectedOutput = "Mark Mc'Cuban-Host"

	result := SingleLine(originalString)

	if result != expectedOutput {
		t.Fatal("Numeric Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = `Mark
Mc'Cuban-Host
something else`
	expectedOutput = "Mark Mc'Cuban-Host something else"

	result = SingleLine(originalString)

	if result != expectedOutput {
		t.Fatal("Numeric Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestTime tests the numeric sanitize method
func TestTime(t *testing.T) {
	var originalString = "t00:00d -EST"
	var expectedOutput = "00:00"

	result := Time(originalString)

	if result != expectedOutput {
		t.Fatal("Time Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestUri tests the uri sanitize method
func TestUri(t *testing.T) {
	var originalString = "Test?=weee! &this=that"
	var expectedOutput = "Test?=weee&this=that"

	result := Uri(originalString)

	if result != expectedOutput {
		t.Fatal("URI Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "Test?=weee! &this=/that/"
	expectedOutput = "Test?=weee&this=/that/"

	result = Uri(originalString)

	if result != expectedOutput {
		t.Fatal("URI Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestUrl tests the url sanitize method
func TestUrl(t *testing.T) {
	var originalString = "Test?=weee! &this=that#works"
	var expectedOutput = "Test?=weee&this=that#works"

	result := Url(originalString)

	if result != expectedOutput {
		t.Fatal("URL Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "Test?=weee! &this=that#works/wee/"
	expectedOutput = "Test?=weee&this=that#works/wee/"

	result = Url(originalString)

	if result != expectedOutput {
		t.Fatal("URL Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestXss tests the xss sanitize method
func TestXss(t *testing.T) {
	originalString := "<script>alert('test');</script>"
	expectedOutput := ">alert('test');</"

	result := Xss(originalString)

	if result != expectedOutput {
		t.Fatal("XSS Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "&lt;script&lt;alert('test');&lt;/script&lt;"
	expectedOutput = "scriptalert('test');/script"

	result = Xss(originalString)

	if result != expectedOutput {
		t.Fatal("XSS Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "javascript:alert('test');"
	expectedOutput = "alert('test');"

	result = Xss(originalString)

	if result != expectedOutput {
		t.Fatal("XSS Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestSocialSecurityNumber tests the social security number sanitize method
func TestSocialSecurityNumber(t *testing.T) {
	originalString := "A123-12-1234"
	expectedOutput := "123-12-1234"

	result := SocialSecurityNumber(originalString)

	if result != expectedOutput {
		t.Fatal("Social Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "A_12s3-1d2-!12f34 "
	expectedOutput = "123-12-1234"

	result = SocialSecurityNumber(originalString)

	if result != expectedOutput {
		t.Fatal("Social Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//TestScripts tests the script removal
func TestScripts(t *testing.T) {
	var originalString = "this <script>$('#something').hide()</script>"
	var expectedOutput = "this "

	result := Scripts(originalString)

	if result != expectedOutput {
		t.Fatal("Scripts Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = "this <script type='text/javascript'>$('#something').hide()</script>"
	expectedOutput = "this "

	result = Scripts(originalString)

	if result != expectedOutput {
		t.Fatal("Scripts Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = `this <script type="text/javascript" class="something">$('#something').hide();</script>`
	expectedOutput = "this "

	result = Scripts(originalString)

	if result != expectedOutput {
		t.Fatal("Scripts Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = `this <iframe width="50" class="something"></iframe>`
	expectedOutput = "this "

	result = Scripts(originalString)

	if result != expectedOutput {
		t.Fatal("Scripts Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = `this <embed width="50" class="something"></embed>`
	expectedOutput = "this "

	result = Scripts(originalString)

	if result != expectedOutput {
		t.Fatal("Scripts Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = `this <object width="50" class="something"></object>`
	expectedOutput = "this "

	result = Scripts(originalString)

	if result != expectedOutput {
		t.Fatal("Scripts Regex did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}
