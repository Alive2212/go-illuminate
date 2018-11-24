package base

import "os"

//GetEnv is used in config file for get environment var
//from os and fallback to default value if it is empty.
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}




