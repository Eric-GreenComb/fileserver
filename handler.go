package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/banerwai/gommon/uuid"
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
	// _uuid := c.PostForm("uuid")
	_uuid := UUID()
	fmt.Println("upload : " + _uuid)
	_name := c.PostForm("name")
	_desc := c.PostForm("desc")

	file, err := c.FormFile("file1")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	_uploadFile := ServerConfig.UploadDir + file.Filename

	fmt.Println("upload : " + _uploadFile)
	if err := c.SaveUploadedFile(file, _uploadFile); err != nil {
		fmt.Println(err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	_hash := md5sum(_uploadFile)

	_url := ServerConfig.UploadURL + file.Filename

	_fileInfo := FileInfo{
		UUID: _uuid,
		Name: _name,
		Desc: _desc,
		File: file.Filename,
		Path: ServerConfig.UploadURL,
		URL:  _url,
		Hash: _hash,
	}

	_json, err := json.Marshal(_fileInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errcode": 1, "errmsg": "json error"})
		return
	}

	err = GetBadgerDB().Write(_uuid, _json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errcode": 1, "errmsg": "badger write error"})
		return
	}

	c.JSON(http.StatusOK, _fileInfo)
}

// HandlerGetTest HandlerGetTest
func HandlerGetTest(c *gin.Context) {
	c.HTML(http.StatusOK, "form-url.html", nil)
}

// HandlerTest HandlerTest
func HandlerTest(c *gin.Context) {

	_uuid := c.PostForm("uuid")
	// _uuid := "test"

	fmt.Println("test : " + _uuid)

	c.JSON(200, gin.H{"errcode": 0, "errmsg": _uuid})
}

// UUID Google UUID
func UUID() string {
	return uuid.UUID()
}
