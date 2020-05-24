package common

const (
	// log level
	DEBUG = "debug"
	ERROR = "error"
	WARN  = "warn"
	INFO  = "info"

	// log format
	JSON = "json"
	TEXT = "text"

	// api server mode
	RELEASE_MODE = "release"
	DEBUG_MODE   = "debug"

	COMMON_YES  = "yes"
	COMMON_NO   = "no"
	COMMON_BACK = "back"

	// micro service type
	SERVICE_TYPE_INFRA    = "infra_service"
	SERVICE_TYPE_AUTH     = "auth_service"
	SERVICE_TYPE_BUSINESS = "business_service"

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
	SRV_NAME_DISTID          = "cn.com.kingtu.srv.distid"
	SRV_NAME_SMS             = "cn.com.kingtu.srv.sms"
	SRV_NAME_CUSTOMER        = "cn.com.kingtu.srv.customer"
	SRV_NAME_UAA             = "cn.com.kingtu.srv.uaa"
	SRV_NAME_FORMMETA        = "cn.com.kingtu.srv.formmeta"
	SRV_NAME_CRM             = "cn.com.kingtu.srv.crm"
	SRV_NAME_CONTRACT        = "cn.com.kingtu.srv.contract"
	SRV_NAME_SERVICECONTRACT = "cn.com.kingtu.srv.servicecontract"
	SRV_NAME_RECRUIT         = "cn.com.kingtu.srv.recruit"
	SRV_NAME_WXGATEWAY       = "cn.com.kingtu.srv.wxgateway"
	SRV_NAME_OSSGATEWAY      = "cn.com.kingtu.srv.ossgateway"
	SRV_NAME_EMPMGM          = "cn.com.kingtu.srv.empmgm"
	SRV_NAME_PROVIDERMGM     = "cn.com.kingtu.srv.providermgm"
	SRV_NAME_REWARDMGM       = "cn.com.kingtu.srv.rewardmgm"
	SRV_NAME_SOCIALINSURANCE = "cn.com.kingtu.srv.socialinsurance"
	SRV_NAME_SETTLEMGM       = "cn.com.kingtu.srv.settlemgm"
	SRV_NAME_TODOMGM         = "cn.com.kingtu.srv.todomgm"
	SRV_NAME_MESSAGEMGM      = "cn.com.kingtu.srv.messagemgm"
	SRV_NAME_RESUMEGATEWAY   = "cn.com.kingtu.srv.resumegateway"
	SRV_NAME_WOGATEWAY       = "cn.com.kingtu.srv.wogateway"
	SRV_NAME_FACEIDGATEWAY   = "cn.com.kingtu.srv.faceidgateway"
	SRV_NAME_BASEGATEWAY     = "cn.com.kingtu.srv.basegateway"
	SRV_NAME_SIPOLICY        = "cn.com.kingtu.srv.sipolicy"
	SRV_NAME_PSIORDER        = "cn.com.kingtu.srv.psiorder"
	SRV_NAME_DAILYWAGE       = "cn.com.kingtu.srv.dailywage"
	SRV_NAME_SOS             = "cn.com.kingtu.srv.sos"
	SRV_NAME_PERDIEM         = "cn.com.kingtu.srv.perdiem"
	SRV_NAME_PAYROLL         = "cn.com.kingtu.srv.payroll"
	SRV_NAME_HRPAYROLL       = "cn.com.kingtu.srv.hr-payroll"
	SRV_NAME_PAYROLL_HR      = "cn.com.kingtu.srv.hrpayroll"
	SRV_NAME_WALLETMGM       = "cn.com.kingtu.srv.walletmgm"
	SRV_NAME_COMMON          = "cn.com.kingtu.srv.common"
	SRV_NAME_ATTENDANCE      = "cn.com.kingtu.srv.attendance"
	SRV_NAME_STATISTICS      = "cn.com.kingtu.srv.statistics"
	SRV_NAME_OUTSOURCING     = "cn.com.kingtu.srv.outsourcing"
	SRV_NAME_SOCIALWORK      = "cn.com.kingtu.srv.socialwork"
	SRV_NAME_IDPHOTO         = "cn.com.kingtu.srv.idphoto"
	SRV_NAME_PHOTOGATEWAY    = "cn.com.kingtu.srv.photogateway"
	SRV_NAME_ANTGATEWAY      = "cn.com.kingtu.srv.antgateway"
	SRV_NAME_PERSONNELMGM    = "cn.com.kingtu.srv.personnelmgm"
	SRV_NAME_TAXGATEWAY      = "cn.com.kingtu.srv.taxgateway"
	SRV_NAME_BILLING         = "cn.com.kingtu.srv.billing"
	SRV_NAME_REMINDER        = "cn.com.kingtu.srv.reminder"
	SRV_NAME_OFFICIALACCOUNT = "cn.com.kingtu.srv.officialaccount"
	SRV_NAME_INTELLIGENCEMGM = "cn.com.kingtu.srv.intelligencemgm"
	SRV_NAME_CALLBOT         = "cn.com.kingtu.srv.callbot"
	SRV_NAME_PERSONTAX       = "cn.com.kingtu.srv.persontax"
	SRV_NAME_VIDEOTRANSCODE  = "cn.com.kingtu.srv.videotranscode"
	SRV_NAME_CHATBOT         = "cn.com.kingtu.srv.chatbot"
	// srv end

	// api begin
	API_NAME_UAA             = "cn.com.kingtu.api.uaa"
	API_NAME_SMS             = "cn.com.kingtu.api.sms"
	API_NAME_FORMMETA        = "cn.com.kingtu.api.formmeta"
	API_NAME_CRM             = "cn.com.kingtu.api.crm"
	API_NAME_CONTRACT        = "cn.com.kingtu.api.contract"
	API_NAME_SERVICECONTRACT = "cn.com.kingtu.api.servicecontract"
	API_NAME_RECRUIT         = "cn.com.kingtu.api.recruit"
	API_NAME_FILEMETA        = "cn.com.kingtu.api.filemeta"
	API_NAME_EMPMGM          = "cn.com.kingtu.api.empmgm"
	API_NAME_PROVIDERMGM     = "cn.com.kingtu.api.providermgm"
	API_NAME_REWARDMGM       = "cn.com.kingtu.api.rewardmgm"
	API_NAME_SOCIALINSURANCE = "cn.com.kingtu.api.socialinsurance"
	API_NAME_SETTLEMGM       = "cn.com.kingtu.api.settlemgm"
	API_NAME_TODOMGM         = "cn.com.kingtu.api.todomgm"
	API_NAME_MESSAGEMGM      = "cn.com.kingtu.api.messagemgm"
	API_NAME_COMMON          = "cn.com.kingtu.api.common"
	API_NAME_SIPOLICY        = "cn.com.kingtu.api.sipolicy"
	API_NAME_PSIORDER        = "cn.com.kingtu.api.psiorder"
	API_NAME_PERDIEM         = "cn.com.kingtu.api.perdiem"
	API_NAME_PAYROLL         = "cn.com.kingtu.api.payroll"
	API_NAME_PAYROLL_HR      = "cn.com.kingtu.api.hrpayroll"
	API_NAME_WALLETMGM       = "cn.com.kingtu.api.walletmgm"
	API_NAME_DAILYWAGE       = "cn.com.kingtu.api.dailywage"
	API_NAME_SOS             = "cn.com.kingtu.api.sos"
	API_NAME_ATTENDANCE      = "cn.com.kingtu.api.attendance"
	API_NAME_OUTSOURCING     = "cn.com.kingtu.api.outsourcing"
	API_NAME_BOT             = "cn.com.kingtu.api.bot"
	API_NAME_SOCIALWORK      = "cn.com.kingtu.api.socialwork"
	API_NAME_IDPHOTO         = "cn.com.kingtu.api.idphoto"
	API_NAME_PHOTOGATEWAY    = "cn.com.kingtu.api.photogateway"
	API_NAME_PERSONNELMGM    = "cn.com.kingtu.api.personnelmgm"
	API_NAME_BILLING         = "cn.com.kingtu.api.billing"
	API_NAME_REMINDER        = "cn.com.kingtu.api.reminder"
	API_NAME_OFFICIALACCOUNT = "cn.com.kingtu.api.officialaccount"
	API_NAME_STATISTICS      = "cn.com.kingtu.api.statistics"
	API_NAME_INTELLIGENCEMGM = "cn.com.kingtu.api.intelligencemgm"
	API_NAME_CALLBOT         = "cn.com.kingtu.api.callbot"
	API_NAME_PERSONTAX       = "cn.com.kingtu.api.persontax"
	API_NAME_CHATBOT         = "cn.com.kingtu.api.chatbot"
	API_NAME_TAXGATEWAY      = "cn.com.kingtu.api.taxbot"
	// api end

	// status begin
	ACTIVE_STATUS      = "active"
	INACTIVE_STATUS    = "inactive"
	OPENING_STATUS     = "opening"
	SEPARATING_STATUS  = "separating"
	TERMINATING_STATUS = "terminating"
	TERMINATED_STATUS  = "terminated"
	OVERDUE_STATUS     = "overdue"
	NEW_STATUS         = "new"
	INTERVIEW_STATUS   = "interview"
	CHECKUP_STATUS     = "checkup"
	ABANDON_STATUS     = "abandon"
	ONBOARD_STATUS     = "onboard"
	ONBOARDED_STATUS   = "onboarded"
	INSERVICE_STATUS   = "in-service" // 在职
	SEPARATE_STATUS    = "separate"   // 离职
	GENERATE_STATUS    = "generate"
	WAIT_STATUS        = "wait"
	CALCULATING_STATUS = "calculating"
	CALCULATED_STATUS  = "calculated"
	PAYING_STATUS      = "paying"
	CONFIRMING_STATUS  = "confirming"
	VERIFYING_STATUS   = "verifying"
	COMPLETED_STATUS   = "completed"
	INIT_STATUS        = "init"
	RUNNING_STATUS     = "running"
	DONE_STATUS        = "done"
	DOING_STATUS       = "doing"
	SUCCESS_STATUS     = "success"
	FAIL_STATUS        = "fail"
	CREATING_STATUS    = "creating"
	PROCESSING_STATUS  = "processing"
	DIALING_STATUS     = "dialing"
	CALLING_STATUS     = "calling"
	FAIL_DIAL_STATUS   = "fail_dial"
	PAUSED_STATUS      = "paused"
	HANGUP_STATUS      = "hangup"
	UNRESPOND_STATUS   = "unrespond"
	DIFF_STATUS        = "diff"
	// status end

	SMS_CODE_TYPE_REGISTER    = "register"
	SMS_CODE_TYPE_BIND_MOBILE = "bindmobile"
	SMS_CODE_TYPE_LOGIN       = "login"
	SMS_CODE_TYPE_RESETPWD    = "resetpwd"
	SMS_CODE_TYPE_PSIORDER    = "psiorder"
	SMS_CODE_TYPE_WALLETMGM   = "walletmgm"
	SMS_CODE_TYPE_BATCHSIGN   = "batchsign"

	TENANT_ADMIN_GLOBAL_ACCESS    = "tenant_admin"
	TENANT_CUSTOMER_GLOBAL_ACCESS = "tenant_customer"
	END_USER_GLOBAL_ACCESS        = "end_user"
	ENT_ADMIN                     = "ent_admin"

	PASSWORD_GRANT               = "password"
	REFRESH_TOKEN_GRANT          = "refresh_token"
	SMSCODE_GRANT                = "smscode"
	AUTHORIZATION_CODE_GRANT     = "authorization_code"
	WX_AUTHORIZATION_CODE_GRANT  = "wx_authorization_code"
	CLIENT_CREDENTIAL_CODE_GRANT = "client_credential"
	WX_APP_CODE_GRANT            = "wx_app_code"

	DEFAULT_OFFSET = 0
	DEFAULT_LIMIT  = 16
	CANCEL_OFFSET  = -1
	CANCEL_LIMIT   = -1
	PAGE_LIMIT     = 100
	START_OFFSET   = 0

	INVALID_INPUT_PARAMS = "Invalid_Input_Params"

	SCOPE_GLOBAL_ACCESS_KEY = "global_access"
	SCOPE_TENANT_KEY        = "tenant"
	SCOPE_CUSTOMER_KEY      = "customer"
	SCOPE_ENT_KEY           = "ent"
	SCOPE_DEPBRANCH_KEY     = "depbranch"

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

	// oss object type begin
	USER_IMAGE_TYPE                     = "user_image"
	CUSTOMER_LOGO_TYPE                  = "customer_logo"
	TENANT_LOGO_TYPE                    = "tenant_logo"
	TENANT_WXQRCODE_TYPE                = "tenant_wx_qrcode"
	POSITION_QRCODE_IMAGE_TYPE          = "pos_qrcode"
	CUSTOMER_POSITION_QRCODE_IMAGE_TYPE = "customer_pos_qrcode"
	OBJECT_TYPE_PSIORDER                = "psiorder"
	OBJECT_TYPE_PAYROLL                 = "payroll"
	OBJECT_TYPE_DAILYWAGE               = "dailywage"
	SERVICE_CONTRACT_OSS_FILE           = "servicecontract"
	SI_POLICY_OSS_FILE                  = "sipolicy"
	MYBANK_TENANT_CERT                  = "mybank_tenant_cert"
	BROKER_POSITION_QRCODE_IMAGE_TYPE   = "broker_pos_qrcode"
	OBJECT_TYPE_BROKER_CANDIDATE        = "broker_candidate"
	OBJECT_TYPE_OUTSOURCING             = "outsourcing"
	OBJECT_TYPE_STATISTICS              = "statistics"
	OBJECT_TYPE_VIDEO_COVER             = "video_cover"
	OBJECT_TYPE_ID_PHOTO                = "id_photo"
	OBJECT_TYPE_SOCIALWORK              = "socialwork"
	OBJECT_TYPE_WALLETMGM               = "walletmgm"
	OBJECT_TYPE_ORGANIXATION            = "organization"
	OBJECT_TYPE_SETTLEMGM               = "settlemgm"
	EMPMGM_NOTICE_QRCODE_IMAGE_TYPE     = "empmgm_notice_qrcode"
	OBJECT_TYPE_IMPORT_TEMPLATES        = "import_templates"
	OBJECT_TYPE_INSWXAPP_PROMOTION      = "inswxapp_promotion"
	OBJECT_TYPE_CALLBOT_RECORD          = "callbot_record"
	OBJECT_TYPE_RESUME_HEAD_IMAGE       = "resume_head_image"
	OBJECT_TYPE_TRAIN_VIDEO_COVER       = "train_video_cover"
	OBJECT_TYPE_UAA                     = "uaa"
	// oss object type end

	DEFAULT_TIME_ZONE = "Asia/Shanghai"

	// mqueue topic begin
	TOPIC_CANDIDATE_ONBOARD                        = "candidate_onboard"
	TOPIC_CANDIDATE_BATCH_ONBOARD                  = "candidate_batch_onboard"
	TOPIC_EMPLOYEE_ONBOARD                         = "employee_onboard"
	TOPIC_EMPLOYEE_NEW_BATCH_ONBOARD               = "employee_new_batch_onboard"
	TOPIC_ONBOARD_TEMPLATE_ITEM_MODIFY             = "onboard_template_item_modify"
	TOPIC_ONBOARD_MATERIAL_SUBMIT                  = "onboard_material_submit"
	TOPIC_EMPLOYEE_SEPARATE_CONTRACT               = "employee_separate_contract"
	TOPIC_EMPLOYEE_SEPARATE                        = "employee_separate"
	TOPIC_EMPLOYEE_SEPARATE_BY_PSI                 = "employee_separate_by_psi"
	TOPIC_PROVIDER_IMPORT                          = "provider_import"
	TOPIC_PROVIDER_TENANT_STATUS_UPDATE            = "provider_tenant_status_update"
	TOPIC_CANDIDATE_SYNC                           = "candidate_sync"
	TOPIC_CANDIDATE_UPDATE_SYNC                    = "candidate_update_sync"
	TOPIC_CANDIDATE_INTERVIEW_SYNC                 = "candidate_interview_sync"
	TOPIC_CANDIDATE_ABANDON                        = "candidate_abandon"
	TOPIC_RESUMES_COUNT_SYNC                       = "new_org_resume" //NEW ADD
	TOPIC_CUSTOMER_BRIEF_UPDATE                    = "customer_brief_update"
	TOPIC_CUSTOMER_BRIEF_BATCH_UPDATE              = "customer_brief_batch_update"
	TOPIC_NOTIFY_CANDIDATE_SYNC                    = "notify_candidate_sync"
	TOPIC_REMIND_NOTIFY_CANDIDATE_SYNC             = "remind_notify_candidate_sync"
	TOPIC_NOTIFY_INTERVIEWER_SYNC                  = "notify_interviewer_sync"
	TOPIC_EMPLOYEE_ONBOARD_SUPPORT                 = "employee_onboard_support"
	TOPIC_CANDIDATE_ONBOARD_SUPPORT                = "candidate_onboard_support"
	TOPIC_EMPLOYEE_IMPORT                          = "employee_import"
	TOPIC_FORM_IMPORT_SYNC                         = "form_import_sync"
	TOPIC_SALARY_IMPORT_SYNC                       = "salary_import_sync"
	TOPIC_SIMULATED_TAX_SALARY_IMPORT_SYNC         = "simulated_tax_import_sync"
	TOPIC_SIMULATED_TAX_TO_SALARY_IMPORT_SYNC      = "simulated_tax_to_salary_import_sync"
	TOPIC_HISTORY_SALARY_IMPORT_SYNC               = "salary_history_import_sync"
	TOPIC_SOCIAL_INSURANCE_IMPORT_SYNC             = "social_insurance_import_sync"
	TOPIC_WELFARE_IMPORT_SYNC                      = "welfare_import_sync"
	TOPIC_SALARY_HR_IMPORT_SYNC                    = "salay_hr_import_sync"
	TOPIC_SALARY_SLIP_IMPORT_SYNC                  = "salary_slip_import_sync"
	TOPIC_SALARY_SLIP_HR_IMPORT_SYNC               = "salary_slip_hr_import_sync"
	TOPIC_SALARY_RESULT_IMPORT_SYNC                = "salary_result_import_sync"
	TOPIC_PERSONNEL_TAX_RESULT_IMPORT_SYNC         = "personnel_tax_result_import_sync"
	TOPIC_WALLETMGM_PAYROLL_SYNC                   = "walletmgm_payroll_sync"
	TOPIC_WALLETMGM_PAYPOINTS_SYNC                 = "walletmgm_paypoints_sync"
	TOPIC_DECLARE_PERSONNEL_TAX_SYNC               = "declare_personnel_tax_sync"
	TOPIC_CONFIRM_SALARY_CYCLE_SYNC                = "confirm_salary_cycle_sync"
	TOPIC_DECLARE_HR_PERSONNEL_TAX_SYNC            = "declare_hr_personnel_tax_sync"
	TOPIC_CALCULATE_PERSONNEL_TAX_SYNC             = "calculate_personnel_tax_sync"
	TOPIC_CALCULATE_HR_PERSONNEL_TAX_SYNC          = "calculate_hr_personnel_tax_sync"
	TOPIC_EMPLOYEE_SALARY_SETTING                  = "employee_salary_setting"
	TOPIC_PAYROLL_SALARY_SETTING                   = "payroll_salary_setting"
	TOPIC_NOTIFY_POSITION_SYNC                     = "notify_position_sync"
	TOPIC_GENERATE_PRE_BILL                        = "generate_pre_bill"
	TOPIC_DATUM_WARNING                            = "datum_warning"
	TOPIC_ENT_INSURED_PERSON_IMPORT                = "ent_insured_person_import"
	TOPIC_ENT_INSURED_PERSON_IMPORT_TEMP           = "ent_insured_person_import_temp"
	TOPIC_ENT_INSURED_PERSON_IMPORT_FILL           = "ent_insured_person_import_fill"
	TOPIC_ENT_INSURED_PERSON_CREATE_RENEW          = "ent_insured_person_create_renew"
	TOPIC_EMPLOYEE_UPDATE_TO_PSIORDER_SYNC         = "employee_update_to_psiorder_sync"
	TOPIC_EMPLOYEE_SI_HF_UPDATE_TO_PSIORDER_SYNC   = "employee_update_si_hf_to_psiorder_sync"
	TOPIC_CANDIDATE_ONBOARDED_SYNC                 = "candidate_onboarded_sync"
	TOPIC_EMPLOYEE_ABANDON                         = "employee_abandon"
	TOPIC_INSURANCE_FORM_IMPORT_SYNC               = "insurance_form_import_sync"
	TOPIC_FORM_FILL                                = "form_fill"
	TOPIC_FEEDBACK_SYNC                            = "feedback_sync"
	TOPIC_FORM_SUBMIT_SYNC                         = "form_submit_sync"
	TOPIC_BACKPROCESS_IMPORT_SYNC                  = "backprocess_import_sync"
	TOPIC_CUSTOMER_OTHER_FEE_IMPORT_SYNC           = "customer_other_fee_import_sync"
	TOPIC_INSURANCE_REAL_IMPORT_SYNC               = "insurance_real_import_sync"
	TOPIC_SIPOLICY_HANDLE_SYNC                     = "sipolicy_handle_sync"
	TOPIC_OUTSOURCING_CHECKINFO                    = "outsourcing_check_info"
	TOPIC_CANDIDATE_IMPORT_JOB                     = "candidate_import_job"
	TOPIC_CANDIDATE_INSERVICE                      = "candidate_inservice"
	TOPIC_NOTIFY_BROKER_AUDITED_SYNC               = "notify_broker_audited_sync"
	TOPIC_BROKER_TRANSFORM                         = "broker_transform_sync"
	TOPIC_NEW_BROKER_WXMESSAGE_SYNC                = "new_broker_wxmessage"
	TOPIC_NOTIFY_BROKER_BATCH_AUDIT_SYNC           = "notify_broker_batch_audit_sync"
	TOPIC_IDPHOTO_ORDER_SYNC                       = "idphoto_order_sync"
	TOPIC_BROKER_ACHIEVEMENT_SYNC                  = "broker_achievement_sync"
	TOPIC_NOTIFY_BOT_MSG                           = "notify_bot_msg"
	TOPIC_CUSTOMER_IMPORT_JOB                      = "customer_import_job"
	TOPIC_CUSTOMER_IMPORT_SYNC                     = "customer_import_sync"
	TOPIC_REAL_NAME_AUTH                           = "real_name_auth"
	TOPIC_USER_INFO_UPDATE                         = "user_info_update"
	TOPIC_RECRUIT_CANDIDATE_BATCH_ONBOARD          = "recruit_candidate_batch_onboard"
	TOPIC_RECRUIT_RESUME_ALLOCATE                  = "recruit_resume_allocate"
	TOPIC_RECRUIT_RESUME_TRANSFER_OUT              = "recruit_resume_transfer_out"
	TOPIC_BOT_NOTIFY                               = "bot_notify"
	TOPIC_CRM_CUSTOMER_ENTTENANT_SYNC              = "crm_sutomer_enttenant_sync"
	TOPIC_SOCIALWORK_TASK_POINT_SYNC               = "socialwork_task_point_sync"
	TOPIC_SOCIALWORK_TASK_REMARK_WORKSTATUS_SYNC   = "socialwork_task_remark_workstatus_sync"
	TOPIC_SOCIALWORK_POINTS_UPDATE_PROCESSING_SYNC = "socialwork_points_update_processing_sync"
	TOPIC_RESUME_IMPORT_JOB                        = "resume_import_job"
	TOPIC_RESUME_IMPORT_BRK                        = "broker_import_brk"
	TOPIC_RESUME_EXTRACT_JOB                       = "resume_extract_job"
	TOPIC_EMPLOYEE_ONBOARD_LAUNCH_CONTRACT         = "employee_onboard_launch_contract"
	TOPIC_POINT_ISSUEMENT_SYNC                     = "point_issuement"
	TOPIC_EMPLOYEE_UPDATE_CANDIDATE_SYNC           = "employee_update_candidate_sync"
	TOPIC_CONTRACT_IMPORT_JOB                      = "contract_import_job"
	TOPIC_SETTLEMGM_ACCOUNT_BOOK_BIND_SYNC         = "settlemgm_account_book_bind"
	TOPIC_SETTLEMGM_BILL_EMPLOYEE_BIND_SYNC        = "settlemgm_bill_employee_bind"
	TOPIC_NOTIFY_BROKERAGE_SETTLEMENT              = "notify_brokerage_settlement"
	TOPIC_BROKER_ATTENDANCE_SYNC                   = "broker_attendance_sync"
	TOPIC_UAA_TENNTADMIN_HANDOVER_SYNC             = "tenantadmin_handover_sync"
	TOPIC_PLAN_EMPLOYEE_BANDING_SYNC               = "plan_employee_banding_sync"
	TOPIC_PLAN_EMPLOYEE_SEPARATE_SYNC              = "plan_employee_separate_sync"
	TOPIC_CONTRACT_TIMEUPDATE_EMPLOYEE_SYNC        = "contract_timeupdate_employee_sync"
	TOPIC_EMPLOYEE_IMPORT_SYNC                     = "employee_import_sync"
	TOPIC_EMPLOYEE_NOTICE_SYNC                     = "employee_notice_sync"
	TOPIC_HISTORICAL_PAY_RECORD_SYNC               = "historical_pay_record_import_sync"    //历史缴纳记录
	TOPIC_PSIORDER_BACKPROCESS_EXT_IMPORT_SYNC     = "psiorder_backprocess_ext_import_sync" //社保后道外部数据导入
	TOPIC_EMPLOYEE_TEMPLATE_SYNC                   = "employee_template_sync"
	TOPIC_EMPLOYEE_SIGN_STATUS_SYNC                = "employee_sign_status_sync"
	TOPIC_EMPLOYEE_USER_SYNC                       = "employee_user_sync"
	TOPIC_EMPLOYEE_RESUME_SYNC                     = "employee_resume_sync"
	TOPIC_EMPMGM_JOB_SETTLEMENT_BIND_SYNC          = "topic_empmgm_job_settlement_bind_sync" // 雇员入离职与结算的绑定
	TOPIC_NOTICE_EMPMGM_UPDATE_SETTLE              = "topic_notice_empmgm_update_settle"     // 通知雇员修改结算设置
	TOPIC_EMPMGM_NOTICE_BILL                       = "topic_empmgm_notice_bill"              // 导入考勤记录通知empmgm，是否需要更新账单
	TOPIC_CALLBOT_BATCH_TASK_GENERATE              = "callbot_batch_task_generate"
	TOPIC_CALLBOT_CALL_TASK_LAUNCH                 = "call_bot_call_task_launch"
	TOPIC_CALLBOT_VOICE_REPORT                     = "callbot_voice_report"
	TOPIC_CALLBOT_VOICE_RECORD_REPORT              = "callbot_voice_record_report"
	TOPIC_CALLBOT_FIELD_RESYNC                     = "callbot_field_resync"
	TOPIC_PERSON_INFORMATION_IMPORT_JOB            = "person_information_import_job"
	TOPIC_EMPLOYEE_INFORMATION_IMPORT_JOB          = "employee_information_import_job"
	TOPIC_AUTO_RECOMMEND_MISMATCH                  = "auto_recommend_mismatch" //自动推荐删除标记
	TOPIC_INTELLIGENCEMGM_TRAINING                 = "start_training"
	TOPIC_INTELLIGENCEMGM_VIDEO_TRANS              = "train_video_trans"
	TOPIC_EMPLOYEE_IMPORT_PERSON_INFORMATION       = "employee_import_person_information"
	TOPIC_EMPLOYEE_IMPORT_HR_SYNC                  = "employee_import_hr_sync"
	TOPIC_SEPARATE_EMPLOYEE_INFORMATION_IMPORT     = "separate_employee_information_import" //离职任职受雇信息
	TOPIC_PERSON_ABNORMAL_IMPORT                   = "person_abnormal_import"               //非正常自然人
	TOPIC_TENANTADMIN_UPDATE_INFO                  = "tenantadmin_update_info"              //变更信息
	TOPIC_SOCIALWORK_UPDATE_FREEDOM                = "topic_socialwork_update_freedom"
	TOPIC_CONTRACT_SIGN_CONTRACT                   = "contract_sign_contract"

	TOPIC_PERSON_TO_SALARY_CYCLE_SYNC = "person_to_salary_cycle_sync" //自然人申报完成反馈   自然人----> 工资

	TOPIC_SALARY_RECORD_DEDUCTION_SYNC     = "salary_record_deduction_status"
	TOPIC_GET_SALARY_RECORD_DEDUCTION_SYNC = "salary_record_deduction_sync"
	TOPIC_ATTENDANCE_PAYROLL_SYNC          = "attendance_payroll_sync"
	TOPIC_OPERATOR_LOGS_SYNC               = "operator_logs_sync"

	TOPIC_PERSON_DECLARE_STATUE_UPDATE = "person_declare_status_update"
	TOPIC_NOTICE_SPECIAL_DECLARE       = "notice_special_declare"

	TOPIC_LEGALENTITY_NOTIFY_SETTLEMGM_SYNC = "legalentity_notify_settlemgm_sync"
	TOPIC_REFERENCE_SERVICECONTRACT_SYNC    = "reference_servicecontract"
	TOPIC_UAA_LOGIN_SYNC                    = "uaa_login_sync"
	TOPIC_CHAECK_RECORD_DEDUCYION_SYNC      = "check_record_deduction"
	TOPIC_PERSON_ADDITION_REMIND            = "person_addition_remind"      //提醒完善专项附加扣除信息
	TOPIC_CALL_NOTICE                       = "call_notice"                 //语音通知
	TOPIC_EMPLOYEE_ACTIVE_UPDATE_SYNC       = "employee_active_update_sync" //雇员激活推送
	TOPIC_ADDITION_REDECLARE_SYNC           = "addition_redeclare_sync"     //专项附加扣除重新申报
	TOPIC_EMPLOYEE_TERMINATE_CONTRACT_SYNC  = "employee_terminate_contract"
	TOPIC_ADDITION_FAILED_SYNC              = "topic_addition_failed_sync" //专项附加扣除失败提醒
	TOPIC_SOCIALWORK_UPDATE_MOBILE          = "socialwork_update_mobile"   // 用工池修改手机号
	TOPIC_PERSON_TO_EMPLOYEE_UPDATE_SYNC    = "person_to_employee_update"
	TOPIC_EMPLOYEE_TO_PERSON_SYNC           = "employee_to_person_sync" //雇员信息同步到自然人
	TOPIC_PERSON_DECLARE_SYNC               = "person_declare_sync"     //自然人申报

	TOPIC_PAYROLL_PERSON_IMPORT    = "person_import_sync"
	TOPIC_BATCH_OPERATE_CLERK_SYNC = "batch_operate_clerk_sync"
	TOPIC_TAX_SCHEDULING           = "tax_scheduling"

	TOPIC_TAX_PERSON_EMPLOYEE_SYNC    = "tax_person_employee_sync"    // 自然人信息同步到雇员
	TOPIC_EMPLOYEE_SALARY_IMPORT_SYNC = "employee_salary_import_sync" //银行卡信息导入

	// mqueue topic end

	// mqueue channel begin
	//TOPIC_NOTIFY_POSITION_SYNC                   = "notify_position_sync"
	CHANNEL_SRV_RESUMES                          = "com.workai.recruit.srv.statistics" //new ADD
	CHANNEL_SRV_EMPMGM                           = "cn.com.kingtu.srv.empmgm"
	CHANNEL_SRV_RECRUIT                          = "cn.com.kingtu.srv.recruit"
	CHANNEL_SRV_PROVIDERMGM                      = "cn.com.kingtu.srv.providermgm"
	CHANNEL_SRV_REWARDMGM                        = "cn.com.kingtu.srv.rewardmgm"
	CHANNEL_SRV_RECRUIT_EMPLOYEE_SEPARATE_RESUME = "cn.com.kingtu.srv.recruit_employee_separate_resume"
	CHANNEL_SRV_PSIORDER                         = "cn.com.kingtu.srv.psiorder"
	CHANNEL_SRV_SIPOLICY                         = "cn.com.kingtu.srv.sipolicy"
	CHANNEL_SRV_SALARY                           = "cn.com.kingtu.srv.salary"
	CHANNEL_SRV_SALARY_HR                        = "cn.com.kingtu.srv.salary_hr"
	CHANNEL_SRV_PERSONNEL_TAX                    = "cn.com.kingtu.srv.personnel_tax"
	CHANNEL_SRV_PERSONNEL_TAX_HR                 = "cn.com.kingtu.srv.personnel_tax_hr"
	CHANNEL_SRV_WALLETMGM                        = "cn.com.kingtu.srv.walletmgm"
	CHANNEL_SRV_RECRUIT_BROKER                   = "cn.com.kingtu.srv.recruit_broker"
	CHANNEL_SRV_OUTSOURCING_CHECKINFO            = "cn.com.kingtu.srv.outsourcing_checkinfo"
	CHANNEL_SRV_IDPHOTO                          = "cn.com.kingtu.srv.idphoto"
	CHANNEL_SRV_CRM                              = "cn.com.kingtu.srv.crm"
	CHANNEL_SRV_UAA                              = "cn.com.kingtu.srv.uaa"
	CHANNEL_SRV_UAA_CRM                          = "cn.com.kingtu.srv.uaa_crm"
	CHANNEL_SRV_UAA_PAYROLL                      = "cn.com.kingtu.srv.uaa_payroll"
	CHANNEL_SRV_UAA_PROVIDER                     = "cn.com.kingtu.srv.uaa_provider"
	CHANNEL_SRV_UAA_PSIORDER                     = "cn.com.kingtu.srv.uaa_psiorder"
	CHANNEL_SRV_UAA_RECRUIT                      = "cn.com.kingtu.srv.uaa_recruit"
	CHANNEL_SRV_UAA_SERVICECONTRACT              = "cn.com.kingtu.srv.uaa_servicecontract"
	CHANNEL_SRV_UAA_SETTLEMGM                    = "cn.com.kingtu.srv.uaa_settlemgm"
	CHANNEL_SRV_UAA_SIPOLICY                     = "cn.com.kingtu.srv.uaa_sipolicy"
	CHANNEL_SRV_CONTRACT                         = "cn.com.kingtu.srv.contract"
	CHANNEL_SRV_COMMON                           = "cn.com.kingtu.srv.common"
	CHANNEL_SRV_SERVICECONTRACT                  = "cn.com.kingtu.srv.servicecontract"
	CHANNEL_SRV_SETTLEMGM                        = "cn.com.kingtu.srv.settlemgm"
	CHANNEL_SRV_ATTENDANCE                       = "cn.com.kingtu.srv.attendance"
	CHANNEL_SRV_FORMMETA                         = "cn.com.kingtu.srv.formmeta"
	CHANNEL_SRV_CALLBOT                          = "cn.com.kingtu.srv.callbot"
	CHANNEL_SRV_OFFICIALACCOUNT                  = "cn.com.kingtu.srv.officialaccount"
	CHANNEL_SRV_SOCIALWORK                       = "cn.com.kingtu.srv.socialwork"
	CHANNEL_SRV_PERSONTAX                        = "cn.com.kingtu.srv.persontax"
	CHANNEL_SRV_INTELLIGENCEMGM                  = "cn.com.kingtu.srv.intelligencemgm"
	CHANNEL_SRV_PERSONNELMGM                     = "cn.com.kingtu.srv.personnelmgm"
	CHANNEL_SRV_OUTSOURCING_STOREINFO            = "cn.com.kingtu.srv.outsourcing_storeinfo"

	// mqueue channel begin

	//file suffix
	XLS_EXT  = ".xls"
	XLSX_EXT = ".xlsx"
	CSV_EXT  = ".csv"
	ZIP_EXT  = ".zip"
	//Excel row category
	DATA_ROW_CATEGORY          = "data"
	BEFORE_HEADER_ROW_CATEGORY = "before-header"
	HEADER_ROW_CATEGORY        = "header"
	SUM_ROW_CATEGORY           = "sum"
	BLANK_ROW_CATEGORY         = "blank"
	EMPLOYEE_ROW_CATEGORY      = "employee"

	// settle frequency
	SETTLE_FREQUENCY_MONTH     = "0"
	SETTLE_FREQUENCY_HALF_YEAR = "1"
	SETTLE_FREQUENCY_YEAR      = "2"

	// attribute type
	ATTRIBUTE_TYPE_CUSTOMER = "customer"
	ATTRIBUTE_TYPE_SELF     = "self"

	OBJECT_TYPE_CONTRACT = "contract"

	// sign type
	CONTRACT_SIGN_NOTSIGNED = "not-signed"
	CONTRACT_SIGN_SIGNING   = "signing"
	CONTRACT_SIGN_SIGNED    = "signed"

	// http request content type
	CONTENT_TYPE_JSON = "application/json"
	CONTENT_TYPE_XML  = "text/xml"

	// ext app cred begin
	EXT_APP_CRED_CATEGORY_SOCIAL_INSURANCE = "social_insurance"
	EXT_APP_CRED_TYPE_WX_APP               = "0"
	// ext app cred end

	// owner org type
	ORG_TYPE_TENANT   = "0"
	ORG_TYPE_CUSTOMER = "1"

	//rongcloud im api format
	IM_API_JSON_FORMAT = "json"

	// excel cell font color
	FONT_COLOR_RED = "FFFF0000"

	// service contract content type
	SERVICE_CONTENT_TYPE_RECRUIT          = "recruit"
	SERVICE_CONTENT_TYPE_PAYROLL          = "payroll"
	SERVICE_CONTENT_TYPE_SOCIAL_INSURANCE = "social_insurance"
	SERVICE_CONTENT_TYPE_SOCIAL_WORK      = "social_work"

	LEGAL_ENTITY          = "legal_entity"
	LEGAL_ENTITY_CUSTOMER = "legal_entity"
	LEGAL_ENTITY_TENANT   = "tenant"

	//HRO_DOMAIN = "oss-cn-beijing.aliyuncs.com"
	// payment method
	PAYMENT_METHOD_WECHATPAY = "WeChatPay"
	PAYMENT_METHOD_ALIPAY    = "AliPay"

	// Content-Type(Mime-Type)
	MIME_TYPE_VIDEO_MPEG4 = "video/mpeg4"
	MIME_TYPE_VIDEO_MP4   = "video/mp4"

	TENANT_TYPE_HRO = "0"
	TENANT_TYPE_ENT = "1"
	TENANT_TYPE_HR  = "2"

	TENANT_TYPE_HRR = "1"

	TENANT_RELATION_TYPE_CUSTOMER = "customer"

	ENT_TENANT_TYPE_SOCIALWORK = "socialwork"

	PERM_VIEW_ALL = "0"

	EMPLOYEE_IMPORT_TYPE                     = "employee"
	RESUME_IMPORT_TYPE                       = "resume"
	CANDIDATE_IMPORT_TYPE                    = "candidate"
	CRM_IMPORT_TYPE                          = "customer"
	CONTRACT_IMPORT_TYPE                     = "contract"
	RECRUIT_BROKER_CANDIDATE_IMPORT_TYPE     = "broker_candidate"
	PSIORDER_IMPORT_TYPE                     = "psiorder"
	PSIORDER_ENTI_INSURED_PEOPLE_IMPORT_TYPE = "psiorder-ent_insured_people"

	YES = "y"
	NO  = "n"

	CUSTOMER_CREATE_REGISTER_TYPE = "register"
)

//context key
const (
	CtxUserNameKey                 = "User-Name"
	CtxUserIdKey                   = "User-Id"
	CtxTenantIdKey                 = "Tenant-Id"
	CtxScopeKey                    = "Scope"
	SIPOLICY_CACL_TYPE_NORMAL      = "normal"
	SIPOLICY_CACL_TYPE_UPDATE_BASE = "update_base"
)

const (
	PENSION_INSURANCE      = "养老保险"
	MEDICAL_INSURANCE      = "医疗保险"
	LNJURY_INSURANCE       = "工伤保险"
	MATERNITY_INSURANCE    = "生育保险"
	UNEMPLOYMENT_INSURANCE = "失业保险"
	PROVIDENT_FUND         = "公积金"
)

// 外呼机器人回填策略
const (
	CALLBOT_NO_RESYNC       = "0"
	CALLBOT_EXIST_NO_RESYNC = "1"
	CALLBOT_RESYNC          = "2"
)

const (
	BELONG_CATEGORY_SELF     = "0" // 自营
	BELONG_CATEGORY_CUSTOMER = "1" // 客户
	BELONG_CATEGORY_PROVIDER = "2" // 供应商
)

const (
	// 工资批次状态
	RELEASE_STATUS_DATA_PROCESSING = "data_processing"             //数据处理中
	RELEASE_STATUS_WAITING_PERSON  = "waiting_person_declaration"  //待申报自然人信息
	RELEASE_STATUS_WAITING_SPECIAL = "waiting_special_declaration" //待申报专项附加扣除
	RELEASE_STATUS_CALCULATING     = "calculating"                 //计算中
	RELEASE_STATUS_ACTIVE          = "active"                      //待确认
	RELEASE_STATUS_PENDING         = "pending"                     //待发放
	RELEASE_STATUS_PAYROLL         = "payroll"                     //申报中／发放中（网商）
	RELEASE_STATUS_COMPLETED       = "completed"                   //已发放
	RELEASE_STATUS_CONFIRMED       = "confirmed"                   //已确认（首次校准）
	RELEASE_STATUS_ABNORMAL        = "abnormal"                    //数据异常
	RELEASE_STATUS_FAILED          = "failed"                      //发放失败
	RELEASE_STATUS_DELETED         = "deleted"                     //已删除

	// 工资详情状态
	RECORD_STATUS_UNISSUED = "UNISSUED" //未发放 for slip
	RECORD_STATUS_INIT     = "init"     //待发放
	RECORD_STATUS_PAYROLL  = "payroll"  //发放中（网商）
	RECORD_STATUS_SUCCESS  = "success"  //发放成功
	RECORD_STATUS_FAILED   = "failed"   //发放失败

	SALARY_BLOCK_STATUS_WAITING_CALCULATE = "waiting"     //待计算
	SALARY_BLOCK_STATUS_CALCULATING       = "calculating" //计算中
	SALARY_BLOCK_STATUS_SYNCING           = "syncing"     //专项附加同步中
	SALARY_BLOCK_STATUS_SYNCDONE          = "syncdone"    //专项附加同步完成
	SALARY_BLOCK_STATUS_SUNCFAILED        = "syncfailed"  //专项附加反馈失败
	SALARY_BLOCKSTATUS_NONE               = "syncnone"    //专项附加同步完成没有数据

	MESSAGE_STATUS_ABNORMAL = "abnormal" //专项反馈异常处理

	SALARY_TAX_STATUS_WAITING_CALCULATE = "waiting_calculate" //待计算---待申报专项附加扣
	SALARY_TAX_STATUS_CALCULATING       = "calculating"       //计算中---计算中
	SALARY_TAX_STATUS_SUCCESS           = "success"           //计算完成---待确认
	SALARY_TAX_STATUS_FAILED            = "failed"            //计算失败---数据异常

	//个税批次申报状态
	TAX_STATUS_WAITING_DECLARE = "waiting_declare" //待申报
	TAX_STATUS_DECLARING       = "declaring"       //申报中
	TAX_STATUS_SUCCESS         = "success"         //申报成功
	TAX_STATUS_FAILED          = "failed"          //申报失败
	TAX_STATUS_PAYMENTING      = "paymenting"      //缴款中
	TAX_STATUS_PAYMENT_SUCCESS = "payment_success" //缴款成功
	TAX_STATUS_PAYMENT_FAILED  = "payment_failed"  //缴款失败
	TAX_STATUS_CANCELED        = "canceled"        //申报作废
	//算税状态码
	CALCULATE_TAX_STATUS_FAILED   = -1 //失败
	CALCULATE_TAX_STATUS_SUCCESS  = 0  //成功
	CALCULATE_TAX_STATUS_REMINDER = 1  //提示

	TAX_CATEGORY_SALARY      = "salary"
	TAX_CATEGORY_SERVICE_FEE = "service_fee"

	PROGRAM_BELONG_TO_CUSTOMER = "customer"
	PROGRAM_BELONG_TO_TENANT   = "tenant"

	//社保来源
	INSURANCE_SOURCE_INTERNAL = "internal"
	INSURANCE_SOURCE_IMPORT   = "import"

	//excel单元格类型
	CELL_TYPE_TEXT   = "text"
	CELL_TYPE_NUMBER = "number"
	CELL_TYPE_DATE   = "date"

	DECLARE_UNIT_TYPE_CUSTOMER = "customer"
	DECLARE_UNIT_TYPE_TENANT   = "tenant"
	DECLARE_UNIT_TYPE_PROVIDER = "provider"

	//个税申报类型
	TAX_CATEGORY_SALARY_INCOME         = "0101" //工资薪金
	TAX_CATEGORY_FOREIGN_SALARY_INCOME = "0102" //外籍工资薪金
	TAX_CATEGORY_FOREIGN_MONTHLY_BONUS = "0104" //外籍人数月奖金
	TAX_CATEGORY_LABOR_INCOME          = "0401" //劳务报酬
	TAX_CATEGORY_BONUS_INCOME          = "0103" //全年一次性奖金
	TAX_CATEGORY_TERMINATE_CONTRACT    = "0108" //解除劳动合同

	//语音通知类型
	CALL_NOTICE_TYPE_TAX_PERSON_INFO_ADD     = "tax_person_info_add"           //自然人信息补充, TTSParam例子: []string{"姓名", "公司名称"}
	CALL_NOTICE_TYPE_TAX_SPECIALL_DEDUCT_ADD = "tax_person_special_deduct_add" //专项附加扣除补充, TTSParam例子: []string{"姓名", "公司名称"}
	CALL_NOTICE_TYPE_TAX_PERSON_COMPOSE_ADD  = "tax_person_compose_add"        //自然人或者专项附加扣除补充, TTSParam例子: []string{"姓名", "公司名称"}

	//自然人完善状态
	PERSON_NOT_PERFECT = "0" //自然人基础信息未完善
	PERSON_HAS_PERFECT = "1" //自然人基础信息已完善

	//自然人报送状态
	PERSON_UNDECLARED       = "0" //待报送
	PERSON_IN_DECLARATION   = "1" //报送中
	PERSON_DECLARE_SUCCESS  = "2" //报送成功
	PERSON_DECLARE_FAILURE  = "3" //报送失败
	PERSON_DECLARE_QUERY    = "4" //报送状态查询
	PERSON_DECLARE_FEEDBACK = "5" //报送反馈
	PERSON_NO_DECLARE_UNIT  = "6" //无报送单位

	//纳税人类型
	TAXPAYER_TYPE_RESIDENT     = "resident"
	TAXPAYER_TYPE_NON_RESIDENT = "non-resident"

	//算税任务类型
	TAX_SCHEDULING_CALCULATING_TAX            = "calculating_tax"            //算税任务
	TAX_SCHEDULING_DECLARING_PERSON           = "declaring_person"           //自然人申报
	TAX_SCHEDULING_DECLARING_SPECIAL_ADDITION = "declaring_special_addition" //专项附加扣除申报

	// attribute
	SYSTEM_ATTRIBUTE = "system"

	//角色
	HRO_ROLE_FINANCE = "finance"

	PERSON_IMPORT_TYPE = "person_import"
)

// 应用类型
const (
	APP_CATEGORY_SW        = "social_work"
	APP_CATEGORY_SI        = "social_insurance"
	APP_CATEGORY_TAX_ROBOT = "tax_robot"

	// default app category
	APP_CATEGORY_DEFAULT = "all"

	TAX_ROBOT_TYPE_REMIND_PERSON_DECLARATION  = "remind_person_declaration"
	TAX_ROBOT_TYPE_REMIND_SAD_CONFIRM         = "remind_sad_confirm"
	TAX_ROBOT_TYPE_INVITATE_ACTIVATE          = "invitate_activate"
	TAX_ROBOT_TYPE_REMIND_IMPROVE_INFORMATION = "remind_improve_information"
	TAX_ROBOT_TYPE_PERSON_DECLARATION_SUCCESS = "person_declaration_success"
	TAX_ROBOT_TYPE_PERSON_DECLARATION_FAILED  = "person_declaration_failed"
	TAX_ROBOT_TYPE_ADD_PERSONAL_EMPLOYEE      = "add_personal_employee"
)

var (
	//缴费方式,1固定值, 0基数*比例
	InsuranceRuleCodePayTypeMap = map[string]string{
		"0":  "0",
		"1":  "1",
		"2":  "1",
		"3":  "0",
		"4":  "0",
		"5":  "0",
		"6":  "1",
		"7":  "1",
		"8":  "1",
		"9":  "1",
		"10": "1",
		"11": "0",
	}
)

var (
	//税友接口错误码 错误提示文案
	TaxErrorCodes = map[string]string{
		"0001": "系统申报错误，请重新申报",
		"0002": "系统申报错误，请重新申报",
		"0003": "系统申报错误，请重新申报",
		"0004": "系统申报错误，请重新申报",
		"0006": "系统申报错误，请重新申报",
		"0007": "系统申报错误，请重新申报",
		"0008": "系统申报错误，请重新申报",
		"0009": "任务处理失败",
		"0010": "任务执行中断",
		"0100": "系统申报错误，请重新申报",
		"0101": "任务执行中断",
		"0102": "任务执行中断",
		"0103": "任务执行中断",
		"0104": "企业纳税人识别号不正确，请修改",
		"0105": "任务执行中断",
		"0111": "系统申报错误，请重新申报",
		"0110": "系统申报错误，请重新申报",
		"0112": "系统申报错误，请重新申报",
		"0119": "系统申报错误，请重新申报",
		"0120": "系统申报错误，请重新申报",
		"0121": "系统申报错误，请重新申报",
		"0122": "存在未填写手机号码的人员",
		"0129": "显示msg信息",
		"0130": "反馈异常",
		"0131": "没有报送成功的人员",
		"0132": "人员存在待报送或者报送失败的情况",
		"0133": "人员存在待反馈的情况",
		"0200": "税款计算错误",
		"0201": "任务执行中断",
		"0202": "任务执行中断",
		"0203": "任务执行中断",
		"0204": "任务执行中断",
		"0205": "任务执行中断",
		"0206": "任务执行中断",
		"0207": "任务执行中断",
		"0211": "税款计算失败",
		"0212": "税款计算失败",
		"0213": "税款计算失败",
		"0214": "税款计算失败",
		"0215": "税款计算失败",
		"0216": "税款计算失败",
		"0221": "税款计算失败",
		"0222": "税款计算失败",
		"0223": "税款计算失败",
		"0224": "税款计算失败",
		"0225": "税款计算失败",
		"0226": "税款计算失败",
		"0251": "税款计算失败",
		"0252": "税款计算失败",
		"0253": "税款计算失败",
		"0254": "税款计算失败",
		"0255": "税款计算失败",
		"0256": "税款计算失败",
		"0257": "税款计算失败",
		"0260": "税款计算失败",
		"0300": "系统申报错误，请重新申报",
		"0400": "系统申报错误，请重新申报",
		"0500": "系统申报错误，请重新申报",
		"0600": "系统申报错误，请重新申报",
		"0700": "任务执行中断",
		"0800": "系统申报错误，请重新申报",
		"0801": "税号对应的企业不存在，请修改",
		"0802": "企业纳税人识别号对应多家实际企业，需要指定企业登记序号",
		"0900": "系统申报错误，请重新申报",
		"9999": "系统申报错误，请重新申报",
		//special area error code
		"999999": "系统申报错误，请重新申报",
		"000001": "任务执行中断",
		"000002": "任务执行中断",
		"000004": "请求处理中",
		"000003": "重复请求",
		"010000": "该企业存在多个登记序号",
		"010001": "实名账号密码校验失败",
		"010002": "申报密码校验失败",
		"010003": "登记序号不匹配",
		"010004": "发送短信验证码失败",
		"010005": "发送短信验证码过于频繁",
		"030000": "任务执行中断",
		"040001": "非法的申报表类型",
		"040003": "申报状态不符合作废要求",
		"100003": "税款计算失败",
		"100004": "税款计算失败",
		"100007": "税款计算失败",
		"100008": "税款计算失败",
		"100009": "税款计算失败",
		"100010": "税款计算失败",
		"100011": "税款计算失败",
		"100012": "税款计算失败",
		"100013": "税款计算失败",
		"100014": "税款计算失败",
		"100015": "税款计算失败",
		"100016": "税款计算失败",
		"200001": "三方信息请求失败",
		"200002": "当前申报状态无法进行三方信息请求",
		"200011": "当前申报状态无法进行欠费查询请求",
		"200021": "扣款请求失败",
		"200022": "当前申报状态无法进行缴款查询请求",
		"200023": "存在多个三方协议号",
		"200024": "无欠费记录",
	}
)

// nlu
const (
	NLU_TAXBOT = "tax_bot"
)

var (
	//税友国籍字典
	TaxNations = []string{"阿富汗", "阿尔巴尼亚", "南极洲", "阿尔及利亚", "美属萨摩亚", "安道尔", "安哥拉", "安提瓜和巴布达", "阿塞拜疆", "阿根廷", "澳大利亚", "奥地利", "巴哈马", "巴林", "孟加拉国", "亚美尼亚", "巴巴多斯", "比利时", "百慕大", "不丹", "玻利维亚", "波黑", "博茨瓦纳", "布维岛", "巴西", "伯利兹", "英属印度洋领地", "所罗门群岛", "英属维尔京群岛", "文莱", "保加利亚", "缅甸", "布隆迪", "白俄罗斯", "柬埔寨", "喀麦隆", "加拿大", "佛得角", "开曼群岛", "中非", "斯里兰卡", "乍得", "智利", "中国", "中国台湾", "圣诞岛", "科科斯（基林）群岛", "哥伦比亚", "科摩罗", "马约特", "刚果（布）", "刚果（金）", "库克群岛", "哥斯达黎加", "克罗地亚", "古巴", "塞浦路斯", "捷克", "贝宁", "丹麦", "多米尼克", "多米尼加", "厄瓜多尔", "萨尔瓦多", "赤道几内亚", "埃塞俄比亚", "厄立特里亚", "爱沙尼亚", "法罗群岛", "福克兰群岛（马尔维纳斯）", "南乔治亚岛和南桑德韦奇岛", "斐济", "芬兰", "法属圭亚那", "法属波利尼西亚", "法属南部领地", "吉布提", "加蓬", "格鲁吉亚", "冈比亚", "巴勒斯坦", "德国", "加纳", "直布罗陀", "基里巴斯", "希腊", "格陵兰", "格林纳达", "瓜德罗普", "关岛", "危地马", "几内亚", "圭亚那", "海地", "赫德岛和麦克唐纳岛", "梵蒂冈", "洪都拉斯", "中国香港", "匈牙利", "冰岛", "印度", "印度尼西亚", "伊朗", "伊拉克", "爱尔兰", "以色列", "意大利", "科特迪瓦", "牙买加", "日本", "哈萨克斯坦", "约旦", "肯尼亚", "朝鲜", "韩国", "科威特", "吉尔吉斯斯坦", "老挝", "黎巴嫩", "莱索托", "拉脱维亚", "利比里亚", "利比亚", "列支敦士登", "立陶宛", "卢森堡", "中国澳门", "马达加斯加", "马拉维", "马来西亚", "马尔代夫", "马里", "马耳他", "马提尼克", "毛里塔尼亚", "毛里求斯", "墨西哥", "摩纳哥", "蒙古", "摩尔多瓦", "黑山", "蒙特塞拉特", "摩洛哥", "莫桑比克", "阿曼", "纳米比亚", "瑙鲁", "尼泊尔", "荷兰", "荷属安的列斯", "阿鲁巴", "新喀里多尼亚", "瓦努阿图", "新西兰", "尼加拉瓜", "尼日尔", "尼日利亚", "纽埃", "诺福克岛", "挪威", "北马里亚纳", "美国本土外小岛屿", "密克罗尼西亚联邦", "马绍尔群岛", "帕劳", "巴基斯坦", "巴拿马", "巴布亚新几内亚", "巴拉圭", "秘鲁", "菲律宾", "皮特凯恩", "波兰", "葡萄牙", "几内亚比绍", "东帝汶", "波多黎各", "卡塔尔", "留尼汪", "罗马尼亚", "俄罗斯联邦", "卢旺达", "圣赫勒拿", "圣基茨和尼维斯", "安圭拉", "圣卢西亚", "圣皮埃尔和密克隆", "圣文森特和格林纳丁斯", "圣马力诺", "圣多美和普林西比", "沙特阿拉伯", "塞内加尔", "塞尔维亚", "塞舌尔", "塞拉利昂", "新加坡", "斯洛伐克", "越南", "斯洛文尼亚", "索马里", "南非", "津巴布韦", "西班牙", "西撒哈拉", "苏丹", "苏里南", "斯瓦尔巴岛和扬马延岛", "斯威士兰", "瑞典", "瑞士", "叙利亚", "塔吉克斯坦", "泰国", "多哥", "托克劳", "汤加", "特立尼达和多巴哥", "阿联酋", "突尼斯", "土耳其", "土库曼斯坦", "特克斯和凯科斯群岛", "图瓦卢", "乌干达", "乌克兰", "前南马其顿", "埃及", "英国", "坦桑尼亚", "美国", "美属维尔京群岛", "布基纳法索", "乌拉圭", "乌兹别克斯坦", "委内瑞拉", "瓦利斯和富图纳", "萨摩亚", "也门", "南斯拉夫", "赞比亚", "法国"}
)

var (
	//权限
	AdminTypePermKeyMap = map[string]string{
		"员工关系专员": "employee",
		"薪酬专员":   "payroll",
		"个税申报专员": "persontax",
		"财务专员":   "finance",
	}
)

var (
	//权限
	PermKeyAdminTypeMap = map[string]string{
		"employee":  "员工关系专员",
		"payroll":   "薪酬专员",
		"persontax": "个税申报专员",
		"finance":   "财务专员",
	}
)

const (
	LEGAL_ENTITY_TYPE_TENANT   = "tenant"
	LEGAL_ENTITY_TYPE_CUSTOMER = "customer"
)

var (
	IDCARD_TYPES = []string{"居民身份证", "中国护照", "港澳居民来往内地通行证", "港澳居民居住证", "台湾居民来往大陆通行证", "台湾居民居住证", "外国护照", "外国人永久居留身份证", "外国人工作许可证（A类）", "外国人工作许可证（B类）", "外国人工作许可证（C类）", "身份证"}
)

const (
	DATA_RANGE_ADDITION = "addition" //专项附加扣除数据可见范围
	DATA_RANGE_PAYROLL  = "payroll"  //薪酬明细数据可见范围
	DATA_RANGE_PREPAY   = "prepay"   //预扣预缴明细数据可见范围

	ADMIN_TYPE_ADMIN     = "系统管理员"
	ADMIN_TYPE_EMPLOYEE  = "员工关系专员"
	ADMIN_TYPE_PAYROLL   = "薪酬专员"
	ADMIN_TYPE_PERSONTAX = "个税申报专员"
	ADMIN_TYPE_FINANCE   = "财务专员"
	ADMIN_TYPE_ALL       = "全部"
)

//银行简称
const (
	BANK_ABB_CBC   = "cbc"   //中国建设银行
	BANK_ABB_CMBC  = "cmbc"  //招商银行
	BANK_ABB_CITIC = "citic" //中信实业银行
)

const (
	PROJECT_NAME = "MallWeb"
	JAEGER_ADDR  = "localhost:6831"
)
