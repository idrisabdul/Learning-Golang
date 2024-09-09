package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
ConvertToPointer converts a value of a flexible type (FlexibleType) into a pointer to that value.
It accepts any value of type FlexibleType and returns a pointer to that value.
However, there's an issue in the code: it creates pointers to local variables, which will result in incorrect behavior.

Usage Example:
stringPtr := utils.ConvertToPointer("success").(*string)
*/
func ConvertToPointer(value FlexibleType) interface{} {
	switch v := value.(type) {
	case int:
		ptr := &v
		return ptr
	case string:
		ptr := &v
		return ptr
	case float64:
		ptr := &v
		return ptr
	default:
		return nil
	}
}

func ConvertToPointerIfNotEmpty(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func GenerateEncoded(Value string) string {
	// Encrypt changeCode to base64
	encoded := base64.StdEncoding.EncodeToString([]byte(Value))

	return encoded
}

func GenerateDecoded(encodedValue string) (string, error) {
	decodedBytes, _ := base64.StdEncoding.DecodeString(encodedValue)
	/* if err != nil {
		return "", err // Return an error if decoding fails
	} */
	return string(decodedBytes), nil
}

func ParseJSONString(jsonString string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func GenerateNumberEncode(item string) string {
	seed := time.Now().UnixNano()

	r := rand.New(rand.NewSource(seed))

	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letter := make([]rune, 16)
	for i := range letter {
		letter[i] = charset[r.Intn(int(len(charset)))]
	}

	code := GenerateEncoded(string(letter) + item)

	return code
}

func GenerateNumberDecode(item string) (*int, error) {

	code, _ := GenerateDecoded(item)

	r := regexp.MustCompile(`\d+`)

	str := r.FindAllString(code, -1)

	decodedStr, errDecodedStr := strconv.Atoi(str[0])
	if errDecodedStr != nil {
		return nil, errDecodedStr
	}
	return &decodedStr, nil
}

// Get time as string
func ConvertTimeToString(datetime time.Time, rules string) string {

	// Format the time
	if rules == "datetime" {
		formattedTime := datetime.Format("at Jan 2, 2006 15:04:05")
		return formattedTime
	}

	if rules == "default" {
		formattedTime := datetime.Format("2006-01-02 15:04:05")
		return formattedTime
	}

	if rules == "normal" {
		formattedTime := datetime.Format("02-01-2006 15:04:05")
		return formattedTime
	}

	if rules == "fullname" {
		formattedTime := datetime.Format("02 January 2006")
		return formattedTime
	}

	// default value as string "-"
	formatedTime := datetime.Format("-")
	return formatedTime
}

func GetDateTime() (time.Time, error) {
	dateTimeString := "2023-10-22 12:12:02"

	layout := "2006-01-02 15:04:05"

	parsedTime, err := time.Parse(layout, dateTimeString)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}

/*
This function will retun datetime now
if you need time give the paramter (time)
if you need date give the paramter (date)
if you need date and time give the paramter (datetime)
*/
func GetTimeNow(type_date string) string {
	indonesiaLocation, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().Local().In(indonesiaLocation)
	if type_date == "date" {
		dateStr := now.Format("2006-01-02")
		return dateStr
	}

	if type_date == "time" {
		timeStr := now.Format("15:04:05")
		return timeStr

	}

	if type_date == "year" {
		dateStr := now.Format("2006")
		return dateStr
	}

	datetime := now.Format("2006-01-02 15:04:05")
	return datetime

}

func ParsedTimeCanNil(DateTime string) *time.Time {
	if DateTime == "" {
		return nil
	}

	parsedTime1, err1 := time.Parse("2006-01-02 15:04:05", DateTime)
	if err1 != nil {
		parsedTime2, err2 := time.Parse("02-01-2006 15:04:05", DateTime)
		if err2 != nil {
			return nil
		}
		return &parsedTime2
	}

	return &parsedTime1
}

func ParsedTime(DateTime string) time.Time {
	parsedTime1, err1 := time.Parse("2006-01-02 15:04:05", DateTime)
	if err1 != nil {
		parsedTime2, _ := time.Parse("02-01-2006 15:04:05", DateTime)
		return parsedTime2
	}

	return parsedTime1
}

// converting time
func ConvertStringToTime(date string) time.Time {
	var dateFormat string

	// String representation of time with format "yyyy-mm-dd"
	// dateString := "2023-11-10"
	// Choose the appropriate format based on the presence of "T"
	fmt.Println("date:", date)

	if strings.Contains(date, "T") {
		dateFormat = "2006-01-02T15:04:05Z"
	} else {
		dateFormat = "2006-01-02 15:04:05"
	}

	// Parse the input string to time.Time
	parsedTime, err := time.Parse(dateFormat, date)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Return the parsed time
	return parsedTime
}

/*
This function is used for convert Struct to Map by marshal the struct to json,
then unmarshal it into map[string]any. Make sure you define json tags on the struct
before use this function
*/
func ConvertStructToMap(data any) (map[string]any, error) {

	result := make(map[string]any, 0)

	toJson, errToJson := json.Marshal(data)
	if errToJson != nil {
		return nil, errToJson
	}

	if errUnmarshal := json.Unmarshal(toJson, &result); errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return result, nil
}
