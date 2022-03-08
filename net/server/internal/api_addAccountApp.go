package databag

import (
  "net/http"
  "time"
  "encoding/hex"
  "databag/internal/store"
  "github.com/theckman/go-securerandom"
)

func AddAccountApp(w http.ResponseWriter, r *http.Request) {

  account, err := AccountLogin(r)
  if err != nil {
    ErrResponse(w, http.StatusUnauthorized, err)
    return
  }

  data, res := securerandom.Bytes(4)
  if res != nil {
    ErrResponse(w, http.StatusInternalServerError, res)
    return
  }
  token := hex.EncodeToString(data)

  accountToken := store.AccountToken{
    AccountID: account.ID,
    TokenType: APP_TOKENATTACH,
    Token: token,
    Expires: time.Now().Unix() + APP_ATTACHEXPIRE,
  }
  if err := store.DB.Create(&accountToken).Error; err != nil {
    ErrResponse(w, http.StatusInternalServerError, err)
    return
  }

  WriteResponse(w, token)
}

