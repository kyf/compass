package admin

import (
	"fmt"
	"github.com/kyf/compass/data"
	"github.com/martini-contrib/sessions"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	APP_UPLOAD_MAX_SIZE int64  = 1024 * 1024 * 20
	APP_STORE_DIR       string = "/work/compass/app/"
)

func AppAdd(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger) {
	r.ParseMultipartForm(APP_UPLOAD_MAX_SIZE)
	form := r.MultipartForm
	tid := form.Value["app_id"]
	tname := form.Value["app_name"]
	ticon := form.File["app_icon"]
	tapk := form.File["app_apk"]
	tdesc := form.Value["app_desc"]

	id := ""
	name := ""
	desc := ""
	icon := ""
	apk := ""

	if len(tid) > 0 {
		id = tid[0]
	}

	if len(tname) > 0 {
		name = tname[0]
	}

	if len(tdesc) > 0 {
		desc = tdesc[0]
	}

	if len(ticon) > 0 {
		fpicon, err := ticon[0].Open()
		if err != nil {
			logger.Printf("appadd err:%v", err)
		} else {
			defer fpicon.Close()
			icon_content, err := ioutil.ReadAll(fpicon)
			if err != nil {
				logger.Printf("read file err:%v", err)
			} else {

				iconpath := fmt.Sprintf("%s%s", APP_STORE_DIR, ticon[0].Filename)
				nf, err := os.Create(iconpath)
				if err != nil {
					logger.Printf("create new file err:%v", err)
				} else {
					defer nf.Close()
					nf.Write(icon_content)
					icon = ticon[0].Filename
				}
			}
		}
	}

	if len(tapk) > 0 {
		fpapk, err := tapk[0].Open()
		if err != nil {
			logger.Printf("appadd err:%v", err)
		} else {
			defer fpapk.Close()
			apk_content, err := ioutil.ReadAll(fpapk)
			if err != nil {
				logger.Printf("read file err:%v", err)
			} else {

				apkpath := fmt.Sprintf("%s%s", APP_STORE_DIR, tapk[0].Filename)
				nf, err := os.Create(apkpath)
				if err != nil {
					logger.Printf("create new file err:%v", err)
				} else {
					defer nf.Close()
					nf.Write(apk_content)
					apk = tapk[0].Filename
				}
			}
		}
	}

	app := &data.App{Logger: logger}
	var err error
	if strings.EqualFold("", id) {
		err = app.Add(name, icon, desc, apk)
	} else {
		nid, err := strconv.Atoi(id)
		if err == nil {
			err = app.Update(nid, name, icon, desc, apk)
		}
	}

	var resdata []byte
	if err != nil {
		resdata = jsonResponse(map[string]interface{}{"status": false, "msg": fmt.Sprintf("%v", err)})
	} else {
		resdata = jsonResponse(map[string]interface{}{"status": true, "msg": "success"})
	}
	w.Write([]byte(fmt.Sprintf("<script type='text/javascript'>parent.callback(%s)</script>", string(resdata))))
}

func AppId(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger) {
	r.ParseForm()
	sid := r.PostForm.Get("id")

	if strings.EqualFold("", sid) {
		w.Write(jsonResponse(map[string]interface{}{"status": false, "msg": "id is empty"}))
		return
	}

	app := &data.App{Logger: logger}
	app = app.Detail(sid)

	if app == nil {
		w.Write(jsonResponse(map[string]interface{}{"status": false, "msg": "app is null"}))
	} else {
		w.Write(jsonResponse(map[string]interface{}{"status": true, "msg": "success", "data": *app}))
	}
}

func AppList(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger) {
	r.ParseForm()

	apps := &data.App{Logger: logger}
	ds := apps.List()
	var result map[string]interface{} = map[string]interface{}{
		"status": true,
		"data":   ds,
	}
	if ds == nil {
		result = map[string]interface{}{
			"status": true,
			"data":   []string{},
		}
	}
	w.Write(jsonResponse(result))

}

func AppDel(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger) {
	r.ParseForm()
	sid := r.PostForm.Get("id")

	if strings.EqualFold("", sid) {
		w.Write(jsonResponse(map[string]interface{}{"status": false, "msg": "id is empty"}))
	} else {
		id, err := strconv.Atoi(sid)
		if err != nil {
			logger.Printf("AppDel ERR:%v", err)
			w.Write(jsonResponse(map[string]interface{}{"status": false, "msg": fmt.Sprintf("%v", err)}))
		} else {
			app := &data.App{Logger: logger}
			err = app.Remove(id)
			if err != nil {
				w.Write(jsonResponse(map[string]interface{}{"status": false, "msg": fmt.Sprintf("%v", err)}))
			} else {
				w.Write(jsonResponse(map[string]interface{}{"status": true, "msg": "success"}))
			}
		}
	}

}

func init() {
	ActionHandlers["app/add"] = AppAdd
	ActionHandlers["app/id"] = AppId
	ActionHandlers["app/list"] = AppList
	ActionHandlers["app/delete"] = AppDel
}
