package common

const (
	TIME_MONTH  = "month"
	TIME_DAY    = "day"
	TIME_HOUR   = "hour"
	TIME_MINUTE = "minute"
	TIME_SECOND = "second"
	TIME_YEAR   = "year"

	SMSCODE_LEN = 6
	// 验证码缓存时间, 单位：秒
	SMSCODE_EXPIRATION = 60 * 5

	COUNTRY_CHINA = "中国"
	COUNTRY_USA   = "美国"

	BASE64_TABLE = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr23456017krhr"

	// srv begin
	SRV_NAME_APIGATEWAY = "API Gateway"
	SRV_NAME_DISTID     = "cn.com.kingtu.srv.distid"
	SRV_NAME_SMS        = "cn.com.kingtu.srv.sms"
	SRV_NAME_CUSTOMER   = "cn.com.kingtu.srv.customer"
	// srv end

	SMSCODE_GRANT = "smscode"

	TIME_FORMAT                 = "2006-01-02T15:04:05-07:00"
	DATE_FORMAT                 = "2006-01-02"
	DATE_FORMAT1                = "2006/01/02"
	DATE_FORMAT2                = "20060102"
	DATE_FORMAT3                = "2006.01.02"
	DATE_FORMAT4                = "2006.01.2"
	DATE_FORMAT5                = "2006.1.02"
	DATE_FORMAT6                = "2006.1.2"
	DATE_FORMAT7                = "2006/1/2"
	DATE_FORMAT8                = "2006/01/2"
	DATE_FORMAT9                = "2006/1/02"
	MONTH_FORMAT                = "2006-01"
	MONTH_FORMAT1               = "2006/01"
	MONTH_FORMAT2               = "200601"
	MONTH_FORMAT3               = "2006.01"
	MONTH_FORMAT4               = "2006.1"
	MONTH_FORMAT5               = "2006/1"
	MONTH_FORMAT6               = "2006-1"
	DATE_TIME_FORMAT            = "2006-01-02 15:04:05"
	DATE_TIME_FORMAT1           = "2006/01/02 15:04:05"
	DATE_HM_FORMAT              = "2006/01/02 15:04"
	CHINA_MONTH_FORMAT          = "2006年1月"
	CHINA_DATE_FORMAT           = "2006年1月2日"
	CHINA_DATE_FORMAT_WITH_ZERO = "2006年01月02日"
	CHINA_DATETIME_FORMAT       = "2006年1月2日 15:04:05"
	CHINA_DATE_HM_FORMAT        = "2006年1月2日 15:04"
	CHINA_HM_FORMAT             = "15:04"
	MONTH_FORMAT_V1             = "200601"
	CHINA_MONTH_FORMAT1         = "1月"
	CHINA_MONTH_FORMAT2         = "2006年01"
	DATETIME_FORMAT2            = "20060102 15:04:05"

	DEFAULT_TIME_ZONE = "Asia/Shanghai"
)
const (
	PROJECT_NAME = "MallWeb"
)
