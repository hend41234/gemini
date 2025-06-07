package geminiutilsetc

func Contains(comp string, list []string) bool {
	for _, val := range list {
		if val == comp {
			return true
		}
	}
	return false
}
