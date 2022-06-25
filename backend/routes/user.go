package routes

import (
  "encoding/json"
  "log"
  "net/http"

  "github.com/julienschmidt/httprouter"
  "host.local/todo-app/backend/models"
)

type UpdateUsernamePayload struct {
  NewUsername string `json:"new_username"`
}

func updateUsername(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  userId := req.Context().Value(idKey).(int64)
  username := req.Context().Value(usernameKey).(string)

  var payload UpdateUsernamePayload

  err := json.NewDecoder(req.Body).Decode(&payload)
  if err != nil {
    jsonMessage(w, http.StatusBadRequest, "Failed to decode JSON")
    return
  }

  exists, err := dbConn.UsernameExists(payload.NewUsername)
  if err != nil {
    log.Println(err)
    jsonMessage(w, http.StatusInternalServerError, "Internal Server Error")
    return
  }

  if exists {
    log.Println("username exists:", payload.NewUsername)
    jsonMessage(w, http.StatusBadRequest, "Username already in use")
    return
  }

  var user = models.User{
    Id:       userId,
    Username: username,
  }

  err = dbConn.UpdateUser(&user, payload.NewUsername)
  if err != nil {
    log.Println(err)
    jsonMessage(w, http.StatusInternalServerError, "Failed to update user")
    return
  }

  token, err := authInstance.CreateJWT(user.Id, payload.NewUsername)
  if err != nil {
    log.Println(err)
    jsonMessage(w, http.StatusInternalServerError, "Failed to create token")
    return
  }

  jsonMessage(w, http.StatusOK, token)
}

type UpdatePasswordPayload struct {
  models.User
  NewPassword string `json:"new_password"`
}

func updatePassword(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  userId := req.Context().Value(idKey).(int64)
  username := req.Context().Value(usernameKey).(string)

  var payload UpdatePasswordPayload
  payload.Id = userId

  err := json.NewDecoder(req.Body).Decode(&payload)
  if err != nil {
    log.Println(err)
    jsonMessage(w, http.StatusBadRequest, "Failed to decode JSON")
    return
  }

  user := payload.User
  user.Id = userId
  user.Username = username
  
  err = dbConn.VerifyUser(&user)
  if err != nil {
    log.Println(err)
    jsonMessage(w, http.StatusBadRequest, "Unauthorized operation")
    return
  }

  hash, err := models.HashPassword(payload.NewPassword)
  if err != nil {
    log.Println(err)
    jsonMessage(w, http.StatusInternalServerError, "Failed to process request")
    return
  }

  user.Password = hash

  err = dbConn.UpdatePassword(&user)
  if err != nil {
    log.Println(err)
    jsonMessage(w, http.StatusInternalServerError, "Failed to update credentials")
    return
  }

  jsonMessage(w, http.StatusOK, "Password updated")
}

func deleteUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  userId := req.Context().Value(idKey).(int64)
  username := req.Context().Value(usernameKey).(string)

  var user models.User
  user.Id = userId
  user.Username = username

  err := json.NewDecoder(req.Body).Decode(&user)
  if err != nil {
    log.Println(err)
    jsonMessage(w, http.StatusBadRequest, "Failed to decode JSON")
    return
  }

  err = dbConn.VerifyUser(&user)
  if err != nil {
    log.Println(err)
    jsonMessage(w, http.StatusBadRequest, "Unauthorized operation")
    return
  }

  err = dbConn.DeleteUser(&user)
  if err != nil {
    log.Println(err)
    jsonMessage(w, http.StatusInternalServerError, "Failed to delete user")
    return
  }

  jsonMessage(w, http.StatusOK, "User deleted")
}
