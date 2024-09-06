package main

import "net/http"

type ImageHandler struct {
}

func (b ImageHandler) PostImage(w http.ResponseWriter, r *http.Request)  {}