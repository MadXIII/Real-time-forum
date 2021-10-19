package server

import (
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

	wantStatus := http.StatusOK

	resp, err := http.Get(srv.URL + "/newpost")
	if err != nil {
		t.Fatal("ERROR TRYING RUN TEST: ", err)
	}

	if resp.StatusCode != wantStatus {
		t.Errorf("Want status: %v, but got: %v", wantStatus, resp.StatusCode)
	}

	resp, err = http.Post(srv.URL+"/newpost", "", nil)
	if err != nil {
		t.Fatal("ERROR TRYING RUN TEST: ", err)
	}

	if resp.StatusCode != wantStatus {
		t.Errorf("Want status: %v but got: %v", wantStatus, resp.StatusCode)
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
			inputPost: models.Post{Title: "Correct Title", Content: "Correct Content"},
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
