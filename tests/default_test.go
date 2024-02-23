package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
	_ "uas-api-pegawai/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestCreatePegawaiEndpoint is a sample to run an endpoint test for creating Pegawai
func TestCreatePegawaiEndpoint(t *testing.T) {
	// Prepare the JSON payload for creating Pegawai
	payload := []byte(`{
		"nama": "John Doe",
		"alamat": "123 Main Street",
		"jenis_kelamin": {"id": 1},
		"agama": {"id": 1},
		"status": {"id": 1}
	}`)

	r, _ := http.NewRequest("POST", "/pegawai", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestCreatePegawaiEndpoint", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Create Pegawai Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		// Add more assertions based on your specific requirements
	})
}
