package session

import (
	"sync"
	"encoding/base64"
	"fmt"
	"io"
	"crypto/rand"
	"net/http"
	"net/url"
	"time"
)

type SessionManager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

type Provider interface {
	SessionInit(sid string)      (Session, error)
	SessionRead(sid string)      (Session, error)
	SessionDestroy(sid string)    error
	SessionGc(maxLifeTime int64)
}

type Session interface {
	Get(key interface{})         interface{}
	Set(key, val interface{})    error
	Delete(key interface{})      error
	SessionId()                  string
}

var globalSessManager *SessionManager
var provides = make(map[string]Provider)

func init() {
	globalSessManager, _ = NewManager("gosessionId", "moemory", 3600)
}

func Register(name string, provider Provider){
	if provider == nil {
		panic("error : session provider is nil")
	}
	if _, dup := provides[name]; dup{
		panic("duplicate error : Register called twice for provider" + name)
	}
	provides[name] = provider
}

func NewManager(cookieName, providerName string, maxLifeTime int64) (*SessionManager, error) {
	provider, ok := provides[providerName]
	if !ok {
		return nil, fmt.Errorf("unknown session provider: %s",providerName)
	}
	return &SessionManager{cookieName:cookieName, provider:provider, maxLifeTime: maxLifeTime}, nil
}

func (manager *SessionManager)SessionId() string{//之后可以改成uid生成
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil{
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *SessionManager) SessionStart (w http.ResponseWriter, r *http.Request) Session{ //开启session
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.SessionId()
		session, _ := manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name:manager.cookieName, Value:url.QueryEscape(sid), Path:"/", HttpOnly:true, MaxAge:int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
		return session
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ := manager.provider.SessionRead(sid)
		return session
	}
}

func (manager *SessionManager)SessionDestory(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		manager.provider.SessionDestroy(sid)
		expiration := time.Now()
		cookie := http.Cookie{Name:manager.cookieName, Path:"/", HttpOnly:true, Expires:expiration, MaxAge:-1}
		http.SetCookie(w, &cookie)
	}
}

func (manager *SessionManager) Gc() { //定时回收
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGc(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() {
		manager.Gc()
	})
}