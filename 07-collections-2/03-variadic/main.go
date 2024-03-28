package main



func DebugLog(args ...string) []string {
	result := []string{"[DEBUG]"}
	result = append(result, args...)
	return result
}

func InfoLog(args ...string) []string {
	result := []string{"[INFO]"}
	result = append(result, args...)
	return result
}

func ErrorLog(args ...string) []string {
	result := []string{"[ERROR]"}
	result = append(result, args...)
	return result
}
