package memory

import (
	"container/list"
	"ncfwxen/common/session"
	"ncfwxen/common/utils"
	"sync"
	"time"
)

type SessionStore struct {
	sid          string
	timeAccessed time.Time                   //最近访问时间
	value        map[interface{}]interface{} //session存储的值
}

type Provider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element //存储的session元素
	list     *list.List               //list指针用于gc
}

var pder = &Provider{list: list.New()}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", pder)
}

func (sess *SessionStore) Set(key, val interface{}) error {
	sess.value[key] = val
	pder.SessionUpdate(sess.sid)
	return nil
}

func (sess *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(sess.sid)
	if v, ok := sess.value[key]; ok {
		return v
	}
	return nil
}

func (sess *SessionStore) Delete(key interface{}) error {
	delete(sess.value, key)
	pder.SessionUpdate(sess.sid)
	return nil
}

func (sess *SessionStore) SessionId() string {
	return sess.sid
}

func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newSession := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newSession)
	pder.sessions[sid] = element
	utils.Nlog.Println("session started")
	return newSession, nil
}

func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		session, err := pder.SessionInit(sid)
		return session, err
	}
}

func (pder *Provider) SessionDestroy(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		pder.list.Remove(element)
		delete(pder.sessions, sid)
		return nil
	}
	return nil
}

func (pder *Provider) SessionGc(maxLifeTime int64) {
	pder.lock.Lock()
	pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		} else {
			if element.Value.(*SessionStore).timeAccessed.Unix()+maxLifeTime < time.Now().Unix() {
				pder.list.Remove(element)
				delete(pder.sessions, element.Value.(*SessionStore).sid)
			} else {
				break
			}
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	utils.Nlog.Println("session update")
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}
