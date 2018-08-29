package util

func CheckErrorRobot(err error) bool {
	if err != nil {
		return true
	}

	return false
}
