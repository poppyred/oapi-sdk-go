//  code generated by oapi sdk gen
package client

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/httpclient"
	"github.com/larksuite/oapi-sdk-go/service/acs/v1"
	"github.com/larksuite/oapi-sdk-go/service/admin/v1"
	"github.com/larksuite/oapi-sdk-go/service/application/v6"
	"github.com/larksuite/oapi-sdk-go/service/approval/v4"
	"github.com/larksuite/oapi-sdk-go/service/attendance/v1"
	"github.com/larksuite/oapi-sdk-go/service/baike/v1"
	"github.com/larksuite/oapi-sdk-go/service/bitable/v1"
	"github.com/larksuite/oapi-sdk-go/service/calendar/v4"
	"github.com/larksuite/oapi-sdk-go/service/contact/v3"
	"github.com/larksuite/oapi-sdk-go/service/docx/v1"
	"github.com/larksuite/oapi-sdk-go/service/drive/v1"
	"github.com/larksuite/oapi-sdk-go/service/ehr/v1"
	"github.com/larksuite/oapi-sdk-go/service/event/v1"
	"github.com/larksuite/oapi-sdk-go/service/gray_test_open_sg/v1"
	"github.com/larksuite/oapi-sdk-go/service/human_authentication/v1"
	"github.com/larksuite/oapi-sdk-go/service/im/v1"
	"github.com/larksuite/oapi-sdk-go/service/mail/v1"
	"github.com/larksuite/oapi-sdk-go/service/optical_char_recognition/v1"
	"github.com/larksuite/oapi-sdk-go/service/passport/v1"
	"github.com/larksuite/oapi-sdk-go/service/search/v2"
	"github.com/larksuite/oapi-sdk-go/service/sheets/v3"
	"github.com/larksuite/oapi-sdk-go/service/speech_to_text/v1"
	"github.com/larksuite/oapi-sdk-go/service/task/v1"
	"github.com/larksuite/oapi-sdk-go/service/tenant/v2"
	"github.com/larksuite/oapi-sdk-go/service/translation/v1"
	"github.com/larksuite/oapi-sdk-go/service/vc/v1"
	"github.com/larksuite/oapi-sdk-go/service/wiki/v2"
)

type Client struct {
	Acs                    *acs.AcsService
	Admin                  *admin.AdminService
	Application            *application.ApplicationService
	Approval               *approval.ApprovalService
	Attendance             *attendance.AttendanceService
	Baike                  *baike.BaikeService
	Bitable                *bitable.BitableService
	Calendar               *calendar.CalendarService
	Contact                *contact.ContactService
	Docx                   *docx.DocxService
	Drive                  *drive.DriveService
	Ehr                    *ehr.EhrService
	Event                  *event.EventService
	GrayTestOpenSg         *gray_test_open_sg.GrayTestOpenSgService
	HumanAuthentication    *human_authentication.HumanAuthenticationService
	Im                     *im.ImService
	Mail                   *mail.MailService
	OpticalCharRecognition *optical_char_recognition.OpticalCharRecognitionService
	Passport               *passport.PassportService
	Search                 *search.SearchService
	Sheets                 *sheets.SheetsService
	SpeechToText           *speech_to_text.SpeechToTextService
	Task                   *task.TaskService
	Tenant                 *tenant.TenantService
	Translation            *translation.TranslationService
	Vc                     *vc.VcService
	Wiki                   *wiki.WikiService
}

type ClientOptionFunc func(config *core.Config)

func WithAppType(appType core.AppType) ClientOptionFunc {
	return func(config *core.Config) {
		config.AppType = appType
	}
}

func WithDisableTokenCache() ClientOptionFunc {
	return func(config *core.Config) {
		config.EnableTokenCache = false
	}
}

func WithLogger(logger core.Logger) ClientOptionFunc {
	return func(config *core.Config) {
		config.Logger = logger
	}
}

func WithDomain(domain string) ClientOptionFunc {
	return func(config *core.Config) {
		config.Domain = domain
	}
}

func WithLogLevel(logLevel core.LogLevel) ClientOptionFunc {
	return func(config *core.Config) {
		config.LogLevel = logLevel
	}
}

func WithTokenCache(cache core.Cache) ClientOptionFunc {
	return func(config *core.Config) {
		config.TokenCache = cache
	}
}

func WithLogReqRespInfoAtDebugLevel(printReqRespLog bool) ClientOptionFunc {
	return func(config *core.Config) {
		config.LogReqRespInfoAtDebugLevel = printReqRespLog
	}
}

func WithHelpdeskCredential(helpdeskID, helpdeskToken string) ClientOptionFunc {
	return func(config *core.Config) {
		config.HelpDeskId = helpdeskID
		config.HelpDeskToken = helpdeskToken
		if helpdeskID != "" && helpdeskToken != "" {
			config.HelpdeskAuthToken = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", helpdeskID, helpdeskToken)))
		}
	}
}

func WithReqTimeout(reqTimeout time.Duration) ClientOptionFunc {
	return func(config *core.Config) {
		config.ReqTimeout = reqTimeout
	}
}

func NewClient(appId, appSecret string, options ...ClientOptionFunc) *Client {
	// 构建配置
	config := &core.Config{
		Domain:           FeishuDomain,
		AppId:            appId,
		AppSecret:        appSecret,
		EnableTokenCache: true,
		AppType:          core.AppTypeCustom,
	}
	for _, option := range options {
		option(config)
	}

	// 构建日志器
	core.NewLogger(config)

	// 构建缓存
	core.NewCache(config)

	// 创建httpclient
	httpClient := httpclient.NewHttpClient(config)

	// 创建sdk-client，并初始化服务
	client := &Client{}
	initService(client, httpClient, config)
	return client
}

func initService(client *Client, httpClient *http.Client, config *core.Config) {
	client.Acs = acs.NewService(httpClient, config)
	client.Admin = admin.NewService(httpClient, config)
	client.Application = application.NewService(httpClient, config)
	client.Approval = approval.NewService(httpClient, config)
	client.Attendance = attendance.NewService(httpClient, config)
	client.Baike = baike.NewService(httpClient, config)
	client.Bitable = bitable.NewService(httpClient, config)
	client.Calendar = calendar.NewService(httpClient, config)
	client.Contact = contact.NewService(httpClient, config)
	client.Docx = docx.NewService(httpClient, config)
	client.Drive = drive.NewService(httpClient, config)
	client.Ehr = ehr.NewService(httpClient, config)
	client.Event = event.NewService(httpClient, config)
	client.GrayTestOpenSg = gray_test_open_sg.NewService(httpClient, config)
	client.HumanAuthentication = human_authentication.NewService(httpClient, config)
	client.Im = im.NewService(httpClient, config)
	client.Mail = mail.NewService(httpClient, config)
	client.OpticalCharRecognition = optical_char_recognition.NewService(httpClient, config)
	client.Passport = passport.NewService(httpClient, config)
	client.Search = search.NewService(httpClient, config)
	client.Sheets = sheets.NewService(httpClient, config)
	client.SpeechToText = speech_to_text.NewService(httpClient, config)
	client.Task = task.NewService(httpClient, config)
	client.Tenant = tenant.NewService(httpClient, config)
	client.Translation = translation.NewService(httpClient, config)
	client.Vc = vc.NewService(httpClient, config)
	client.Wiki = wiki.NewService(httpClient, config)

}

var FeishuDomain = "https://open.feishu.cn"
var LarkDomain = "https://open.larksuite.com"
