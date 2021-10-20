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
		inputBody  []byte
		wantStatus int
	}{
		"Wait error with nil request body": {
			inputBody:  nil,
			wantStatus: http.StatusBadRequest,
		},
		"Succes": {
			inputBody:  []byte(`{"title":"1","content":"1"}`),
			wantStatus: http.StatusOK,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {})
	}

	wantStatus := http.StatusOK

	resp, err := http.Get(srv.URL + "/newpost")
	if err != nil {
		t.Fatal("ERROR TRYING RUN TEST: ", err)
	}

	if resp.StatusCode != wantStatus {
		t.Errorf("Want status: %v, but got: %v", wantStatus, resp.StatusCode)
	}

	//--------------------------------------
	//need to refactor
	buf := bytes.NewBuffer([]byte(`{"title":"1","content":"1"}`))
	resp, err = http.Post(srv.URL+"/newpost", "", buf)
	if err != nil {
		t.Fatal("ERROR TRYING RUN TEST: ", err)
	}
	if resp.StatusCode != wantStatus {
		t.Errorf("Want status: %v but got: %v", wantStatus, resp.StatusCode)
	}
	//--------------------------------------

	wantStatus = http.StatusBadRequest

	resp, err = http.Post(srv.URL+"/newpost", "", nil)
	if err != nil {
		t.Fatal("ERROR TRYING RUN TEST: ", err)
	}
	if resp.StatusCode != wantStatus {
		t.Errorf("Want status: %v, but got: %v", wantStatus, resp.StatusCode)
	}

}

func TestCheckNewPostDatas(t *testing.T) {
	tests := map[string]struct {
		inputPost models.Post
		wantError error
	}{
		"Wait error with empty title": {
			inputPost: models.Post{Title: "", Content: "some content"},
			wantError: newErr.ErrPostTitle,
		},
		"Wait error with overflow title": {
			inputPost: models.Post{Title: string([]byte{32: '0'}), Content: "some content"},
			wantError: newErr.ErrPostTitle,
		},
		"Wait error with empty content": {
			inputPost: models.Post{Title: "Some Tittle", Content: ""},
			wantError: newErr.ErrPostContent,
		},
		"Success": {
			inputPost: models.Post{Title: "Correct Title", Content: ""},
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
