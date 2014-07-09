package auth

import (
  "code.google.com/p/go.crypto/scrypt"
  "encoding/json"
  //"github.com/astaxie/beego/session"
  "net/http"
  "regexp"
  "time"
)


type User struct {
  Username string
  pass string
  Is_anonymous bool
  Is_autenthicated bool
  Last_login string
  Perms []string
}

func (self *User) Anonymous() {
  self.Username = "Anonymous"
  self.pass = ""
  self.Is_anonymous = true
  self.Is_autenthicated = false
  self.Last_login = ""
}

func encryptpass(pass string) string {
  salt := "$"
  key, _ := scrypt.Key([]byte(pass), []byte(salt), 16384, 8, 1, 32)

  return string(key[:])
}

func (self *User) check(username, pass string) bool {
  //TODO: query para obtener el usuario 
  /*
  if self.pass == encryptpass(pass) {
    return true
  }
  */
  return false
}

func (self *User) checkUsername(username string) bool {
  if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{1,32}$", username); !ok {
    return false
  }
  return true
}

func (self *User) authenticate(username, pass string) bool {
  if !self.check(username, pass) {
    return false
  }
  return true
}

func (self *User) Login(w http.ResponseWriter, r *http.Request) bool {
  now := time.Now().Format("2006-01-02 15:04:05")

  self.Last_login = now
  self.Is_autenthicated = true
  self.Is_anonymous = false

  //sess := globalSessions.SessionStart(userCtrl.Ctx.ResponseWriter, userCtrl.Ctx.Request)
  sess := globalSessions.SessionStart(w, r)
  encoded, _ := json.Marshal(self)
  sess.Set("user", encoded)

  return true
}
