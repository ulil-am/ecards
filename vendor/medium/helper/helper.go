package helper

import (
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"math"
	"math/big"
	"math/rand"
	"medium/helper/constant"
	"medium/helper/timetn"
	"medium/structs"
	structsRPC "medium/structs/api/grpc"
	logicStruct "medium/structs/logic"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	js "github.com/json-iterator/go"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/context"
)

var secureDigits = 6

// JS JSON FAST
var JS js.API

// Bm ...
var Bm cache.Cache

func init() {
	bm, err := cache.NewCache("memory", `{"interval":0}`)
	Bm = bm
	if err != nil {
		beego.Warning("error set cache", err)
	}
	// JS = js.ConfigFastest
	JS = js.ConfigCompatibleWithStandardLibrary
}

// HeaderAll to get all header request
func HeaderAll(c *context.Context) string {
	headerAll := make(map[string]string)
	for k, v := range c.Request.Header {
		headerAll[k] = v[0]
	}
	strJSON, err := json.Marshal(headerAll)
	if err != nil {
		CheckErr("error Marshal header request", err)
	}

	return string(strJSON)
}

// GetHeader - GetHeaderParseToStruct
func GetHeader(headerAll string) (header structs.ReqHTTPHeader) {
	err := json.Unmarshal([]byte(headerAll), &header)
	if err != nil {
		CheckErr("failed unmarshal header", err)
	}

	return
}

// ContextStruct - Fill data into this ContextStruct
func ContextStruct(c *context.Context) (contextStruct logicStruct.ContextStruct) {
	headerAll := HeaderAll(c)
	header := GetHeader(headerAll)
	jobID := GetJobID(c)

	contextStruct.HeaderAll = headerAll
	contextStruct.Header = header
	contextStruct.JobID = jobID
	contextStruct.HeaderTracer.HTTPGetHeaderTrace(c)

	return
}

//IsValidUUID ...
func IsValidUUID(uuid string) bool {
	g := `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}` +
		`-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`
	r := regexp.MustCompile(g)
	return r.MatchString(uuid)
}

// SetHeaderRPC ...
func SetHeaderRPC(befMsStr string, reqID string,
	errorHeader structs.TypeGRPCError) []byte {
	nowMs := timetn.Now().UnixNano() / int64(time.Millisecond)

	befMs, err := strconv.ParseInt(befMsStr, 10, 64)
	CheckErr("", err)
	afterMs := nowMs - befMs

	header := structsRPC.TypeHeaderRPC{
		ReqID:       reqID,
		Date:        timetn.Now(),
		ContentType: "application/grpc",
		RoundTrip:   strconv.FormatInt(afterMs, 10),
		Error:       errorHeader,
	}

	jsonByte, err := json.Marshal(header)
	CheckErr("", err)

	return jsonByte
}

// IsNewRequest ...
func IsNewRequest(c *context.Context) bool {
	newRequest := true
	newRequests, err := json.Marshal(c.Input.GetData("new_request"))
	CheckErr("", err)
	newRequestst := strings.Replace(string(newRequests), `"`, ``, -1)
	newRequestInt, err := strconv.Atoi(newRequestst)

	if err != nil {
		beego.Error(err)
	}

	if newRequestInt == 0 {
		newRequest = false
	}

	return newRequest
}

// GetReqID Get x-request-id
func GetReqID(ctx *context.Context) string {
	return ctx.Input.Header("x-request-id")
}

// GetJobID Get job-id
func GetJobID(ctx *context.Context) string {
	jobIDs, err := json.Marshal(ctx.Input.GetData("job-id"))
	CheckErr("Error marshal job-id", err)
	jobID := strings.Replace(string(jobIDs), `"`, ``, -1)

	return jobID
}

// GetMessageID Get message-id
func GetMessageID(ctx *context.Context) string {
	msgIDs, err := json.Marshal(ctx.Input.GetData("message-id"))
	CheckErr("Error marshal message-id", err)
	msgID := strings.Replace(string(msgIDs), `"`, ``, -1)

	return msgID
}

// GetRoundTrip ...
func GetRoundTrip(ctx *context.Context) string {
	intfData, err := json.Marshal(ctx.Input.GetData("x-roundtrip"))
	CheckErr("", err)
	sanitizeMs := strings.Replace(string(intfData), `"`, ``, -1)

	beforeMs, err := strconv.ParseInt(sanitizeMs, 10, 64)
	ms := timetn.Now().UnixNano() / int64(time.Millisecond)
	afterMs := ms - beforeMs
	CheckErr("", err)

	return strconv.FormatInt(afterMs, 10)
}

// GetRoundTripInternal ...
func GetRoundTripInternal(beforeMs int64) string {
	ms := timetn.Now().UnixNano() / int64(time.Millisecond)
	afterMs := ms - beforeMs

	return strconv.FormatInt(afterMs, 10)
}

// GetRqBody Get RqBody
func GetRqBody(ctx *context.Context) []byte {
	var reqData structs.ReqData
	err := JS.Unmarshal(ctx.Input.RequestBody, &reqData)
	CheckErr("Error Marshal Get RqBody", err)

	bodyByte, err := JS.Marshal(reqData.ReqBody)
	CheckErr("Error Marshal Get RqBody 2", err)

	return bodyByte
}

// GetRqBodyRev ...
func GetRqBodyRev(ctx *context.Context,
	errCode *[]structs.TypeError) []byte {
	var reqData structs.ReqData
	err1 := JS.Unmarshal(ctx.Input.RequestBody, &reqData)
	CheckErr("Error Marshal Get RqBody", err1)

	bodyByte, err2 := JS.Marshal(reqData.ReqBody)
	CheckErr("Error Marshal Get RqBody 2", err2)

	if err1 != nil || err2 != nil {
		structs.ErrorCode.RequestMalformed.String(errCode)
	}

	return bodyByte
}

// ConstructJSONHeader ...
func ConstructJSONHeader(ctx *context.Context) structs.JSONHeaderResp {
	URL := ctx.Input.URL()
	URLs := strings.Split(URL, "/")
	version := URLs[1]

	now := timetn.Now()
	JSONHeader := structs.JSONHeaderResp{
		JobID:         GetJobID(ctx),
		MessageID:     GetMessageID(ctx),
		CallRtnSvc:    constant.GOAPP,
		CallRtnSvcVer: version,
		ReqDmn:        beego.AppConfig.String("domain::domain"),
		DateTime:      now,
		AcceptLang:    "en/id",
		AcceptEnc:     "UTF-8",
	}
	return JSONHeader
}

// ConstructHTTPHeader ...
func ConstructHTTPHeader(ctx *context.Context) structs.ResHTTPHeader {
	now := timetn.Now()
	nowISO, err := json.Marshal(now)
	CheckErr("", err)
	nowISOString := strings.Replace(string(nowISO), `"`, ``, -1)

	reqID := GetReqID(ctx)
	jobID := GetJobID(ctx)
	roundtrip := GetRoundTrip(ctx)
	ctnType := "application/json"

	ctx.Output.Header("x-request-id", reqID)
	ctx.Output.Header("x-job-id", jobID)
	ctx.Output.Header("datetime", nowISOString)
	ctx.Output.Header("content-type", ctnType)
	ctx.Output.Header("x-roundtrip", roundtrip)
	ctx.Output.Header("Server", "TN Engineer. We are hiring")

	resHeader := structs.ResHTTPHeader{
		XRequestID:  GetReqID(ctx),
		XJobID:      jobID,
		Datetime:    now,
		ContentType: ctnType,
		XRoundTrip:  roundtrip,
	}

	return resHeader
}

// ToFixedRoundDigits ...
func ToFixedRoundDigits(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	// return float64(round(num*output)) / output
	return float64(int64(num*output+0.5)) / output
}

// FloatToString ...
func FloatToString(inputNum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}

// StringToFloat ...
func StringToFloat(elem string) float64 {
	i, err := strconv.ParseFloat(elem, 64)
	if err != nil {
		beego.Warning("Failed Parse To Float")
		beego.Warning(err)
	}
	return i
}

// ContainsArray ...
func ContainsArray(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// GenJobID Generate Job ID
func GenJobID() string {
	return algorithmJobID()
}

func algorithmJobID() string {
	crc := GenCRC32()
	domain := constant.DOMAIN
	scrnm := getSecureRandomNumber(secureDigits)

	jobID := strings.ToUpper(crc + domain + scrnm)
	dateString := strconv.Itoa(timetn.Now().Year())[2:] +
		filterStrMonthDay(strconv.Itoa(int(timetn.Now().Month()))) +
		filterStrMonthDay(strconv.Itoa(timetn.Now().Day()))

	return dateString + jobID
}

func filterStrMonthDay(strData string) string {
	if len(strData) < 2 {
		return `0` + strData
	}
	return strData
}

// Random ...
func Random(min, max int) string {
	rand.Seed(timetn.Now().Unix())
	randDigits := rand.Intn(max-min) + min
	return strconv.Itoa(randDigits)
}

func getSecureRandomNumber(length int) string {
	token := ""

	secureParam := "0123456789"

	for i := 0; i < length; i++ {
		token += string(secureParam[cryptoRandSecure(int64(len(secureParam)))])
	}
	return token
}

func cryptoRandSecure(max int64) int64 {
	nBig, err := crand.Int(crand.Reader, big.NewInt(max))
	if err != nil {
		beego.Warning(err)
	}
	return nBig.Int64()
}

// GenCRC32 ...
func GenCRC32() string {
	return fmt.Sprintf("%08x", crc32.ChecksumIEEE([]byte(time.Now().String())))
}

// RangeRandomNumber Create Random Account with spesific range
func RangeRandomNumber(min int64, max int64) int64 {
	rand.Seed(timetn.Now().Unix())

	return rand.Int63n(max-min) + min
}

// PathLogFile Create Path Log File
func PathLogFile(pathfile string) string {
	hostname, err := os.Hostname()
	if err != nil {
		beego.Warning("error get hostname")
		return ""
	}
	now := timetn.Now()

	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hours := now.Hour()

	strTime := strconv.Itoa(year) + `-` +
		strconv.Itoa(month) + `-` +
		strconv.Itoa(day) + `-` +
		strconv.Itoa(hours)

	pathLog := pathfile + "/" + hostname + "_" + strTime + "_" +
		constant.GOAPP + ".log"
	return pathLog
}
