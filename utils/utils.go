package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	mathrand "math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// GetRandomId
// 随机生成ID
func GetRandomId() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

func IsEmpty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)

	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

func IsBlankString(str string) bool {
	if len(str) == 0 {
		return true
	}
	nStr := strings.Replace(str, " ", "", -1)
	if len(nStr) == 0 {
		return true
	} else if nStr == "\r" || nStr == "\n" || nStr == "\r\n" {
		return true
	}
	return false
}

func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

func RandomNumber(length int) string {
	table := [...]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

func RandomString(length int) string {
	mathrand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[mathrand.Intn(len(letters))]
	}
	return string(b)
}

func ParseTimeDur(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	if strings.HasSuffix(d, "h") {
		index := strings.Index(d, "h")
		hour, _ := strconv.Atoi(d[:index])
		return time.Hour * time.Duration(hour), nil
	} else if strings.HasSuffix(d, "d") {
		index := strings.Index(d, "d")
		day, _ := strconv.Atoi(d[:index])
		return time.Hour * 24 * time.Duration(day), nil
	}
	return time.Duration(0), errors.New("配置有误，解析失败，时间信息为: " + d)
}

func CalPageSize(total int64) int64 {
	if (total % 11) != 0 {
		return total/11 + 1
	}
	return total / 11
}
