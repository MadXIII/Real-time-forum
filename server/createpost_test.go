package server

import (
	"forum/database/testdb"
	newErr "forum/internal/error"
	"forum/models"
	"forum/sessions/testsession"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePost(t *testing.T) {
	mysrv := Init(&testdb.TestDB{}, &testsession.TestSession{})
	mysrv.Conf()
	srv := httptest.NewServer(&mysrv.router)
	// srv.URL

	wantStatus := http.StatusOK
	resp, err := http.Post(srv.URL+"/newpost", "", nil)
	if err != nil {
		t.Fatal("ERROR TRYING RUN TEST: ", err)
	}

	if resp.StatusCode != wantStatus {
		t.Errorf("want status: %v but got %v", wantStatus, resp.StatusCode)
	}
	log.Println(srv.URL)
	// tests := map[string]struct {
	// 	post       models.Post
	// 	wantStatus error
	// 	wantBody   []byte
	// }{}

	// for name, test := range tests {
	// 	t.Run(name, func(t *testing.T) {
	// 		if err := checkNewPostDatas(test.inputPost); err != test.wantErr {
	// 			t.Errorf("Wait for '%v', but got '%v'", test.wantErr, err)
	// 		}
	// 	})
	// }
}

func TestCheckNewPostDatas(t *testing.T) {
	tests := map[string]struct {
		inputPost models.Post
		wantErr   error
	}{
		"Wait error with empty title": {
			inputPost: models.Post{Title: "", Content: "some content"},
			wantErr:   newErr.ErrPostTitle,
		},
		"Wait error with overflow title": {
			inputPost: models.Post{Title: string([]byte{32: '0'}), Content: "some content"},
			wantErr:   newErr.ErrPostTitle,
		},
		"Wait error with empty content": {
			inputPost: models.Post{Title: "Some Tittle", Content: ""},
			wantErr:   newErr.ErrPostContent,
		},
		"Success": {
			inputPost: models.Post{Title: "Correct Title", Content: "Correct Content"},
			wantErr:   nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if err := checkNewPostDatas(test.inputPost); err != test.wantErr {
				t.Errorf("Wait for '%v', but got '%v'", test.wantErr, err)
			}
		})
	}
}
