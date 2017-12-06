package controllers

import (
	"github.com/devfeel/dotweb"
	"fmt"
	"dotapp_server/const"
	"dotweb-start/util/file"
	"time"
	"strings"
	"strconv"
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

//多张图片上传
func (this *PublicFunc) FileUploadMore(ctx dotweb.Context) (err error) {
	fileMap, err:=ctx.Request().FormFiles()
	l := len(fileMap)
	if l==0 {
		this.Respone(ctx,_const.CODE_FIAL,"请上传图片" + err.Error(),nil)
		return
	}
	 retString:= make([]string,len(fileMap))
	if err!= nil{
		this.Respone(ctx,_const.CODE_FIAL,"FormFile error " + err.Error(),nil)
		return
	}else {
		 i := 0
		for _, upload:=range fileMap{
			if upload.Size()>1024*10 {
				this.Respone(ctx,_const.CODE_FIAL,"文件太大，请重新上传",nil)
				return
			}
			filepath:="../static/upload/" +strconv.FormatInt(time.Now().Unix(),10)+upload.FileName()
			_, err = upload.SaveFile(filepath)
			if err != nil {
				this.Respone(ctx,_const.CODE_FIAL,"SaveFile error " + err.Error(),nil)
				return
			} else {
				retString[i] = strings.Replace(filepath,"../","",1)
				i++
			}
		}
	}
	this.Respone(ctx,_const.CODE_SUCCESS,"上传成功 " ,retString)
	return
}
