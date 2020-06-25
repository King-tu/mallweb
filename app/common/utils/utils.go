package utils

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/pem"
	se "errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/king-tu/mallweb/app/global"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	_ "image/gif"
	_ "image/jpeg"
	"io"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	_ "strconv"
	"strings"
	"unicode"

	"encoding/base64"
	"encoding/json"
	"hash/fnv"
	"reflect"
	"time"

	"github.com/king-tu/mallweb/app/common"
	"github.com/king-tu/mallweb/app/common/errors"
	"github.com/king-tu/mallweb/app/common/i18n"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
)

func AbortWithError(c *gin.Context, err *errors.APIError) {
	returnErr := err

	xRequestId := ""
	xRequestIdObj := c.Value("x-request-id")
	if xRequestIdObj != nil {
		xRequestId = xRequestIdObj.(string)
	}

	lang := "zh-CN"
	localeLang := c.Value("locale-lang")
	if localeLang != nil {
		lang = localeLang.(string)
	}
	if i18n.DefaultLang != lang {
		returnErr = err.Copy4LangWithRequestId(lang, xRequestId)
	} else {
		returnErr.XRequestId = xRequestId
	}

	c.JSON(returnErr.Code, returnErr.Response())
	c.AbortWithError(returnErr.Code, returnErr)
}

func StringHashInt(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}

func HashPassword(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pwd), nil
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		global.Logger.Bg().Error(false, "CompareHashAndPassword err", zap.Error(err))
		return false
	}
	return true
}

func ValidateInRange(ele string, eles []string) bool {
	for _, v := range eles {
		if ele == v {
			return true
		}
	}
	return false
}

func Contains(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

func ValidateMobileNumber(mobile string) *errors.APIError {
	mobilePatten := "^(1)[1-9]\\d{9}$"
	matched, err := regexp.MatchString(mobilePatten, mobile)
	if err != nil {
		return errors.ErrGeneralInternalFault
	}
	if !matched {
		return errors.ErrMobileNumberIncorrect
	}

	return nil
}

func ValidatePassword(passwd string) error {
	// 要求：由数字和字母组成，并且要同时含有数字和字母，且长度要在8-16位之间。
	mobilePatten := "[0-9A-Za-z]{6,16}"
	matched, err := regexp.MatchString(mobilePatten, passwd)
	if err != nil {
		return errors.ErrGeneralInternalFault
	}
	if !matched {
		return se.New("密码格式不正确: 由数字和字母的一种或两种组成, 密码长度要在8-16位之间")
	}

	return nil
}

func ValidateTelephoneNumber(telephone string) *errors.APIError {
	telephonePatten := "^(0)\\d{2,3}-\\d{7,8}$"
	matched, err := regexp.MatchString(telephonePatten, telephone)
	if err != nil {
		return errors.ErrGeneralInternalFault
	}
	if !matched {
		return errors.ErrTelephoneNumberIncorrect
	}

	return nil
}

func ValidateEmail(email string) bool {
	rxEmail := regexp.MustCompile("^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")

	return rxEmail.MatchString(email)
}

func BuildRDBQuery(query map[string]map[string]interface{}) (string, []interface{}, *errors.APIError) {
	var queryString string = ""
	values := make([]interface{}, 0)
	i := 0
	for k, v := range query {
		if _, ok := v["op"]; !ok {
			return "", nil, errors.ErrGeneralInternalFault
		}

		if _, ok := v["value"]; !ok {
			return "", nil, errors.ErrGeneralInternalFault
		}

		//construct the ? part
		var uncertain string
		switch v["op"].(string) {
		case "in":
			uncertain = k + v["op"].(string) + " in (?)"
		case "between":
			uncertain = k + v["op"].(string) + " between ? and ?"
		default:
			uncertain = k + v["op"].(string) + " ? "
		}
		if i == 0 {
			queryString = uncertain
		} else {
			queryString = queryString + " and " + uncertain
		}

		if reflect.TypeOf(v["value"]).Kind().String() == "slice" {
			eles := v["value"].([]interface{})
			values = append(values, eles...)
		}
	}
	return queryString, values, nil
}

func FindNextElement(current string, set []string) string {
	for index, ele := range set {
		if current == ele && index < len(set)-1 {
			return set[index+1]
		}
	}
	return "!NotFound"
}

func ReplaceAllPlaceHolders(target string, placeHolders map[string]string) string {
	for k, v := range placeHolders {
		target = strings.Replace(target, k, v, -1)
	}
	return target
}

func GetIntegerRangeByLength(length int) (int64, int64, *errors.APIError) {
	if length <= 0 || length > 18 {
		return 0, 0, errors.ErrGeneralInvalidParam
	}

	var buffer bytes.Buffer
	for index := 0; index < length; index++ {
		buffer.WriteString("9")
	}
	abs := buffer.String()

	min := "-" + abs
	minInt64, err := strconv.ParseInt(min, 10, 64)
	if err != nil {
		return 0, 0, errors.ErrGeneralInvalidParam
	}
	maxInt64, err := strconv.ParseInt(abs, 10, 64)
	if err != nil {
		return 0, 0, errors.ErrGeneralInvalidParam
	}
	return minInt64, maxInt64, nil
}

func AddTimeByUnit(t time.Time, timeUnit string, delta int) time.Time {
	switch timeUnit {
	case common.TIME_YEAR:
		return t.AddDate(delta, 0, 0)
	case common.TIME_MONTH:
		return t.AddDate(0, delta, 0)
	case common.TIME_DAY:
		return t.AddDate(0, 0, delta)
	case common.TIME_HOUR:
		return t.Add(time.Duration(delta) * time.Hour)
	case common.TIME_MINUTE:
		return t.Add(time.Duration(delta) * time.Minute)
	case common.TIME_SECOND:
		return t.Add(time.Duration(delta) * time.Second)
	}

	return t
}

func GetValidCountry() []string {
	return []string{common.COUNTRY_CHINA, common.COUNTRY_USA}
}

// Get all the updated fields from the http request input struct
// @param arg: should be a struct type
// All the fields of the struct should be of pointer type
func GetUpdatedFields(arg interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	r_value := reflect.Indirect(reflect.ValueOf(arg))
	r_type := r_value.Type()
	//only handle the struct type
	if r_type.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < r_type.NumField(); i++ {
		field := r_type.Field(i)
		name := r_type.Field(i).Name

		fieldValue := r_value.FieldByName(name)

		_kind := fieldValue.Type().Kind()

		//only handle all the pointer field
		if _kind == reflect.Ptr && _kind == reflect.Slice && _kind == reflect.Map {
			if !r_value.FieldByName(name).IsNil() {
				if fieldValue.Type().Kind() == reflect.Slice {
					result[field.Tag.Get("json")] = r_value.FieldByName(name).Interface()
				} else {
					result[field.Tag.Get("json")] = r_value.FieldByName(name).Elem().Interface()
				}
			}
		} else {
			result[field.Tag.Get("json")] = r_value.FieldByName(name).Interface()
		}
	}
	return result
}

var (
	defaultRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Creates a random string based on a variety of options, using
// supplied source of randomness.
//
// If start and end are both 0, start and end are set
// to ' ' and 'z', the ASCII printable
// characters, will be used, unless letters and numbers are both
// false, in which case, start and end are set to 0 and math.MaxInt32.
//
// If set is not nil, characters between start and end are chosen.
//
// This method accepts a user-supplied rand.Rand
// instance to use as a source of randomness. By seeding a single
// rand.Rand instance with a fixed seed and using it for each call,
// the same random sequence of strings can be generated repeatedly
// and predictably.
func RandomSpec0(count uint, start, end int, letters, numbers bool, chars []rune, rand *rand.Rand) string {
	if count == 0 {
		return ""
	}
	if start == 0 && end == 0 {
		end = 'z' + 1
		start = ' '
		if !letters && !numbers {
			start = 0
			end = math.MaxInt32
		}
	}
	buffer := make([]rune, count)
	gap := end - start
	for count != 0 {
		count--
		var ch rune
		if len(chars) == 0 {
			ch = rune(rand.Intn(gap) + start)
		} else {
			ch = chars[rand.Intn(gap)+start]
		}
		if letters && ((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) || numbers && (ch >= '0' && ch <= '9') || (!letters && !numbers) {
			if ch >= rune(56320) && ch <= rune(57343) {
				if count == 0 {
					count++
				} else {
					buffer[count] = ch
					count--
					buffer[count] = rune(55296 + rand.Intn(128))
				}
			} else if ch >= rune(55296) && ch <= rune(56191) {
				if count == 0 {
					count++
				} else {
					// high surrogate, insert low surrogate before putting it in
					buffer[count] = rune(56320 + rand.Intn(128))
					count--
					buffer[count] = ch
				}
			} else if ch >= rune(56192) && ch <= rune(56319) {
				// private high surrogate, no effing clue, so skip it
				count++
			} else {
				buffer[count] = ch
			}
		} else {
			count++
		}
	}

	return string(buffer)
}

// Creates a random string whose length is the number of characters specified.
//
// Characters will be chosen from the set of alpha-numeric
// characters as indicated by the arguments.
//
// Param count - the length of random string to create
// Param start - the position in set of chars to start at
// Param end   - the position in set of chars to end before
// Param letters - if true, generated string will include
//                 alphabetic characters
// Param numbers - if true, generated string will include
//                 numeric characters
func RandomSpec1(count uint, start, end int, letters, numbers bool) string {
	return RandomSpec0(count, start, end, letters, numbers, nil, defaultRand)
}

// Creates a random string whose length is the number of characters specified.
//
// Characters will be chosen from the set of alpha-numeric
// characters as indicated by the arguments.
//
// Param count - the length of random string to create
// Param letters - if true, generated string will include
//                 alphabetic characters
// Param numbers - if true, generated string will include
//                 numeric characters
func RandomAlphaOrNumeric(count uint, letters, numbers bool) string {
	return RandomSpec1(count, 0, 0, letters, numbers)
}

func RandomString(count uint) string {
	return RandomAlphaOrNumeric(count, false, false)
}

func RandomStringSpec0(count uint, set []rune) string {
	return RandomSpec0(count, 0, len(set)-1, false, false, set, defaultRand)
}

func RandomStringSpec1(count uint, set string) string {
	return RandomStringSpec0(count, []rune(set))
}

// Creates a random string whose length is the number of characters
// specified.
//
// Characters will be chosen from the set of characters whose
// ASCII value is between 32 and 126 (inclusive).
func RandomAscii(count uint) string {
	return RandomSpec1(count, 32, 127, false, false)
}

// Creates a random string whose length is the number of characters specified.
// Characters will be chosen from the set of alphabetic characters.
func RandomAlphabetic(count uint) string {
	return RandomAlphaOrNumeric(count, true, false)
}

// Creates a random string whose length is the number of characters specified.
// Characters will be chosen from the set of alpha-numeric characters.
func RandomAlphanumeric(count uint) string {
	return RandomAlphaOrNumeric(count, true, true)
}

// Creates a random string whose length is the number of characters specified.
// Characters will be chosen from the set of numeric characters.
func RandomNumeric(count uint) string {
	return RandomAlphaOrNumeric(count, false, true)
}

var coder = base64.NewEncoding(common.BASE64_TABLE)

func Base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, *errors.APIError) {
	encode, err := coder.DecodeString(string(src))
	if err != nil {
		global.Logger.Bg().Error(false, "Base64 decode failed", zap.Error(err))
		return nil, errors.ErrGeneralInternalFault
	}
	return encode, nil
}

func QueryParametersFormat(data map[string]interface{}, obj interface{}) (map[string]interface{}, *errors.APIError) {
	Data := map[string]interface{}{}
	types := reflect.TypeOf(obj).Elem()

	fieldNum := types.NumField()
	for i := 0; i < fieldNum; i++ {
		tag := types.Field(i).Tag.Get("json")
		mapValue, exist := data[tag]
		if !exist {
			continue
		}
		switch types.Field(i).Type.Kind() {
		case reflect.Bool:
			s, err := strconv.ParseBool(mapValue.(string))
			if err != nil {
				return nil, errors.ErrQueryParameterFault
			}
			Data[tag] = s
		case reflect.Int:
			s, err := strconv.ParseInt(mapValue.(string), 10, 32)
			if err != nil {
				return nil, errors.ErrQueryParameterFault
			}
			Data[tag] = s
		case reflect.String:
			Data[tag] = mapValue.(string)
		}
	}

	return Data, nil
}

func RemoveDuplicateStringMember(members []string) []string {
	m := make(map[string]int)

	for _, member := range members {
		if count, ok := m[member]; ok {
			m[member] = count + 1
		} else {
			m[member] = 1
		}
	}

	ret := make([]string, len(m))
	index := 0
	for k, _ := range m {
		ret[index] = k
		index = index + 1
	}

	return ret
}

func Struct2Map(model interface{}) (map[string]interface{}, *errors.APIError) {
	bytes, jsonErr := json.Marshal(model)
	if jsonErr != nil {
		global.Logger.Bg().Error(false, "json marshal error:", zap.Error(jsonErr))
		return nil, errors.ErrGeneralInternalFault
	}
	result := make(map[string]interface{})
	jsonErr = json.Unmarshal(bytes, &result)
	if jsonErr != nil {
		global.Logger.Bg().Error(false, "json marshal error:", zap.Error(jsonErr))
		return nil, errors.ErrGeneralInternalFault
	}
	return result, nil

}

func Compare(Src, Dst []string) []string {
	m := make(map[string]int)
	for _, dst := range Dst {
		m[dst]++
	}

	var ret []string
	for _, src := range Src {
		if m[src] > 0 {
			continue
		}
		ret = append(ret, src)
	}

	return ret
}

func CompareInterfaceArray(Src, Dst []interface{}, exclude interface{}) []interface{} {
	m := make(map[interface{}]int)
	for _, dst := range Dst {
		m[dst]++
	}

	var ret []interface{}
	for _, src := range Src {
		if src == exclude || m[src] > 0 {
			continue
		}
		ret = append(ret, src)
	}

	return ret
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
func DeleteDuplicate(Src []string) []string {
	cleaned := []string{}
	for _, value := range Src {
		if !stringInSlice(value, cleaned) {
			cleaned = append(cleaned, value)
		}
	}

	return cleaned
}

func GetStructTags(data interface{}, tagKey string) ([]string, *errors.APIError) {
	var ret []string
	s := reflect.TypeOf(data).Elem()
	for i := 0; i < s.NumField(); i++ {
		tagValue := s.Field(i).Tag.Get(tagKey)
		ret = append(ret, tagValue)
	}

	return ret, nil
}

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}

func TrimAllSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func FormatFloat(num float64, prec int) float64 {
	formatFloat := strconv.FormatFloat(num, 'f', prec, 64)
	retFloat, _ := strconv.ParseFloat(formatFloat, 64)
	return retFloat
}
func FormatFloatFor32(num float32, prec int) float32 {
	_num := float64(num)
	nStr := fmt.Sprintf("%v", num)
	if strings.Contains(nStr, ".") {
		_num += 0.0001 //补充丢失的高精度数，比如1.5转换成float64后可能变成1.499999888
	}
	formatFloat := strconv.FormatFloat(_num, 'f', prec, 64)
	retFloat, _ := strconv.ParseFloat(formatFloat, 64)
	return float32(retFloat)
}
func TrimDashInId(uuid string) string {
	return strings.Replace(uuid, "-", "", -1)
}

func GetRandomNumberId() string {
	return fmt.Sprintf("%d%d", time.Now().UnixNano()/1000, rand.Intn(1000))
}

func GetBankCode(bankName string) string {
	bankCodeMap := map[string]string{
		"中国工商银行": "01020000",
		"中国农业银行": "01030000",
		"招商银行":   "03080000",
		"中国建设银行": "01040000",
		"交通银行":   "03010000",
		"光大银行":   "03030000",
		"中国民生银行": "03050000",
		"北京银行":   "04031000",
		"中信银行":   "03020000",
	}

	bCode := ""

	for k, v := range bankCodeMap {
		if strings.Index(bankName, k) != -1 {
			bCode = v
			break
		}
	}

	return bCode
}

func GetUUID() string {
	return uuid.NewV1().String()
}

func ExtractUpdates(arg interface{}, fields []string) map[string]interface{} {
	result := make(map[string]interface{})
	r_value := reflect.Indirect(reflect.ValueOf(arg))
	r_type := r_value.Type()
	if r_type.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < r_type.NumField(); i++ {
		field := r_type.Field(i)
		name := r_type.Field(i).Name
		jsonTagName := field.Tag.Get("json")

		if !Contains(jsonTagName, fields) {
			continue
		}
		if !r_value.FieldByName(name).IsValid() {
			continue
		}

		fieldValue := r_value.FieldByName(name)
		_kind := fieldValue.Type().Kind()
		if _kind == reflect.Ptr {
			result[field.Tag.Get("json")] = r_value.FieldByName(name).Elem().Interface()
		} else {

			result[field.Tag.Get("json")] = r_value.FieldByName(name).Interface()
		}
	}
	return result
}

func initializeStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			if f.IsNil() {
				f.Set(reflect.MakeMap(ft.Type))
			}
		case reflect.Slice:
			if f.IsNil() {
				f.Set(reflect.MakeSlice(ft.Type, 0, 0))
			} else {
				for i := 0; i < f.Len(); i++ {
					_v := f.Index(i)
					if _v.Type().Kind() == reflect.Ptr && _v.Type().Elem().Kind() == reflect.Struct {
						initializeStruct(_v.Type().Elem(), _v.Elem())
					}
				}
			}
		case reflect.Chan:
			if f.IsNil() {
				f.Set(reflect.MakeChan(ft.Type, 0))
			}
		case reflect.Struct:
			initializeStruct(ft.Type, f)
		case reflect.Ptr:
			if f.IsNil() {
				fv := reflect.New(ft.Type.Elem())
				if ft.Type.Elem().Kind() == reflect.Struct {
					initializeStruct(ft.Type.Elem(), fv.Elem())
				}
				f.Set(fv)
			} else {
				initializeStruct(ft.Type.Elem(), f.Elem())
			}
		default:
		}
	}
}

// To initialize the nil in the struct
// t should be a pointer
func RefineNil(t interface{}) {
	initializeStruct(reflect.TypeOf(t).Elem(), reflect.ValueOf(t).Elem())
}

func Map2Xml(mReq map[string]interface{}) (xml string) {
	sb := bytes.Buffer{}
	sb.WriteString("<xml>")
	for k, v := range mReq {
		str := fmt.Sprintf("%v", v)
		sb.WriteString("<" + k + ">" + str + "</" + k + ">")
	}
	sb.WriteString("</xml>")

	return sb.String()
}

//微信支付计算签名的函数
func WxPayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings)) //
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	zap.L().Sugar().Errorf("signStrings:%v, upperSign:%v", signStrings, upperSign)
	return upperSign
}
func CalcSign(mReq map[string]interface{}, key string) (sign string) {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings)) //
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	zap.L().Sugar().Debugf("signStrings:%v, upperSign:%v", signStrings, upperSign)
	return upperSign
}

//微信支付签名验证函数
func WxpayVerifySign(needVerifyM map[string]interface{}, sign, secretKey string) bool {
	signCalc := WxPayCalcSign(needVerifyM, secretKey)
	if sign == signCalc {
		return true
	}

	return false
}

func MybankCalcSrc(mReq map[string]interface{}) string {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//remove last "&"
	signStrings = signStrings[0 : len(signStrings)-1]
	return signStrings
}

func fromBase64(s []byte) (buf []byte, err error) {
	buflen := base64.StdEncoding.DecodedLen(len(s))
	buf = make([]byte, buflen)
	n, err := base64.StdEncoding.Decode(buf, s)
	buf = buf[:n]
	return
}

func MybankVerifySign(src map[string]interface{}, signature, publicKey string) bool {

	format := `-----BEGIN PUBLIC KEY-----
%s
-----END PUBLIC KEY-----`
	pubKeyPEM := fmt.Sprintf(format, publicKey)
	zap.L().Sugar().Debugf("pubKeyPEM:%v,publicKey:%v", pubKeyPEM, publicKey)
	PEMBlock, _ := pem.Decode([]byte(pubKeyPEM))
	if PEMBlock == nil {
		return false
	}
	pubKey, err := x509.ParsePKIXPublicKey(PEMBlock.Bytes)
	if err != nil {
		return false
	}

	//compute the sha1
	h := sha1.New()
	data := MybankCalcSrc(src)
	h.Write([]byte(data))

	zap.L().Sugar().Debugf("data:%v,src:%v", data, src)
	//sign base64 decode
	sig, err := fromBase64([]byte(signature))
	if err != nil {
		return false
	}
	zap.L().Sugar().Debugf("signature:%v, sig:%v", signature, string(sig))
	//read the signature from stdin
	zap.L().Sugar().Debugf("crypto.SHA256:%v, src len:%v", crypto.SHA1.Size(), len(h.Sum(nil)))
	if err := rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA1, h.Sum(nil), sig); err != nil {
		zap.L().Sugar().Debugf("verify failed,err:%v", zap.Error(err))
		return false
	}
	zap.L().Debug("verify success")
	return true
}

// ZipFiles compresses one or many files into a single zip archive file
func ZipFiles(filename string, files []string) *errors.APIError {

	newfile, err := os.Create(filename)
	if err != nil {
		return errors.ErrGeneralInternalFault
	}
	defer newfile.Close()

	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		zap.L().Sugar().Debugf("zip file %s", file)

		zipfile, err := os.Open(file)
		if err != nil {
			return errors.ErrGeneralInternalFault
		}
		defer zipfile.Close()

		// Get the file information
		info, err := zipfile.Stat()
		if err != nil {
			return errors.ErrGeneralInternalFault
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return errors.ErrGeneralInternalFault
		}
		// Change to deflate to gain better compression
		// see http://golang.org/pkg/archive/zip/#pkg-constants
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return errors.ErrGeneralInternalFault
		}
		_, err = io.Copy(writer, zipfile)
		if err != nil {
			return errors.ErrGeneralInternalFault
		}
	}

	return nil
}

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {
			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return filenames, err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return filenames, err
			}
		}
	}

	return filenames, nil
}

/*// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
// dir contains gd232 chars
func Unzipgb232(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		utf8path,err:=gbtoutf8.ConvertString(fpath,"GB2312","utf-8")
		filenames = append(filenames, utf8path)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(utf8path, os.ModePerm)

		} else {
			// Make File
			if err = os.MkdirAll(filepath.Dir(utf8path), os.ModePerm); err != nil {
				return filenames, err
			}

			outFile, err := os.OpenFile(utf8path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return filenames, err
			}
		}
	}

	return filenames, nil
}*/
func ValidateMobile(mobileNum string) bool {
	regular := "^((1[3,5,6,7,8,9][0-9])|(14[5,7])|(17[0,3,6,7,8])|(19[7]))\\d{8}$"
	reg := regexp.MustCompile(regular)

	return reg.MatchString(mobileNum)
}

func byte2int(x byte) byte {
	if x == 88 {
		return 'X'
	}
	return (x - 48) // 'X' - 48 = 40;
}

func checkId(id [17]byte) int {
	arry := make([]int, 17)

	for index, value := range id {
		arry[index], _ = strconv.Atoi(string(value))
	}

	var res int
	var wi [17]int = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	for i := 0; i < 17; i++ {
		res += arry[i] * wi[i]
	}

	return (res % 11)
}

func verifyId(verify int, idValue byte) bool {
	var temp byte
	var i int
	a18 := [11]byte{1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2}

	for i = 0; i < 11; i++ {
		if i == verify {
			temp = a18[i]
			break
		}
	}
	if temp == idValue {
		return true
	}

	return false
}

/*
	idCardString 如果有'x'会转成'X'
*/
func ValidateIdCard(idCardString string) bool {
	//to upper
	idCardString = strings.ToUpper(idCardString)
	if len(idCardString) != 18 {
		return false
	}

	var idCard [18]byte // 'X' == byte(88)， 'X'在byte中表示为88
	var idCardCopy [17]byte

	if len(idCardString) < 18 {
		return false
	}

	// 将字符串，转换成[]byte,并保存到idCard[]数组当中
	for k, v := range []byte(idCardString) {
		idCard[k] = byte(v)
	}
	//复制idCard[18]前17位元素到idCardCopy[]数组当中
	for j := 0; j < 17; j++ {
		idCardCopy[j] = idCard[j]
	}

	return verifyId(checkId(idCardCopy), byte2int(idCard[17]))
}

func JSONMarshal(v interface{}, unescape bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if unescape {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

func AgeByBirthDate(birthday time.Time) int {
	now := time.Now()
	years := now.Year() - birthday.Year()
	if now.YearDay() < birthday.YearDay() {
		years--
	}
	return years
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//创建目录
//如果已经存在,则不创建; 如果已存在,则清空内容
func MakePathAndClear(path string) (bool, error) {
	exist, err := PathExists(path)
	if exist {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println("delete failed![%v]\n", err)
			return false, err
		}
	}

	// 创建文件夹
	_, err = exec.Command("/bin/sh", "-c", fmt.Sprintf(`mkdir -p %s`, path)).Output()
	if err != nil {
		fmt.Printf("mkdir failed![%v]\n", err)
		return false, err
	}
	fmt.Printf("mkdir success!\n")
	return true, nil
}

func MakePathMany(path string) (bool, error) {
	exist, err := PathExists(path)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return false, err
	}
	if exist {
		fmt.Printf("has dir![%v]\n", path)
		return true, nil
	}

	// 创建文件夹
	_, err = exec.Command("/bin/sh", "-c", fmt.Sprintf(`mkdir -p %s`, path)).Output()
	if err != nil {
		fmt.Printf("mkdir failed![%v]\n", err)
		return false, err
	}
	fmt.Printf("mkdir success!\n")
	return true, nil
}

//创建目录
//如果已经存在,则不创建
func MakePath(path string) (bool, error) {
	exist, err := PathExists(path)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return false, err
	}
	if exist {
		fmt.Printf("has dir![%v]\n", path)
		return true, nil
	}

	// 创建文件夹
	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
		fmt.Printf("mkdir failed![%v]\n", err)
		return false, err
	}
	fmt.Printf("mkdir success!\n")
	return true, nil
}

// 手机号打码
func MaskMobile(mobile string, maskNumber int) string {
	offset := 3
	if len(mobile) > 11 {
		offset += 2
	}
	maskNumber += offset

	ret := mobile
	if len(mobile) > maskNumber {
		_cardNo := make([]rune, len(mobile))
		for index, _c := range mobile {
			if index >= offset && index < maskNumber {
				_cardNo[index] = '*'
				continue
			}
			_cardNo[index] = _c
		}

		ret = string(_cardNo)
	}

	return ret
}

// 去除头尾空白字符和不可见字符
func TrimSpace(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r) || r == '\u202d' || r == '\u202c'
	})
}

//ValidateChinese 检验字符串是不是中文
func ValidateChinese(s string) bool {
	for _, r := range s {
		if !unicode.Is(unicode.Scripts["Han"], r) {
			return false
		}
	}

	return true
}

//ValidateNum 检验字符串是否全是数据
func ValidateNum(s string) bool {
	zap.L().Sugar().Debugf("字符串数据:%s", s)
	for _, r := range s {
		zap.L().Sugar().Debugf("数字%v", r)
		if r < 48 || r > 57 {
			return false
		}
	}

	return true
}

// DumpStructToMap 将数据库的结构体解析成map并去掉id,created_at updated_at字段
func DumpStructToMap(dbStruct interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	data, err := json.Marshal(dbStruct)
	if err != nil {
		return nil, errors.NewError(errors.ErrGeneralInternalFault, "DumpStructToMap failed to json.Marshal, err:%v", err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.NewError(errors.ErrGeneralInternalFault, "DumpStructToMap failed to json.Unmarshal, err:%v", err)
	}

	/**
	 * 去掉id, created_at updated_at字段
	 */
	deleteField := []string{"id", "created_at", "updated_at", "version"}
	for _, field := range deleteField {
		delete(result, field)
	}

	return result, nil
}

//DumpStructToMapForInsert 同上，只删除created_at和updated_at
func DumpStructToMapForInsert(dbStruct interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	data, err := json.Marshal(dbStruct)
	if err != nil {
		return nil, errors.NewError(errors.ErrGeneralInternalFault, "DumpStructToMap failed to json.Marshal, err:%v", err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.NewError(errors.ErrGeneralInternalFault, "DumpStructToMap failed to json.Unmarshal, err:%v", err)
	}

	/**
	 * 去掉id, created_at updated_at字段
	 */
	deleteField := []string{"created_at", "updated_at"}
	for _, field := range deleteField {
		delete(result, field)
	}

	return result, nil
}

//StringSliceContain ...
func StringSliceContain(slice []string, sub string) bool {
	for _, item := range slice {
		if item == sub {
			return true
		}
	}
	return false
}

// get the number substring
func GetNumSubStr(s string) string {
	reg, _ := regexp.Compile("[0-9]+")
	return reg.FindString(s)
}

//判断文件或文件夹子是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//CalculateApprociateGroutine ...
func CalculateApprociateGroutine(maxNum, numPerG int) int {
	a, b := maxNum/numPerG, maxNum%numPerG

	if b == 0 {
		return a
	} else {
		return a + 1
	}
}

func Struct2MapV2(v interface{}) map[string]interface{} {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Struct {
		return nil
	}

	rt := reflect.TypeOf(v)
	count := rv.NumField()

	m := map[string]interface{}{}
	for i := 0; i < count; i++ {
		fieldType := rt.Field(i)

		mKey := fieldType.Tag.Get("json")
		if mKey == "" {
			mKey = fieldType.Name
		}

		m[mKey] = rv.Field(i).Interface()
	}

	return m
}

func MergeMap2Struct(v interface{}, mv map[string]interface{}) {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return
	}

	rv := reflect.ValueOf(v).Elem()
	if rv.Kind() != reflect.Struct {
		return
	}

	rt := rv.Type()
	count := rv.NumField()

	fm := make(map[string]string)
	for i := 0; i < count; i++ {
		fieldType := rt.Field(i)

		mKey := fieldType.Tag.Get("json")
		if mKey == "" {
			mKey = fieldType.Name
		}
		fm[mKey] = fieldType.Name
	}

	for k, v := range fm {
		_v := rv.FieldByName(v)
		if !_v.IsValid() || mv[k] == nil {
			continue
		}

		_v.Set(reflect.ValueOf(mv[k]))
	}
}

/****** add by kingtu *******/
func SendSms(phoneNum, msg string) error {
	//客户端
	client, err := sdk.NewClientWithAccessKey(global.Config.SmsConfig.RegionId, global.Config.SmsConfig.AccessKeyId, global.Config.SmsConfig.AccessKeySecret)
	if err != nil {
		return err
	}
	//请求包体
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = global.Config.SmsConfig.RegionId
	//用户模版
	request.QueryParams["PhoneNumbers"] = phoneNum
	//短信签名
	request.QueryParams["SignName"] = global.Config.SmsConfig.SignName
	//短信模板 id
	request.QueryParams["TemplateCode"] = global.Config.SmsConfig.TemplateCode

	//短信模板变量对应的实际值，JSON格式。
	//如果JSON中需要带换行符，请参照标准的JSON协议处理。
	//"code" 为短信模版中的 变量
	request.QueryParams["TemplateParam"] = `{"code":` + msg + `}`
	//发送
	if _, err := client.ProcessCommonRequest(request); err != nil {
		return err
	}
	//调用接口之后的响应信息
	return nil
}

func GetGRPCService(srvName string) service.Service {
	// 配置注册中心
	etcdRegistry := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{global.Config.ETCD.Addr}
	})

	return grpc.NewService(
		service.Name(srvName),
		service.Version(global.Config.System.Version),
		service.Registry(etcdRegistry),
		//TTL指定一次注册在注册中心的有效期，过期后便删除
		service.RegisterTTL(time.Second*time.Duration(global.Config.ETCD.RegisterTTL)),
		//间隔时间注册则是定时向注册中心重新注册以保证服务仍在线
		service.RegisterInterval(time.Second*time.Duration(global.Config.ETCD.RegisterInterval)),
		service.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
}

/************ end ****************/
