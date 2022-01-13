/*
 * DataBag
 *
 * DataBag provides storage for decentralized identity based self-hosting apps. It is intended to support sharing of personal data and hosting group conversations. 
 *
 * API version: 0.0.1
 * Contact: roland.osborne@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package databag

import (
  "log"
  "encoding/json"
	"net/http"
  "gorm.io/gorm"
  "golang.org/x/crypto/bcrypt"
  "databag/internal/store"
)

func AddNodeAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetNodeAccountImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetNodeAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetNodeClaimable(w http.ResponseWriter, r *http.Request) {

  if _configured {
    w.WriteHeader(http.StatusNotAcceptable)
  } else {
    w.WriteHeader(http.StatusOK)
  }
}

func GetNodeConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ImportAccount(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
}

func RemoveNodeAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func SetNodeAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func SetNodeClaim(w http.ResponseWriter, r *http.Request) {

  // confirm node is claimable
  if _configured {
    w.WriteHeader(http.StatusUnauthorized)
    return
  }

  // extract credentials
  username, password, ok := r.BasicAuth();
  if !ok {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    log.Printf("SetNodeClaim - failed to hash password");
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  // store credentials
  err = store.DB.Transaction(func(tx *gorm.DB) error {
    if res := tx.Create(&store.Config{ConfigId: CONFIG_USERNAME, StrValue: username}).Error; res != nil {
      return res
    }
    if res := tx.Create(&store.Config{ConfigId: CONFIG_PASSWORD, BinValue: hashedPassword}).Error; res != nil {
      return res
    }
    return nil;
  })
  if(err != nil) {
    log.Printf("SetNodeCalim - failed to store credentials");
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  // set global values
  _adminUsername = username
  _adminPassword = hashedPassword

	w.WriteHeader(http.StatusOK)
}

func SetNodeConfig(w http.ResponseWriter, r *http.Request) {

  // validate admin password
  username, password, ok := r.BasicAuth();
  if !ok {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  if username != _adminUsername || bcrypt.CompareHashAndPassword(_adminPassword, []byte(password)) != nil {
    log.Printf("SetNodeConfig - invalid admin credentials");
    w.WriteHeader(http.StatusUnauthorized);
    return
  }

  // parse node config
  r.Body = http.MaxBytesReader(w, r.Body, CONFIG_BODYLIMIT)
  dec := json.NewDecoder(r.Body)
  dec.DisallowUnknownFields()
  var config NodeConfig;
  res := dec.Decode(&config);
  if res != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  // store credentials
  err := store.DB.Transaction(func(tx *gorm.DB) error {
    if res := tx.Create(&store.Config{ConfigId: CONFIG_DOMAIN, StrValue: config.Domain}).Error; res != nil {
      return res
    }
    if res := tx.Create(&store.Config{ConfigId: CONFIG_PUBLICLIMIT, NumValue: config.PublicLimit}).Error; res != nil {
      return res
    }
    if res := tx.Create(&store.Config{ConfigId: CONFIG_STORAGE, NumValue: config.AccountStorage}).Error; res != nil {
      return res
    }
    return nil;
  })
  if(err != nil) {
    log.Printf("SetNodeConfig - failed to store config");
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  // set global values
  _nodeDomain = config.Domain
  _publicLimit = config.PublicLimit
  _accountStorage = config.AccountStorage

	w.WriteHeader(http.StatusOK)
}
