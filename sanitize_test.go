/*
Package gosanitize implements a simple library of various sanitation methods for data transformation.

If you have any suggestions or comments, please feel free to open an issue on this project's GitHub page.

Author: MrZ
*/
package gosanitize

import (
	"fmt"
	"testing"
)

//TestAlpha tests the alpha sanitize method
func TestAlpha(t *testing.T) {

	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Test removing spaces and punctuation - preserve the case of the letters
	originalString = "Test This String-!123"
	expectedOutput = "TestThisString"
	methodName = "Alpha"

	result := Alpha(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing various symbols
	originalString = `~!@#$%^&*()-_Symbols=+[{]};:'"<>,./?`
	expectedOutput = "Symbols"

	result = Alpha(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing all carriage returns
	originalString = "\nThis\nThat"
	expectedOutput = "ThisThat"

	result = Alpha(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing fancy quotes and microsoft symbols
	originalString = "“This is a quote with tick`s … ” ☺ "
	expectedOutput = "Thisisaquotewithticks"

	result = Alpha(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//
	//==================================================================================================================
	//

	//Test removing spaces and punctuation - preserve the case of the letters
	originalString = "Test This String-!123"
	expectedOutput = "Test This String"

	result = Alpha(originalString, true)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing various symbols
	originalString = `~!@#$%^&*()-_Symbols=+[{]};:'"<>,./?`
	expectedOutput = "Symbols"

	result = Alpha(originalString, true)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing all carriage returns
	originalString = "\nThis\nThat"
	expectedOutput = `
This
That`

	result = Alpha(originalString, true)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing fancy quotes and microsoft symbols
	originalString = "“This is a quote with tick`s … ” ☺ "
	expectedOutput = "This is a quote with ticks    "

	result = Alpha(originalString, true)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkAlphaNoSpaces benchmarks the Alpha method
func BenchmarkAlphaNoSpaces(b *testing.B) {
	testString := "This is the test string."
	for i := 0; i < b.N; i++ {
		_ = Alpha(testString, false)
	}
}

//ExampleAlpha_noSpaces example with no spaces
func ExampleAlpha_noSpaces() {
	fmt.Println(Alpha("Example String!", false))
	// Output: Example String!
}

//BenchmarkAlphaWithSpaces benchmarks the Alpha method
func BenchmarkAlphaWithSpaces(b *testing.B) {
	testString := "This is the test string."
	for i := 0; i < b.N; i++ {
		_ = Alpha(testString, true)
	}
}

//TestAlphaNumeric tests the alpha numeric sanitize method
func TestAlphaNumeric(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Test the base string with mixed characters
	originalString = "Test This String-!123"
	expectedOutput = "TestThisString123"
	methodName = "AlphaNumeric"

	result := AlphaNumeric(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing various symbols
	originalString = `~!@#$%^&*()-_Symbols=+[{]};:'"<>,./?`
	expectedOutput = "Symbols"

	result = AlphaNumeric(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing all carriage returns
	originalString = "\nThis1\nThat2"
	expectedOutput = "This1That2"

	result = AlphaNumeric(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing fancy quotes and microsoft symbols
	originalString = "“This is a quote with tick`s … ” ☺ 342"
	expectedOutput = "Thisisaquotewithticks342"

	result = AlphaNumeric(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//
	//==================================================================================================================
	//

	//Test removing spaces and punctuation - preserve the case of the letters
	originalString = "Test This String-! 123"
	expectedOutput = "Test This String 123"

	result = AlphaNumeric(originalString, true)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing various symbols
	originalString = `~!@#$%^&*()-_Symbols 123=+[{]};:'"<>,./?`
	expectedOutput = "Symbols 123"

	result = AlphaNumeric(originalString, true)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing all carriage returns
	originalString = "\nThis1\nThat2"
	expectedOutput = `
This1
That2`

	result = AlphaNumeric(originalString, true)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing fancy quotes and microsoft symbols
	originalString = "“This is a quote with tick`s…”☺ 123"
	expectedOutput = "This is a quote with ticks 123"

	result = AlphaNumeric(originalString, true)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkAlphaNumericNoSpaces benchmarks the AlphaNumeric method
func BenchmarkAlphaNumericNoSpaces(b *testing.B) {
	testString := "This is the test string 12345."
	for i := 0; i < b.N; i++ {
		_ = AlphaNumeric(testString, false)
	}
}

//BenchmarkAlphaNumericWithSpaces benchmarks the AlphaNumeric method
func BenchmarkAlphaNumericWithSpaces(b *testing.B) {
	testString := "This is the test string 12345."
	for i := 0; i < b.N; i++ {
		_ = AlphaNumeric(testString, true)
	}
}

//TestCustom tests the custom sanitize method
func TestCustom(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
		regString      string
	)

	//Test custom Alpha Numeric
	originalString = "ThisWorks123!"
	expectedOutput = "ThisWorks123"
	methodName = "Custom"
	regString = `[^a-zA-Z0-9]`

	result := Custom(originalString, regString)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test custom Decimal
	originalString = "ThisWorks1.23!"
	expectedOutput = "1.23"
	regString = `[^0-9.-]`

	result = Custom(originalString, regString)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test invalid regString
	//
	// will panic()
	//
}

//BenchmarkCustom benchmarks the Custom method
func BenchmarkCustom(b *testing.B) {
	testString := "This is the test string 12345."
	for i := 0; i < b.N; i++ {
		_ = Custom(testString, `[^a-zA-Z0-9]`)
	}
}

//TestDecimal tests the decimal sanitize method
func TestDecimal(t *testing.T) {

	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Combined with letters
	originalString = "String1.23"
	expectedOutput = "1.23"
	methodName = "Decimal"

	result := Decimal(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove all symbols, spaces, words
	originalString = "  $-1.034234  Price"
	expectedOutput = "-1.034234"

	result = Decimal(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//More additional symbols, line returns
	originalString = "/n<<  $-1.034234  >>/n"
	expectedOutput = "-1.034234"

	result = Decimal(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkDecimal benchmarks the Decimal method
func BenchmarkDecimal(b *testing.B) {
	testString := "String: -123.12345"
	for i := 0; i < b.N; i++ {
		_ = Decimal(testString)
	}
}

//TestDomain tests the domain sanitize method
func TestDomain(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Start with an invalid domain name
	originalString = "http://www.I am a domain.com"
	expectedOutput = "Iamadomain.com"
	methodName = "Domain"

	result, err := Domain(originalString, true, false)
	if err == nil {
		t.Fatal(methodName, "method should have failed here. Expected:", expectedOutput, "original string:", originalString, "returned: ", result, err)
	}

	//Another invalid domain name
	originalString = "!I_am a domain.com"
	expectedOutput = ""

	result, err = Domain(originalString, true, false)
	if err == nil {
		t.Fatal(methodName, "method did not function properly. Expected:", expectedOutput, "original string:", originalString, "returned: ", result, err)
	}

	//Another invalid domain name
	originalString = ""
	expectedOutput = ""

	result, err = Domain(originalString, true, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not function properly. Expected:", expectedOutput, "original string:", originalString, "returned: ", result, err)
	} else if err != nil {
		t.Fatal(methodName, "method did not function properly. Expected:", expectedOutput, "original string:", originalString, "returned: ", result, err)
	}

	//A valid url (preserve the case)
	originalString = "http://IAmDomain.com"
	expectedOutput = "IAmDomain.com"

	result, err = Domain(originalString, true, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not function properly. Expected:", expectedOutput, "original string:", originalString, "returned: ", result, err)
	}

	//A valid url - lowercase
	originalString = "http://IAmDomain.com"
	expectedOutput = "iamdomain.com"

	result, err = Domain(originalString, false, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not function properly. Expected:", expectedOutput, "original string:", originalString, "returned: ", result, err)
	}

	//A valid url - lowercase
	originalString = "https://IAmDomain.com/?this=that#plusThis"
	expectedOutput = "iamdomain.com"

	result, err = Domain(originalString, false, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not function properly. Expected:", expectedOutput, "original string:", originalString, "returned: ", result, err)
	}

	//A valid url - lowercase
	originalString = "http://www.IAmDomain.com/?this=that#plusThis"
	expectedOutput = "iamdomain.com"

	result, err = Domain(originalString, false, true)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not function properly. Expected:", expectedOutput, "original string:", originalString, "returned: ", result, err)
	}

	//A valid url - lowercase, remove WWW
	originalString = "IAmDomain.com/?this=that#plusThis"
	expectedOutput = "iamdomain.com"

	result, err = Domain(originalString, false, true)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not function properly. Expected:", expectedOutput, "original string:", originalString, "returned: ", result, err)
	}

	//A valid url - lowercase, leave WWW
	originalString = "www.IAmDomain.com/?this=that#plusThis"
	expectedOutput = "www.iamdomain.com"

	result, err = Domain(originalString, false, false)
	if result != expectedOutput {
		t.Fatal(methodName, "method did not function properly. Expected:", expectedOutput, "original string:", originalString, "returned: ", result, err)
	}

}

//BenchmarkDomain benchmarks the Domain method
func BenchmarkDomain(b *testing.B) {
	testString := "https://Example.COM/?param=value"
	for i := 0; i < b.N; i++ {
		_, _ = Domain(testString, false, false)
	}
}

//BenchmarkDomainPreserveCase benchmarks the Domain method
func BenchmarkDomainPreserveCase(b *testing.B) {
	testString := "https://Example.COM/?param=value"
	for i := 0; i < b.N; i++ {
		_, _ = Domain(testString, true, false)
	}
}

//BenchmarkDomainRemoveWww benchmarks the Domain method
func BenchmarkDomainRemoveWww(b *testing.B) {
	testString := "https://Example.COM/?param=value"
	for i := 0; i < b.N; i++ {
		_, _ = Domain(testString, false, true)
	}
}

//TestEmail tests the email sanitize method
func TestEmail(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Test removing the mailto: and lower case
	methodName = "Email"
	originalString = "mailto:testME@GmAil.com"
	expectedOutput = "testme@gmail.com"

	result := Email(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Lowercase
	originalString = "test_ME@GmAil.com"
	expectedOutput = "test_me@gmail.com"

	result = Email(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Supports valid email, lowercase
	originalString = "test-ME@GmAil.com"
	expectedOutput = "test-me@gmail.com"

	result = Email(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Valid email, lowercase
	originalString = "test.ME@GmAil.com"
	expectedOutput = "test.me@gmail.com"

	result = Email(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove all spaces
	originalString = " test_ME @GmAil.com "
	expectedOutput = "test_me@gmail.com"

	result = Email(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove all invalid characters
	originalString = " <<test_ME @GmAil.com!>> "
	expectedOutput = "test_me@gmail.com"

	result = Email(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Allowed plus signs
	originalString = " test_ME+2@GmAil.com "
	expectedOutput = "test_me+2@gmail.com"

	result = Email(originalString, false)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Allowed plus signs (preserve case)
	originalString = " test_ME+2@GmAil.com "
	expectedOutput = "test_ME+2@GmAil.com"

	result = Email(originalString, true)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkEmail benchmarks the Email method
func BenchmarkEmail(b *testing.B) {
	testString := "mailto:Person@Example.COM "
	for i := 0; i < b.N; i++ {
		_ = Email(testString, false)
	}
}

//BenchmarkEmailPreserveCase benchmarks the Email method
func BenchmarkEmailPreserveCase(b *testing.B) {
	testString := "mailto:Person@Example.COM "
	for i := 0; i < b.N; i++ {
		_ = Email(testString, true)
	}
}

//TestFirstToUpper tests the first to upper method
func TestFirstToUpper(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Test turning to uppercase
	originalString = "thisworks"
	expectedOutput = "Thisworks"
	methodName = "FirstToUpper"

	result := FirstToUpper(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test keeping it uppercase
	originalString = "Thisworks"
	expectedOutput = "Thisworks"

	result = FirstToUpper(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Convert first letter to uppercase
	originalString = "this"
	expectedOutput = "This"

	result = FirstToUpper(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Single letter test
	originalString = "t"
	expectedOutput = "T"

	result = FirstToUpper(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Two letter test
	originalString = "tt"
	expectedOutput = "Tt"

	result = FirstToUpper(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkFirstToUpper benchmarks the FirstToUpper method
func BenchmarkFirstToUpper(b *testing.B) {
	testString := "make this upper"
	for i := 0; i < b.N; i++ {
		_ = FirstToUpper(testString)
	}
}

//TestFormalName tests the formal name method
func TestFormalName(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Test a valid name
	originalString = "Mark Mc'Cuban-Host"
	expectedOutput = "Mark Mc'Cuban-Host"
	methodName = "FormalName"

	result := FormalName(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test a valid name
	originalString = "Mark Mc'Cuban-Host the SR."
	expectedOutput = "Mark Mc'Cuban-Host the SR."

	result = FormalName(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test a valid name
	originalString = "Mark Mc'Cuban-Host the Second."
	expectedOutput = "Mark Mc'Cuban-Host the Second."

	result = FormalName(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test another valid name
	originalString = "Johnny Apple.Seed, Martin"
	expectedOutput = "Johnny Apple.Seed, Martin"

	result = FormalName(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test invalid characters
	originalString = "Does #Not Work!"
	expectedOutput = "Does Not Work"

	result = FormalName(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkFormalName benchmarks the FormalName method
func BenchmarkFormalName(b *testing.B) {
	testString := "John McDonald Jr."
	for i := 0; i < b.N; i++ {
		_ = FormalName(testString)
	}
}

//TestHTML tests the HTML sanitize method
func TestHTML(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Test basic HTML removal
	methodName = "HTML"
	originalString = "<b>This works?</b>"
	expectedOutput = "This works?"

	result := HTML(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Advanced HTML removal
	originalString = "<html><b>This works?</b><i></i></br></html>"
	expectedOutput = "This works?"

	result = HTML(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkHTML benchmarks the HTML method
func BenchmarkHTML(b *testing.B) {
	testString := "<html><b>Test This!</b></html>"
	for i := 0; i < b.N; i++ {
		_ = FormalName(testString)
	}
}

//TestIPAddress tests the ip address sanitize method
func TestIPAddress(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Basic IPV4 check
	originalString = "192.168.3.6"
	expectedOutput = "192.168.3.6"
	methodName = "IPAddress"

	result := IPAddress(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Basic IPV4 check (gateway mask)
	originalString = "255.255.255.255"
	expectedOutput = "255.255.255.255"

	result = IPAddress(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Invalid IPV4 out of range
	originalString = "304.255.255.255"
	expectedOutput = ""

	result = IPAddress(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Invalid ip address
	originalString = "fail"
	expectedOutput = ""

	result = IPAddress(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Invalid ip address
	originalString = "192-123-122-123"
	expectedOutput = ""

	result = IPAddress(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Valid IPV6
	originalString = "2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f"
	expectedOutput = "2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f"

	result = IPAddress(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Valid formatted IPV6
	originalString = "2602:305:bceb:1bd0:44ef:2:2:2"
	expectedOutput = "2602:305:bceb:1bd0:44ef:2:2:2"

	result = IPAddress(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Valid formatted IPV6
	originalString = "2:2:2:2:2:2:2:2"
	expectedOutput = "2:2:2:2:2:2:2:2"

	result = IPAddress(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Invalid IPv4
	originalString = "192.2"
	expectedOutput = ""

	result = IPAddress(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Invalid IPV4 - bad character and too short
	originalString = "192.2! "
	expectedOutput = ""

	result = IPAddress(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Invalid string characters
	originalString = "IP: 192.168.0.1 "
	expectedOutput = ""

	result = IPAddress(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove space characters
	originalString = " 192.168.0.1 "
	expectedOutput = "192.168.0.1"

	result = IPAddress(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkIPAddress benchmarks the IPAddress method
func BenchmarkIPAddress(b *testing.B) {
	testString := " 192.168.0.1 "
	for i := 0; i < b.N; i++ {
		_ = IPAddress(testString)
	}
}

//TestNumeric tests the numeric sanitize method
func TestNumeric(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Remove everything and leave just numbers
	originalString = "Test This String-!1234"
	expectedOutput = "1234"
	methodName = "Numeric"

	result := Numeric(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove everything and leave just numbers
	originalString = " $1.00 Price!"
	expectedOutput = "100"

	result = Numeric(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkNumeric benchmarks the numeric method
func BenchmarkNumeric(b *testing.B) {
	testString := " 192.168.0.1 "
	for i := 0; i < b.N; i++ {
		_ = Numeric(testString)
	}
}

//TestPathName tests the path name sanitize method
func TestPathName(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Invalid path name
	originalString = "My BadPath (10)"
	expectedOutput = "MyBadPath10"
	methodName = "PathName"

	result := PathName(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Invalid characters
	originalString = "My BadPath (10)[]()!$"
	expectedOutput = "MyBadPath10"

	result = PathName(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Valid Path
	originalString = "My_Folder-Path-123_TEST"
	expectedOutput = "My_Folder-Path-123_TEST"

	result = PathName(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkPathName benchmarks the PathName method
func BenchmarkPathName(b *testing.B) {
	testString := "/This-Path-Name_Works-/"
	for i := 0; i < b.N; i++ {
		_ = PathName(testString)
	}
}

//TestPunctuation tests the punctuation method
func TestPunctuation(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Keep standard punctuation
	originalString = "Mark Mc'Cuban-Host"
	expectedOutput = "Mark Mc'Cuban-Host"
	methodName = "Punctuation"

	result := Punctuation(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Keep periods and commas
	originalString = "Johnny Apple.Seed, Martin"
	expectedOutput = "Johnny Apple.Seed, Martin"

	result = Punctuation(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Keep hashes and exclamation
	originalString = "Does #Not Work!"
	expectedOutput = "Does #Not Work!"

	result = Punctuation(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Keep question marks
	originalString = "Does #Not Work!?"
	expectedOutput = "Does #Not Work!?"

	result = Punctuation(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Keep ampersands
	originalString = "Does #Not Work! & this"
	expectedOutput = "Does #Not Work! & this"

	result = Punctuation(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Keep quotes
	originalString = `[@"Does" 'this' work?@]this`
	expectedOutput = `"Does" 'this' work?this`

	result = Punctuation(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove invalid characters
	originalString = "Does, 123^* Not & Work!?"
	expectedOutput = "Does, 123 Not & Work!?"

	result = Punctuation(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkPunctuation benchmarks the Punctuation method
func BenchmarkPunctuation(b *testing.B) {
	testString := "Does this work? The're doing it?"
	for i := 0; i < b.N; i++ {
		_ = Punctuation(testString)
	}
}

//TestScripts tests the script removal
func TestScripts(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Test removing a script
	originalString = "this <script>$('#something').hide()</script>"
	expectedOutput = "this "
	methodName = "Scripts"

	result := Scripts(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove JS script
	originalString = "this <script type='text/javascript'>$('#something').hide()</script>"
	expectedOutput = "this "

	result = Scripts(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove JS script
	originalString = `this <script type="text/javascript" class="something">$('#something').hide();</script>`
	expectedOutput = "this "

	result = Scripts(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove iframe
	originalString = `this <iframe width="50" class="something"></iframe>`
	expectedOutput = "this "

	result = Scripts(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove embed tag
	originalString = `this <embed width="50" class="something"></embed>`
	expectedOutput = "this "

	result = Scripts(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove object
	originalString = `this <object width="50" class="something"></object>`
	expectedOutput = "this "

	result = Scripts(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkScripts benchmarks the Scripts method
func BenchmarkScripts(b *testing.B) {
	testString := "<script>$(){ var remove='me'; }</script>"
	for i := 0; i < b.N; i++ {
		_ = Scripts(testString)
	}
}

//TestSingleLine test the single line sanitize method
func TestSingleLine(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	methodName = "SingleLine"
	originalString = `Mark
Mc'Cuban-Host`
	expectedOutput = "Mark Mc'Cuban-Host"

	result := SingleLine(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	originalString = `Mark
Mc'Cuban-Host
something else`
	expectedOutput = "Mark Mc'Cuban-Host something else"

	result = SingleLine(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkSingleLine benchmarks the SingleLine method
func BenchmarkSingleLine(b *testing.B) {
	testString := `This line
That Line
Another Line`
	for i := 0; i < b.N; i++ {
		_ = SingleLine(testString)
	}
}

//TestTime tests the time sanitize method
func TestTime(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Just the timestamp, no timezone
	originalString = "t00:00d -EST"
	expectedOutput = "00:00"
	methodName = "Time"

	result := Time(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Just the timestamp, no other characters
	originalString = "t00:00:00d -EST"
	expectedOutput = "00:00:00"

	result = Time(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Just the timestamp, remove everything else
	originalString = "SOMETHING t00:00:00d -EST DAY"
	expectedOutput = "00:00:00"

	result = Time(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkTime benchmarks the Time method
func BenchmarkTime(b *testing.B) {
	testString := "Time is 05:10:23"
	for i := 0; i < b.N; i++ {
		_ = Time(testString)
	}
}

//TestURI tests the URI sanitize method
func TestURI(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Test remove spaces
	originalString = "Test?=weee! &this=that"
	expectedOutput = "Test?=weee&this=that"
	methodName = "URI"

	result := URI(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test removing invalid symbols
	originalString = "Test?=weee! &this=/that/!()*^"
	expectedOutput = "Test?=weee&this=/that/"

	result = URI(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Test valid url
	originalString = "/This/Works/?woot=123&this#page10%"
	expectedOutput = "/This/Works/?woot=123&this#page10%"

	result = URI(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkURI benchmarks the URI method
func BenchmarkURI(b *testing.B) {
	testString := "/Test/This/Url/?param=value"
	for i := 0; i < b.N; i++ {
		_ = XSS(testString)
	}
}

//TestURL tests the URL sanitize method
func TestURL(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Invalid url, remove spaces
	originalString = "Test?=weee! &this=that#works"
	expectedOutput = "Test?=weee&this=that#works"
	methodName = "URL"

	result := URL(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Invalid characters
	originalString = "/this/test?dfsf$"
	expectedOutput = "/this/test?dfsf"

	result = URL(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Invalid characters
	originalString = "https://domain.com/this/test?dfsf$!@()[]{}'<>"
	expectedOutput = "https://domain.com/this/test?dfsf"

	result = URL(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Valid url
	originalString = "https://domain.com/this/test?this=value&another=123%#page"
	expectedOutput = "https://domain.com/this/test?this=value&another=123%#page"

	result = URL(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkURL benchmarks the URL method
func BenchmarkURL(b *testing.B) {
	testString := "/Test/This/Url/?param=value"
	for i := 0; i < b.N; i++ {
		_ = XSS(testString)
	}
}

//TestXML tests the XML sanitize method
func TestXML(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Test basic HTML removal
	methodName = "XML"
	originalString = `<?xml version="1.0" encoding="UTF-8"?><note>Something</note>`
	expectedOutput = "Something"

	result := XML(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Advanced XML removal
	originalString = `<body>This works?</body><title>Something</title>`
	expectedOutput = "This works?Something"

	result = XML(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkXML benchmarks the XML method
func BenchmarkXML(b *testing.B) {
	testString := "<xml>Test This!</xml>"
	for i := 0; i < b.N; i++ {
		_ = XML(testString)
	}
}

//TestXSS tests the XSS sanitize method
func TestXSS(t *testing.T) {
	var (
		expectedOutput string
		methodName     string
		originalString string
	)

	//Remove the script tags
	originalString = "<script>alert('test');</script>"
	expectedOutput = ">alert('test');</"
	methodName = "XSS"

	result := XSS(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove the lt and rt characters
	originalString = "&lt;script&lt;alert('test');&lt;/script&lt;"
	expectedOutput = "scriptalert('test');/script"

	result = XSS(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove javascript:
	originalString = "javascript:alert('test');"
	expectedOutput = "alert('test');"

	result = XSS(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove eval
	originalString = "eval('test');"
	expectedOutput = "'test');"

	result = XSS(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove js
	originalString = "javascript&#58;('test');"
	expectedOutput = "('test');"

	result = XSS(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove char code
	originalString = "fromCharCode('test');"
	expectedOutput = "('test');"

	result = XSS(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}

	//Remove char code, 60 and 62
	originalString = "&#60;&#62;fromCharCode('test');&#62;&#60;"
	expectedOutput = "('test');"

	result = XSS(originalString)
	if result != expectedOutput {
		t.Fatal(methodName, "did not work properly, expected result: [", expectedOutput, "] but received: [", result, "]")
	}
}

//BenchmarkXSS benchmarks the XSS method
func BenchmarkXSS(b *testing.B) {
	testString := "<script>Test This!</script>"
	for i := 0; i < b.N; i++ {
		_ = XSS(testString)
	}
}
