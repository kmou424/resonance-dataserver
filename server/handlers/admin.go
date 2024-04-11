package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gookit/goutil/fsutil"
	"github.com/kmou424/resonance-dataserver/cmd/cli"
	"github.com/kmou424/resonance-dataserver/database/model"
	"github.com/kmou424/resonance-dataserver/database/repositories"
	"github.com/kmou424/resonance-dataserver/server/errors"
	"net/http"
	"os"
	"strings"
	"sync"
)

func queryUser(c *gin.Context) string {
	user := c.Query("user")
	if user == "" {
		c.Abort()
		panic(errors.BadRequest(fmt.Sprintf("user not provided")))
	}
	return user
}

var AddUUID gin.HandlerFunc = func(c *gin.Context) {
	user := queryUser(c)

	if strings.Contains(user, "admin") {
		c.Abort()
		panic(errors.BadRequest(fmt.Sprintf("can't create admin user")))
	}

	if repositories.AuthKey.HasUser(user) {
		c.Abort()
		panic(errors.BadRequest(fmt.Sprintf("user already exists: %s", user)))
	}

	authKey := model.AuthKey{
		UUID: uuid.New().String(),
		User: user,
	}

	repositories.AuthKey.Create(&authKey)
	syncAuthKeysToFilesystem(c)

	c.JSON(http.StatusOK, authKey)
}

var UpdateUUID gin.HandlerFunc = func(c *gin.Context) {
	user := queryUser(c)

	authKey := repositories.AuthKey.FindByUser(user)
	if authKey == nil {
		c.Abort()
		panic(errors.BadRequest(fmt.Sprintf("invalid user")))
	}

	authKey.UUID = uuid.New().String()

	repositories.AuthKey.UpdateByUser(authKey)
	syncAuthKeysToFilesystem(c)

	c.JSON(http.StatusOK, authKey)
}

var DeleteUUID gin.HandlerFunc = func(c *gin.Context) {
	user := queryUser(c)

	if strings.Contains(user, "admin") {
		c.Abort()
		panic(errors.BadRequest(fmt.Sprintf("admin uuid can't be deleted")))
	}

	authKey := repositories.AuthKey.FindByUser(user)
	if authKey == nil {
		c.Abort()
		panic(errors.BadRequest(fmt.Sprintf("invalid user")))
	}

	repositories.AuthKey.DeleteByUUID(authKey)
	syncAuthKeysToFilesystem(c)

	c.JSON(http.StatusOK, authKey)
}

var authKeysFileMutex = sync.Mutex{}

func syncAuthKeysToFilesystem(c *gin.Context) {
	authKeysFileMutex.Lock()
	defer authKeysFileMutex.Unlock()
	var err error

	defer func() {
		if err != nil {
			c.Abort()
			panic(errors.InternalServerError(fmt.Sprintf("error syncing auth keys to filesystem: %s", err)))
		}
	}()

	authKeys := repositories.AuthKey.FindAll()
	marshaled, err := json.MarshalIndent(authKeys, "", "    ")
	if err != nil {
		return
	}
	err = fsutil.WriteFile(cli.AuthKeysFile, marshaled, os.ModePerm)
	if err != nil {
		return
	}
}
