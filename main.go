package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/hugomd/ascii-live/art"
	"github.com/hugomd/ascii-live/frames"
)

const (
	maxUploadSize = 5 << 20 // 5 MiB
	artWidth      = 80
)

var notFoundMessage = map[string]string{
	"error": "Art not found. Use /list for built-in animations or POST an image to /upload.",
}

var notCurledMessage = map[string]string{
	"error": "You almost ruined a good surprise. Come on, curl it in terminal.",
}

var availableFrames []string

var uploadedArt = struct {
	sync.RWMutex
	items map[string]string
}{items: make(map[string]string)}

func init() {
	for k := range frames.FrameMap {
		availableFrames = append(availableFrames, k)
	}
}

func writeJSON(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		glog.Warningf("writing JSON response: %v", err)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string][]string{"frames": availableFrames}, http.StatusOK)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, notFoundMessage, http.StatusNotFound)
}

func notCurledHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, notCurledMessage, http.StatusExpectationFailed)
}

func curlOnly(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("User-Agent")), "curl")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, map[string]string{"error": "use POST /upload with a PNG or JPEG image"}, http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		writeJSON(w, map[string]string{"error": "upload must be multipart form data with a field named image and be no larger than 5 MiB"}, http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("image")
	if err != nil {
		writeJSON(w, map[string]string{"error": "missing multipart field named image"}, http.StatusBadRequest)
		return
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil || (format != "png" && format != "jpeg") {
		writeJSON(w, map[string]string{"error": "only PNG and JPEG images are supported"}, http.StatusBadRequest)
		return
	}
	frame, err := art.FromImage(img, art.Options{MaxWidth: artWidth})
	if err != nil {
		writeJSON(w, map[string]string{"error": "could not convert image to terminal art"}, http.StatusBadRequest)
		return
	}

	id, err := newArtID()
	if err != nil {
		writeJSON(w, map[string]string{"error": "could not create a shareable art ID"}, http.StatusInternalServerError)
		return
	}
	uploadedArt.Lock()
	uploadedArt.items[id] = frame
	uploadedArt.Unlock()

	baseURL := requestBaseURL(r)
	writeJSON(w, map[string]string{
		"id":      id,
		"format":  format,
		"url":     baseURL + "/art/" + id,
		"command": "curl " + baseURL + "/art/" + id,
	}, http.StatusCreated)
}

func newArtID() (string, error) {
	var bytes [6]byte
	if _, err := io.ReadFull(rand.Reader, bytes[:]); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes[:]), nil
}

func requestBaseURL(r *http.Request) string {
	scheme := "http"
	if forwarded := r.Header.Get("X-Forwarded-Proto"); forwarded != "" {
		scheme = strings.TrimSpace(strings.Split(forwarded, ",")[0])
	}
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}

func streamFrame(w http.ResponseWriter, r *http.Request, getFrame func() string, sleep time.Duration) {
	if !curlOnly(r) {
		notCurledHandler(w, r)
		return
	}
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming is not supported", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.WriteHeader(http.StatusOK)
	for {
		select {
		case <-r.Context().Done():
			glog.Infof("Client stopped listening")
			return
		case <-time.After(sleep):
			fmt.Fprint(w, "\033[2J\033[H"+getFrame()+"\n")
			flusher.Flush()
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["frameSource"]
	frameSource, ok := frames.FrameMap[name]
	if !ok {
		notFoundHandler(w, r)
		return
	}
	index := 0
	streamFrame(w, r, func() string {
		frame := frameSource.GetFrame(index)
		index = (index + 1) % frameSource.GetLength()
		return frame
	}, frameSource.GetSleep())
}

func artHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uploadedArt.RLock()
	frame, ok := uploadedArt.items[id]
	uploadedArt.RUnlock()
	if !ok {
		notFoundHandler(w, r)
		return
	}
	streamFrame(w, r, func() string { return frame }, 700*time.Millisecond)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/list", listHandler).Methods(http.MethodGet)
	r.HandleFunc("/upload", uploadHandler).Methods(http.MethodPost)
	r.HandleFunc("/art/{id}", artHandler).Methods(http.MethodGet)
	r.HandleFunc("/{frameSource}", handler).Methods(http.MethodGet)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	srv := &http.Server{Handler: r, Addr: ":8080"}
	glog.Infof("Serving...")
	glog.Fatal(srv.ListenAndServe())
}
