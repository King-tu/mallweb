package errors

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/king-tu/mallweb/common/i18n"
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

type APIError struct {
	Code       int                 `json:"code"`
	Message    string              `json:"message"`
	Title      string              `json:"title"`
	Errors     map[string][]string `json:"errors"`
	XRequestId string              `json:"x_request_id"`
}

func NewAPIError(code int, title string, message string) *APIError {
	return &APIError{
		Code:    code,
		Title:   title,
		Message: message,
		Errors:  map[string][]string{},
	}
}

func NewFullAPIError(code int, title string, message string, errors map[string][]string, xRequestId string) *APIError {
	return &APIError{
		Code:       code,
		Title:      title,
		Message:    message,
		Errors:     errors,
		XRequestId: xRequestId,
	}
}

func SubError(apiError *APIError, msg string) *APIError {
	newErr := *apiError
	newErr.Message = msg
	return &newErr
}

func ConvertFrom(message string) *APIError {
	return NewAPIError(http.StatusInternalServerError, "GENERAL_INTERNAL_SERVER_ERROR", message)
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	}
	return fmt.Sprintf("%s:%d", file, line)
}

//NewError 将具体的错误内容包进通用的错误结构体中,并返回新的变量地址
func NewError(apiError *APIError, format string, a ...interface{}) *APIError {
	newErr := *apiError
	newErr.Errors = map[string][]string{
		"err_information": []string{
			fmt.Sprintf(format, a...),
		},
		"file_no": []string{
			fileInfo(2),
		},
	}

	return &newErr
}

//Newerror 将具体的错误内容包进通用的错误结构体中,并返回新的变量地址
//func Newerror(err error, format string, a ...interface{}) *APIError {
//	return NewError(ConvertFrom(err), format, a...)
//}

var (
	ErrPanic                        = NewAPIError(http.StatusInternalServerError, "ERR_PANIC", "系统出错")
	ErrRetry                        = NewAPIError(http.StatusInternalServerError, "ERR_RETRY", "内部错误，请稍后重试")
	ErrIndexConflict                = NewAPIError(http.StatusBadRequest, "ERR_INDEX_CONFLIT", "唯一索引或主键冲突")
	ErrNoRowsFind                   = NewAPIError(http.StatusNotFound, "INVALID_INPUT_PARAMS", "sql: no rows in result set")
	ErrLimitInvalid                 = NewAPIError(http.StatusBadRequest, "INVALID_INPUT_PARAMS", "limit is invalid")
	ErrOffsetInvalid                = NewAPIError(http.StatusBadRequest, "INVALID_INPUT_PARAMS", "offet is invalid")
	ErrGlobalAccessMissingInScope   = NewAPIError(http.StatusBadRequest, "GLOBAL_ACCESS_MISSING_IN_SCOPE", "scope里没有global_access")
	ErrGlobalAccessIncorrectInScope = NewAPIError(http.StatusBadRequest, "GLOBAL_ACCESS_INCORRECT_IN_SCOPE", "scope里global_access不正确")
	ErrTenantMissingInScope         = NewAPIError(http.StatusBadRequest, "TENANT_MISSING_IN_SCOPE", "scope里没有指定租户")
	ErrResourceIdMissing            = NewAPIError(http.StatusInternalServerError, "RESOURCE_ID_MISSING", "资源ID未提供")
	ErrResourceConflict             = NewAPIError(http.StatusConflict, "RESOURCE_CONFLICT", "资源冲突")
	ErrResourceNoUpdate             = NewAPIError(http.StatusBadRequest, "RESOURCE_NO_UPDATE", "资源没有可以更新的地方")

	ErrRequestInputReadError = NewAPIError(http.StatusInternalServerError, "REQUEST_INPUT_READ_ERROR", "读取请求输入发生错误")

	ErrRequestBodyEmpty = NewAPIError(http.StatusBadRequest, "REQUEST_BODY_EMPTY", "没有提供请求输入")

	ErrJsonParseError = NewAPIError(http.StatusBadRequest, "JSON_PARSE_ERROR", "请求输入JSON不能被解析")

	ErrInvalidInputParams = NewAPIError(http.StatusBadRequest, "INVALID_INPUT_PARAMS", "请求输入存在不合法的参数")

	ErrAuthenticationFail = NewAPIError(http.StatusUnauthorized, "AUTHORIZE_FAIL", "用户认证授权失败")

	ErrGeneralInternalFault = NewAPIError(http.StatusInternalServerError, "GENERAL_INTERNAL_SERVER_ERROR", "发生内部服务错误")
	ErrInvalidTemplate      = NewAPIError(http.StatusInternalServerError, "GENERAL_INTERNAL_SERVER_ERROR", "该模板需要重新更新")

	ErrResourceNotFound = NewAPIError(http.StatusNotFound, "RESOURCE_NOT_FOUND", "服务端找不到相应资源")

	ErrActionForbidden = NewAPIError(http.StatusForbidden, "ACTION_FORBIDDEN", "您被禁止执行该操作")

	ErrDeleteOrgForbidden = NewAPIError(http.StatusForbidden, "ACTION_FORBIDDEN", "部门有员工或岗位不能删除")

	ErrResourceDeleteFail = NewAPIError(http.StatusInternalServerError, "RESOURCE_DELETE_FAIL", "删除该资源不成功")

	ErrDisableDeleteUnique = NewAPIError(http.StatusInternalServerError, "RESOURCE_DELETE_FAIL", "唯一业务员/检测员不能删除")

	ErrResourceUpdateFail = NewAPIError(http.StatusInternalServerError, "RESOURCE_UPDATE_FAIL", "更新该资源不成功")

	ErrResourceUpdateConflict = NewAPIError(http.StatusConflict, "RESOURCE_UPDATE_CONFLICT", "资源更新时发生冲突")

	ErrMappingInvalid = NewAPIError(http.StatusBadRequest, "MAPPING_INVALID", "字段映射不合法")

	ErrNoRowFound = NewAPIError(http.StatusBadRequest, "EMPTY_SHEET", "Sheet为空没有发现数据")

	ErrNoRecordFound = NewAPIError(http.StatusBadRequest, "NO_RECORD_FOUND", "Sheet为空没有发现数据")

	ErrNoSheetFound = NewAPIError(http.StatusBadRequest, "NO_SHEET_FOUND", "没有找到sheet")

	ErrOneSheetSupported = NewAPIError(http.StatusBadRequest, "ONLY_ONE_SHEET_SUPPORT", "只支持一个sheet")

	ErrFileOpenFail = NewAPIError(http.StatusBadRequest, "FILE_READ_FAIL", "不能读取该文件内容")

	ErrExcelFileExtensionUnsupported = NewAPIError(http.StatusBadRequest, "EXCEL_FILE_EXT_UNSUPPORTED", "导入文件必须是xls或xlsx格式文件")

	ErrResumeFileExtensionUnsupported = NewAPIError(http.StatusBadRequest, "RESUME_FILE_EXT_UNSUPPORTED", "导入简历文件必须是doc或docx格式pdf文件")

	ErrResourceInactive = NewAPIError(http.StatusBadRequest, "RESOURCE_INACTIVE", "资源处于inactive状态")

	ErrRequiredColumnMissing = NewAPIError(http.StatusBadRequest, "REQUIRED_COLUMN_MISSING", "需要的列缺失")

	ErrTitleRowNotFound = NewAPIError(http.StatusBadRequest, "TITLE_ROW_NOT_FOUND", "没有找到包含员工姓名相关的信息表头")

	ErrGeneralTitleRowNotFound = NewAPIError(http.StatusBadRequest, "TITLE_ROW_NOT_FOUND", "没有找到可以识别的信息表头行")

	ErrEmailInvalid = NewAPIError(http.StatusBadRequest, "EMAIL_INVALID", "电子邮箱格式不正确")
	ErrIdNumInvalid = NewAPIError(http.StatusBadRequest, "IDNUM_INVALID", "身份证格式不正确")

	ErrFormFieldExist = NewAPIError(http.StatusBadRequest, "FORM_FIELD_EXIST", "表单字段已存在")

	ErrBatchInsertError = NewAPIError(http.StatusInternalServerError, "BATCH_INSERT_DATA_FAIL", "批量导入数据失败")

	ErrBatchInsertTargetOwnersError = NewAPIError(http.StatusInternalServerError, "BATCH_INSERT_DATA_FAIL", "批量插入数据失败")

	ErrBatchUpdateError = NewAPIError(http.StatusInternalServerError, "BATCH_UPDATE_DATA_FAIL", "批量更新数据失败")
	ErrBatchDeleteError = NewAPIError(http.StatusInternalServerError, "BATCH_DELETE_DATA_FAIL", "批量删除数据失败")

	ErrEmptyUpdateInput = NewAPIError(http.StatusBadRequest, "EMPTY_UPDATE_INPUT", "没有需要更新的内容")

	ErrUpdateSalaryCycleError = NewAPIError(http.StatusBadRequest, "SALARY_CYCLE_HAS_ISSUED", "该批次工资条已经发放")

	ErrGeneralInvalidParam = NewAPIError(http.StatusBadRequest, "General_Invalid_Param", "Parameters are invalid")

	ErrGeneralInvalidParamDabing = NewAPIError(http.StatusBadRequest, "General_Invalid_Param", "大病险种参数配置不正确")

	// file begin
	ErrOpenFileError   = NewAPIError(http.StatusBadRequest, "OPEN_FILE_ERROR", "Open file fail")
	ErrCreateFileError = NewAPIError(http.StatusBadRequest, "CREATE_FILE_ERROR", "Create file fail")
	ErrReadFileError   = NewAPIError(http.StatusBadRequest, "READ_FILE_ERROR", "Read file fail")
	ErrWriteFileError  = NewAPIError(http.StatusBadRequest, "WRITE_FILE_ERROR", "Write file fail")
	ErrDeleteFileError = NewAPIError(http.StatusBadRequest, "DELETE_FILE_ERROR", "Delete file fail")
	// file end

	// stream begin
	ErrCreateStreamError  = NewAPIError(http.StatusBadRequest, "CREATE_STREAM_ERROR", "Create steram fail")
	ErrSendStreamError    = NewAPIError(http.StatusBadRequest, "SEND_STREAM_ERROR", "Send strem fail")
	ErrReceiveStreamError = NewAPIError(http.StatusBadRequest, "RECEIVE_STREAM_ERROR", "Receive strem fail")
	// stream end

	ErrSheetTitleRowMissing = NewAPIError(http.StatusBadRequest, "SHEET_TITLE_ROW_MISSING", "表头列字段是必选项,请检查表头")

	ErrQueryParameterFault = &APIError{
		Code:    http.StatusBadRequest,
		Title:   "Query_Invalid_Param",
		Message: "Query parameter invalid",
	}

	//resource record begin
	ErrResourceCreateFail = NewAPIError(http.StatusInternalServerError, "RESOURCE_CREATE_FAIL", "创建资源失败")

	ErrCredentialNumberAndMobileDuplicate = NewAPIError(http.StatusConflict, "CREDENTIALNUMBER_MOBILE_DUPLICATE", "身份证号码/或手机号码重复，请重新导入工资条")
	ErrResourceDuplicate                  = NewAPIError(http.StatusConflict, "RESOURCE_DUPLICATE", "创建的资源和已有的资源冲突")
	ErrResourceDuplicateDetail            = NewAPIError(http.StatusConflict, "RESOURCE_DUPLICATE_DETAIL", "您已申请此项证明，请耐心等待审批")
	//resource record end

	ErrMobileNumberIncorrect = &APIError{
		Code:    http.StatusBadRequest,
		Title:   "Mobile_Number_Format_Incorrect",
		Message: "Mobile number format is incorrect",
	}

	ErrTelephoneNumberIncorrect = &APIError{
		Code:    http.StatusBadRequest,
		Title:   "Telephone_Number_Format_Incorrect",
		Message: "Telephone number format is incorrect",
	}

	ErrImTokenGetFail = &APIError{
		Code:    http.StatusInternalServerError,
		Title:   "Internal_Server_Error",
		Message: "Failed to get the im token from the im server",
	}

	ErrClaimMissing = &APIError{
		Code:    http.StatusUnauthorized,
		Title:   "UNAUTHORIZED",
		Message: "Required claims are missing in the token",
	}

	ErrTokenGenErr = &APIError{
		Code:    http.StatusInternalServerError,
		Title:   "Internal_Server_Error",
		Message: "Error happened when generating the access token",
	}

	ErrHTTPRequestFail = &APIError{
		Code:    http.StatusBadRequest,
		Title:   "HTTP_Request_Fail",
		Message: "HTTP request fail",
	}
	ErrDifferentTenant = &APIError{
		Code:    http.StatusForbidden,
		Title:   "Different_Tenant",
		Message: "Is not the same tenant",
	}

	// S3
	ErrS3GetBucketInfo = &APIError{
		Code:    http.StatusBadRequest,
		Title:   "S3_Get_Bucket_Info",
		Message: "Get s3 bucket info failed",
	}
	ErrS3GetSecurityTokenFaild = &APIError{
		Code:    http.StatusBadRequest,
		Title:   "S3_Get_Security_Token",
		Message: "Get s3 security token failed",
	}
	ErrS3CreateObject = &APIError{
		Code:    http.StatusBadRequest,
		Title:   "S3_Create_Object",
		Message: "Create s3 object failed",
	}
	ErrS3DeleteObject = &APIError{
		Code:    http.StatusBadRequest,
		Title:   "S3_Delete_Object",
		Message: "Delete s3 object failed",
	}

	ErrResourceTotalCountGetFail = &APIError{
		Code:    http.StatusInternalServerError,
		Title:   "Resource_Total_Count_Get_Fail",
		Message: "Failed to get the total count for the resouce",
	}
	ErrNoUpdateFound = NewAPIError(http.StatusBadRequest, "NO_UPDATE_FOUND", "没有更新需要操作")

	ErrNoFileDownload = NewAPIError(http.StatusBadRequest, "NO_FILE_DOWNLOAD", "没有文件可下载")

	ErrNoFileUpload = NewAPIError(http.StatusBadRequest, "NO_FILE_DOWNLOAD", "没有文件可上传")

	ErrImportJobInRun = NewAPIError(http.StatusForbidden, "IMPORT_JOB_IN_RUN", "导入job正在运行中")

	ErrImageFileInvalid = NewAPIError(http.StatusBadRequest, "IMAGE_FILE_INVALID", "图片文件无效")

	ErrTxHandlerRequire = NewAPIError(http.StatusInternalServerError, "TX_HANDLER_REQUIRE", "没有提供事物处理器")

	ErrTulingParseFault = NewAPIError(http.StatusInternalServerError, "ERR_TULING_PARSE_FAULT", "图灵机器人解析失败")

	ErrOperateLogNotSupport = NewAPIError(http.StatusForbidden, "ERR_OPERATE_LOG_UNSUPPORT", "操作日志不支持")

	ErrImportDisallowed = NewAPIError(http.StatusForbidden, "ERR_IMPORT_DISALLOWED", "您禁止导入该文件")

	ErrDataFormatUnsupport = NewAPIError(http.StatusBadRequest, "ERR_DATA_FORMAT_UNSUPPORT", "数据格式不支持")

	ErrRefreshNotExpires = NewAPIError(http.StatusBadRequest, "REFRESH_TIME_NOT_EXPIRES", "验证码还在有效期内")

	ErrSmsCodeNotFound = NewAPIError(http.StatusNotFound, "SMS_CODE_NOT_FOUND", "验证码不存在")

	ErrSmsCodeExpires = NewAPIError(http.StatusBadRequest, "SMS_CODE_EXPIRE", "验证码已过期")

	ErrSmsCodeIncorrect = NewAPIError(http.StatusBadRequest, "SMS_CODE_INCORRECT", "验证码输入错误")

	ErrPasswdIncorrect = NewAPIError(http.StatusBadRequest, "PASSWD_FORMAT_INCORRECT", "密码格式不正确: 要求同时含有数字和字母，且长度要在8-16位之间")

	ErrMobileNumberAlreadyRegistered = NewAPIError(http.StatusBadRequest, "MOBILE_NUMBER_ALREADY_REGISTERED", "手机号码已注册")
)

func (e *APIError) Response() gin.H {
	return gin.H{"code": e.Code, "title": e.Title, "message": e.Message, "errors": e.Errors, "x_request_id": e.XRequestId}
}

//Implement the built-in error interface
func (e *APIError) Error() string {

	b, _ := json.Marshal(e)
	return string(b)
}

func (e APIError) Copy4Lang(lang string) *APIError {
	T := i18n.T(lang)
	return NewFullAPIError(e.Code, e.Title, T(e.Message), copyNestedErrors(e.Errors), e.XRequestId)
}

func (e APIError) Copy4LangWithRequestId(lang string, requestId string) *APIError {
	T := i18n.T(lang)
	return NewFullAPIError(e.Code, e.Title, T(e.Message), copyNestedErrors(e.Errors), requestId)
}

func (e *APIError) AppendFieldError(field string, fieldErrs []string) {
	errors, ok := e.Errors[field]
	if !ok {
		errors = fieldErrs
	} else {
		errors = append(errors, fieldErrs...)
	}

	e.Errors[field] = errors
}

func (e *APIError) Equal(other *APIError) bool {
	return reflect.DeepEqual(e, other)
}

func copyNestedErrors(errors map[string][]string) map[string][]string {
	ret := make(map[string][]string)
	for k, v := range errors {
		errMsgs := make([]string, len(v))
		for index, errMsg := range v {
			errMsgs[index] = errMsg
		}
		ret[k] = errMsgs
	}

	return ret
}

func IsDuplicate(err error) bool {
	return strings.Index(strings.ToLower(err.Error()), "duplicate") != -1
}
