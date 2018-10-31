package structs

import (
	"reflect"
	"strings"

	c "ecards/helper/constant"
)

type Code int

var strs = make(map[reflect.Type]map[int]map[string]string)

// String convert
func (e Code) String(errCode *[]TypeError, customErrorMessage ...string) {
	str := getStr(e)
	strMsg := str["msg"]
	if len(customErrorMessage) > 0 {
		strMsg += " " + strings.Join(customErrorMessage, ", ")
	}

	*errCode = append(*errCode, TypeError{
		Code:    c.PREFIXERRCODE + str["code"],
		Case:    str["case"],
		Message: strMsg,
	})
}

func GetCodeError(code []string, errCode *[]TypeError) {
	// v := reflect.ValueOf(ErrorCode)
	// vt := v.Type()
	// for _, val := range code {
	// 	filter
	// }
}

func getStr(e interface{}) map[string]string {
	t := reflect.TypeOf(e)
	v := int(reflect.ValueOf(e).Int())
	str := strs[t][v]

	return str
}

// ErrorCode Struct
var ErrorCode struct {

	// Common Format Error //
	UnexpectedError          Code `code:"0000" case:"Unexpected error" msg:"Unexpected error"`
	RPCInvalid               Code `code:"0001" case:"Connection RPC Invalid" msg:"Connection RPC Invalid"`
	ServiceNotAllowedToBatch Code `code:"0002" case:"Service Not Allowed to Batch" msg:"Service Not Allowed to Batch"`
	URLFilterInvalid         Code `code:"0003" case:"URL Filter Invalid" msg:"URL Filter Invalid"`

	RequestMalformed   Code `code:"0004" case:"Request malformed / Unacceptable" msg:"Request malformed / Unacceptable"`
	DatabaseError      Code `code:"0005" case:"Database Error" msg:"Database Error"`
	DatabaseConnError  Code `code:"0006" case:"Database Connection Error" msg:"Database Connection Error"`
	FormatError        Code `code:"0007" case:"Format Error" msg:"Format Error"`
	MismatchParamValue Code `code:"0008" case:"Mismatch Parameter Value" msg:"Mismatch Parameter Value"`

	MissingField Code `code:"0009" case:"Missing Field {field}" msg:"{field} is Required"`

	ErorTypeString  Code `code:"0010" case:"Key type is string, but the value isn't" msg:"Field {field} must be string only"`
	ErorTypeNumeric Code `code:"0011" case:"Key type is numeric, but the value isn't" msg:"Field {field} must be numeric only"`
	ErorTypeTime    Code `code:"0012" case:"Key Datetime is (ISO 8601), but the value isn't" msg:"Field {field} must be ISO-8601 only"`

	TokenInvalid        Code `code:"0013" case:"Token Invalid" msg:"Token Invalid"`
	TokenExpired        Code `code:"0014" case:"Token Expired" msg:"Token Expired"`
	TokenGenerateFailed Code `code:"0015" case:"Generate Token Failed" msg:"Generate Token Failed"`
	TokenGenerateDenied Code `code:"0016" case:"Permission Denied" msg:"Permission Denied"`
	// Common Format Error //

	// Start Defined Format Error //
	BranchCodeEmpty             Code `code:"0301" case:"Branch code is blank" msg:"Branch Code is required"`
	BranchCodeCantContainLetter Code `code:"0302" case:"Branch Code Contain Letter" msg:"Branch Code must be numeric only"`
	BranchCodeLength            Code `code:"0303" case:"Branch Code lenght is not 4 char" msg:"Branch Code must be 4 Characters"`
	// End Defined Format Error //
}

type (
	// TypeError ...
	TypeError struct {
		Code    string `json:"errorCode"`
		Case    string `json:"errorCase"`
		Message string `json:"errorDesc"`
	}

	// TypeGRPCError ...
	TypeGRPCError struct {
		Error []TypeError `json:"error"`
	}
	// TypeHTTPError ...
	TypeHTTPError struct {
		Error []TypeError `json:"error"`
	}
)
