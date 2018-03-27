package main

import ()

// FileInfo FileInfo
type FileInfo struct {
	UUID string `form:"uuid" json:"uuid"`
	Name string `form:"name" json:"name"`
	Desc string `form:"desc" json:"desc"`
	URL  string `form:"url" json:"url"`
	Hash string `form:"hash" json:"hash"`
}
