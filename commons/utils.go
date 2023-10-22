package commons

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	gin "github.com/gin-gonic/gin"
	"golang.org/x/sys/unix"
)

// read json file and parse into struct with standard go library
func ReadJSON(filename string, v any) error { //TODO: 是否可改成pointer to any type?
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return err
	}

	// shared lock (read lock) - other transactions can read but cannot write
	if err = unix.Flock(int(file.Fd()), unix.LOCK_SH); err != nil {
		return err
	}
	defer unix.Flock(int(file.Fd()), unix.LOCK_UN) // release a lock

	json.NewDecoder(file).Decode(v)

	return nil
}

func HandleError(c *gin.Context, err error) {
	fmt.Printf("error: %v\n", err)
	c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	return
}
