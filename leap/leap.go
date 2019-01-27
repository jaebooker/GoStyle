package leap
//create function to test leapyear
func IsLeapYear(year int) bool {
	//check if year is divisible by 400, before proceding
	if (year % 400 == 0) {
		return true
	}
	//check if year is divisible by 4, making sure it is not new century
	if (year % 4 == 0) && (year % 100 != 0) {
		return true
	}
	//if neither is true, return false
	return false
}
