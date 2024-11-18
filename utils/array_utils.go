package utils

func StringFoundInArray(ary []string, stringToSearch string) bool {
	for _, item := range ary {
		if item == stringToSearch {
			return true
		}
	}
	return false
}
