package auth

import (
  //"encoding/json"
  //"fmt"
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/session"
  "time"
)

const (
  cookie = "gosessionid"
)

var globalSessions *session.Manager

type BaseController struct {
	beego.Controller
  Sess session.SessionStore
}

func (self *BaseController) Session() {
  globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
  go globalSessions.GC()

  if self.Ctx.GetCookie(cookie) == "" {
    return
  }

  w := self.Ctx.ResponseWriter
  r := self.Ctx.Request
  sess := globalSessions.SessionStart(w, r)

  createtime := sess.Get("createtime")
  if createtime == nil {
    sess.Set("createtime", time.Now().Unix())
  } else if (createtime.(int64) + 60) < (time.Now().Unix()) {
    globalSessions.SessionDestroy(w, r)
    sess = globalSessions.SessionStart(w, r)
  }

  self.Sess = sess
}
