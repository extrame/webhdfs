package webhdfs

import (
	"fmt"
)

const (
	ACTION_FS_APPEND = "append"
	ACTION_FS_CREATE = "create"
)

type HttpError struct {
	action      string
	remote_loc  string
	status_code int
}

func (h *HttpError) Error() string {
	return fmt.Sprintf("%s(%s) - File not created.  Server returned status %v", h.action, h.remote_loc, h.status_code)
}

func IsFileNotFoundException(err error) bool {
	if re, ok := err.(RemoteException); ok {
		return re.JavaClassName == "java.io.FileNotFoundException"
	}
	return false
}
