package go_user_manage


import (
	"net/http"
	"github.com/hadihabashi/sredis"

)

type IUserState interface {
	UserRights(req *http.Request) bool
	HasUser(username string) bool
	BooleanField(username, fieldname string) bool
	SetBooleanField(username, fieldname string, val bool)
	IsConfirmed(username string) bool
	IsLoggedIn(username string) bool
	AdminRights(req *http.Request) bool
	IsAdmin(username string) bool
	UsernameCookie(req *http.Request) (string, error)
	SetUsernameCookie(w http.ResponseWriter, username string) error
	AllUsernames() ([]string, error)
	Email(username string) (string, error)
	PasswordHash(username string) (string, error)
	AllUnconfirmedUsernames() ([]string, error)
	ConfirmationCode(username string) (string, error)
	AddUnconfirmed(username, confirmationCode string)
	RemoveUnconfirmed(username string)
	MarkConfirmed(username string)
	RemoveUser(username string)
	SetAdminStatus(username string)
	RemoveAdminStatus(username string)
	AddUser(username, password, email string)
	SetLoggedIn(username string)
	SetLoggedOut(username string)
	Login(w http.ResponseWriter, username string) error
	ClearCookie(w http.ResponseWriter)
	Logout(username string)
	Username(req *http.Request) string
	CookieTimeout(username string) int64
	SetCookieTimeout(cookieTime int64)
	CookieSecret() string
	SetCookieSecret(cookieSecret string)
	PasswordAlgo() string
	SetPasswordAlgo(algorithm string) error
	HashPassword(username, password string) string
	SetPassword(username, password string)
	CorrectPassword(username, password string) bool
	AlreadyHasConfirmationCode(confirmationCode string) bool
	FindUserByConfirmationCode(confirmationcode string) (string, error)
	Confirm(username string)
	ConfirmUserByConfirmationCode(confirmationcode string) error
	SetMinimumConfirmationCodeLength(length int)
	GenerateUniqueConfirmationCode() (string, error)

	Users() sredis.IHashMap
	Host() IHost
	Creator() sredis.ICreator
}

type IHost interface {
	Ping() error
	Close()
}

type IPermissions interface {
	SetDenyFunction(f http.HandlerFunc)
	DenyFunction() http.HandlerFunc
	UserState() IUserState
	Clear()
	AddAdminPath(prefix string)
	AddUserPath(prefix string)
	AddPublicPath(prefix string)
	SetAdminPath(pathPrefixes []string)
	SetUserPath(pathPrefixes []string)
	SetPublicPath(pathPrefixes []string)
	Rejected(w http.ResponseWriter, req *http.Request) bool
	ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc)
}