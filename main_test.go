package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUploadHandlerStoresPNGAndReturnsShareURL(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 4, 2))
	for y := 0; y < 2; y++ {
		for x := 0; x < 4; x++ {
			img.Set(x, y, color.White)
		}
	}
	var imageData bytes.Buffer
	if err := png.Encode(&imageData, img); err != nil {
		t.Fatal(err)
	}

	var body bytes.Buffer
	form := multipart.NewWriter(&body)
	file, err := form.CreateFormFile("image", "test.png")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := file.Write(imageData.Bytes()); err != nil {
		t.Fatal(err)
	}
	if err := form.Close(); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "http://example.com/upload", &body)
	req.Header.Set("Content-Type", form.FormDataContentType())
	req.Header.Set("X-Forwarded-Proto", "https")
	recorder := httptest.NewRecorder()
	uploadHandler(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d: %s", http.StatusCreated, recorder.Code, recorder.Body.String())
	}
	var response struct {
		ID      string `json:"id"`
		Format  string `json:"format"`
		URL     string `json:"url"`
		Command string `json:"command"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	if response.ID == "" || response.Format != "png" {
		t.Fatalf("unexpected upload response: %+v", response)
	}
	if !strings.HasPrefix(response.URL, "https://example.com/art/") || response.Command != "curl "+response.URL {
		t.Fatalf("unexpected share URL or command: %+v", response)
	}

	uploadedArt.RLock()
	frame, ok := uploadedArt.items[response.ID]
	uploadedArt.RUnlock()
	if !ok || frame == "" {
		t.Fatalf("uploaded image was not stored as terminal art for ID %q", response.ID)
	}
}

func TestImageDimensionsAllowed(t *testing.T) {
	tests := []struct {
		name          string
		width, height int
		want          bool
	}{
		{name: "normal", width: 640, height: 480, want: true},
		{name: "pixel cap", width: 8192, height: 2048, want: true},
		{name: "too many pixels", width: 8192, height: 2049, want: false},
		{name: "dimension cap", width: 8193, height: 1, want: false},
		{name: "zero width", width: 0, height: 1, want: false},
		{name: "zero height", width: 1, height: 0, want: false},
		{name: "negative dimensions", width: -1, height: 1, want: false},
		{name: "overflow-sized dimensions", width: int(^uint(0) >> 1), height: int(^uint(0) >> 1), want: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := imageDimensionsAllowed(test.width, test.height); got != test.want {
				t.Fatalf("imageDimensionsAllowed(%d, %d) = %t, want %t", test.width, test.height, got, test.want)
			}
		})
	}
}
