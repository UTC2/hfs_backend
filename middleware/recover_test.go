package middleware

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"hfs_backend/common"
)

func newCtx() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c, w
}

func TestRecover_AppError(t *testing.T) {
	c, w := newCtx()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("middleware re-panicked with %v", r)
		}
	}()
	r := Recover(nil)
	r(c)
	func() {
		defer recoverCallback(c)
		panic(common.ErrInternal(errors.New("boom")))
	}()
	if w.Code != 500 {
		t.Errorf("expected 500, got %d", w.Code)
	}
}

// recoverCallback simulates the deferred recover from Recover.
// Inlined here for the test only.
func recoverCallback(c *gin.Context) {
	if r := recover(); r != nil {
		handlePanic(c, r)
	}
}
