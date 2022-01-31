package server

import (
	"bytes"
	"forum/internal/database/testdb"
	newErr "forum/internal/error"
	"forum/internal/models"
	"forum/internal/sessions/testsession"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	db := &testdb.TestDB{}
	session := &testsession.TestSession{}
	mysrv := Init(db, session)
	mysrv.router.HandleFunc("/newpost", mysrv.CreatePost)
	srv := httptest.NewServer(&mysrv.router)

	tests := map[string]struct {
		method     string
		post       models.Post
		sessionID  int
		inputBody  []byte
		wantStatus int
		wantError  error
	}{
		// "Succes GET": {
		// 	method:     "GET",
		// 	inputBody:  nil,
		// 	wantStatus: http.StatusOK,
		// },
		"Success POST": {
			method:     "POST",
			post:       models.Post{ID: 1, CategoryID: 1, Username: "User", Title: "Title", Content: "Content"},
			inputBody:  []byte(`{"id":1,"category_id":1,"username":"User","title":"Title","content":"Content"}`),
			wantStatus: http.StatusOK,
		},
		// "Wait BadRequest with nil request body": {
		// 	method:     "POST",
		// 	inputBody:  nil,
		// 	wantStatus: http.StatusBadRequest,
		// },
		// "Wait BadRequest with empty request body": {
		// 	method:     "POST",
		// 	inputBody:  []byte(`{"title":"","content":""}`),
		// 	wantStatus: http.StatusBadRequest,
		// },
		// "Wait MethodNotAllowed": {
		// 	method:     "ERROR",
		// 	inputBody:  nil,
		// 	wantStatus: http.StatusMethodNotAllowed,
		// },
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			db.On("InsertPost", test.post).Return(test.post.ID, test.wantError)

			db.On("CheckCategoryID", test.post.CategoryID).Return(test.wantError)

			req, err := http.NewRequest(test.method, srv.URL+"/newpost", bytes.NewBuffer(test.inputBody))
			assert.Nil(t, err)

			session.On("GetIDByCookie", req).Return(test.sessionID, test.wantError)
			db.On("GetUsernameByID", test.sessionID).Return(test.post.Username, test.wantError)
			resp, err := http.DefaultClient.Do(req)
			assert.Nil(t, err)

			if err != nil {
				assert.Equal(t, test.wantError, err.Error())
				assert.Equal(t, test.wantStatus, resp.StatusCode)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, test.wantError, err.Error())
			}

		})
	}
}

func (s *Server) TestCheckNewPostDatas(t *testing.T) {
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
			if err := s.checkNewPostDatas(&test.inputPost); err != test.wantError {
				t.Errorf("Wait for '%v', but got '%v'", test.wantError, err)
			}
		})
	}
}
