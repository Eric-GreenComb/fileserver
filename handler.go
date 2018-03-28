package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerIndex HandlerIndex
func HandlerIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// HandlerHealth HandlerHealth
func HandlerHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}

// HandlerCreateFile HandlerCreateFile
func HandlerCreateFile(c *gin.Context) {

	_uuid := c.PostForm("uuid")
	_name := c.PostForm("name")
	_desc := c.PostForm("desc")
	_url := c.PostForm("url")
	_hash := c.PostForm("hash")

	_fileInfo := FileInfo{
		UUID: _uuid,
		Name: _name,
		Desc: _desc,
		URL:  _url,
		Hash: _hash,
	}

	_json, err := json.Marshal(_fileInfo)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "errmsg": "json error"})
		return
	}

	err = GetBadgerDB().Write(_uuid, _json)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "errmsg": "badger write error"})
		return
	}

	c.JSON(200, gin.H{"errcode": 0, "errmsg": _fileInfo})
}

// HandlerFileInfo HandlerFileInfo
func HandlerFileInfo(c *gin.Context) {

	_uuid := c.Params.ByName("uuid")

	_value, err := GetBadgerDB().Read(_uuid)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "errmsg": err.Error()})
		return
	}

	var _fileInfo FileInfo
	err = json.Unmarshal(_value, &_fileInfo)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "errmsg": "json Unmarshal error"})
		return
	}
	c.JSON(http.StatusOK, _fileInfo)
}

// HandlerGetUploadFile HandlerGetUploadFile
func HandlerGetUploadFile(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

// HandlerUploadFile HandlerUploadFile
func HandlerUploadFile(c *gin.Context) {
	_uuid := c.PostForm("uuid")
	_desc := c.PostForm("desc")

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	_uploadFile := ServerConfig.UploadDir + file.Filename
	if err := c.SaveUploadedFile(file, _uploadFile); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	_hash := md5sum(_uploadFile)

	_url := ServerConfig.UploadURL + file.Filename

	_fileInfo := FileInfo{
		UUID: _uuid,
		Name: file.Filename,
		Desc: _desc,
		URL:  _url,
		Hash: _hash,
	}

	_json, err := json.Marshal(_fileInfo)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "errmsg": "json error"})
		return
	}

	err = GetBadgerDB().Write(_uuid, _json)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "errmsg": "badger write error"})
		return
	}

	c.JSON(200, gin.H{"errcode": 0, "errmsg": _fileInfo})
}
