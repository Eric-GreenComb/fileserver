package main

import ()

// FileInfo FileInfo
type FileInfo struct {
	UUID string `form:"uuid" json:"uuid"`
	Name string `form:"name" json:"name"`
	Desc string `form:"desc" json:"desc"`
	File string `form:"file" json:"file"`
	Path string `form:"path" json:"path"`
	URL  string `form:"url" json:"url"`
	Hash string `form:"hash" json:"hash"`
}
