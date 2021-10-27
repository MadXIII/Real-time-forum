package server

import (
	"bytes"
	"forum/database/testdb"
	newErr "forum/internal/error"
	"forum/models"
	"forum/sessions/testsession"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePost(t *testing.T) {
	mysrv := Init(&testdb.TestDB{}, &testsession.TestSession{})
	mysrv.router.HandleFunc("/newpost", mysrv.CreatePost)
	srv := httptest.NewServer(&mysrv.router)

	tests := map[string]struct {
		method     string
		inputBody  []byte
		wantStatus int
	}{
		"Succes GET": {
			method:     "GET",
			inputBody:  nil,
			wantStatus: http.StatusOK,
		},
		"Succes POST": {
			method:     "POST",
			inputBody:  []byte(`{"title":"1","content":"1"}`),
			wantStatus: http.StatusOK,
		},
		"Wait BadRequest with nil request body": {
			method:     "POST",
			inputBody:  nil,
			wantStatus: http.StatusBadRequest,
		},
		"Wait BadRequest with empty request body": {
			method:     "POST",
			inputBody:  []byte(`{"title":"","content":""}`),
			wantStatus: http.StatusBadRequest,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r, err := http.NewRequest(test.method, srv.URL+"/newpost", bytes.NewBuffer(test.inputBody))
			if err != nil {
				t.Fatal("Problem with TEST", err)
			}
			resp, err := http.DefaultClient.Do(r)
			if err != nil {
				t.Fatal("ERROR TRYING RUN TEST: ")
			}
			if resp.StatusCode != test.wantStatus {
				t.Errorf("Want status: %v, but got: %v", test.wantStatus, resp.StatusCode)
			}
		})
	}
}

func TestCheckNewPostDatas(t *testing.T) {
	tests := map[string]struct {
		inputPost models.Post
		wantError error
	}{
		"Wait error with empty title": {
			inputPost: models.Post{Title: "", Content: "Content"},
			wantError: newErr.ErrPostTitle,
		},
		"Wait error with overflow title": {
			inputPost: models.Post{Title: string([]byte{32: '0'}), Content: "Content"},
			wantError: newErr.ErrPostTitle,
		},
		"Wait error with empty content": {
			inputPost: models.Post{Title: "Some Tittle", Content: ""},
			wantError: newErr.ErrPostContent,
		},
		"Success": {
			inputPost: models.Post{Title: "Title", Content: "Content"},
			wantError: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if err := checkNewPostDatas(test.inputPost); err != test.wantError {
				t.Errorf("Wait for '%v', but got '%v'", test.wantError, err)
			}
		})
	}
}
