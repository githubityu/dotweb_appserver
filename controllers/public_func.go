package controllers

import (
	"github.com/devfeel/dotweb"
	"fmt"
	"dotapp_server/const"
	"dotweb-start/util/file"
	"time"
	"strings"
	"strconv"
	"os"
	"io"
)

type PublicFunc struct {
	BaseController
}
//单张图片上传
func (this *PublicFunc) FileUpload(ctx dotweb.Context) (err error) {
	upload, err := ctx.Request().FormFile("file")
	if err != nil {
		this.Respone(ctx,_const.CODE_FIAL,"FormFile error " + err.Error(),nil)
		return
	} else {
		filepath:="../static/upload/" + strconv.FormatInt(time.Now().Unix(),10)+file.GetFileExt(upload.FileName())
		_, err = upload.SaveFile(filepath)
		if err != nil {
			this.Respone(ctx,_const.CODE_FIAL,"SaveFile error " + err.Error(),nil)
			return
		} else {
			this.Respone(ctx,_const.CODE_SUCCESS,"SaveFile success || " + upload.FileName() + " || " + upload.GetFileExt() + " || " + fmt.Sprint(upload.Size()),strings.Replace(filepath,"../","",0))
			return
		}
	}

}

//单张图片上传
func (this *PublicFunc) FileUploadMore(ctx dotweb.Context) (err error) {
	fhs:= ctx.Request().MultipartForm.File["file[]"]
	l := len(fhs)
	optionDirs := make([]string, l)
	for i := 0; i < l; i++ {
		file, err := fhs[i].Open()
		if err != nil {
			this.Respone(ctx,_const.CODE_FIAL,"SaveFile error " + err.Error(),nil)
		}
		filename := fhs[i].Filename
		f, err := os.OpenFile("static/upload/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			this.Respone(ctx,_const.CODE_FIAL,"SaveFile error " + err.Error(),nil)
		}
		defer f.Close()
		io.Copy(f, file)
		optionDirs = append(optionDirs, filename)
	}
	this.Respone(ctx,_const.CODE_SUCCESS,"上传成功",optionDirs)
	return
}
