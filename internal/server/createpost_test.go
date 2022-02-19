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
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	db := &testdb.TestDB{}
	session := &testsession.TestSession{}
	mysrv := Init(db, session)
	mysrv.router.HandleFunc("/newpost", mysrv.CreatePost)
	srv := httptest.NewServer(&mysrv.router)

	tests := map[string]struct {
		method         string
		inputPost      models.Post
		inputCategorys interface{}
		sessionID      int
		inputBody      []byte
		wantStatus     int
		wantError      error
	}{
		"Success GET": {
			method:         http.MethodGet,
			inputCategorys: []models.Categories{{ID: 1, Name: "ALL"}, {ID: 2, Name: "UFC"}},
			wantStatus:     http.StatusOK,
		},
		// "Wait InternalError with wr": {
		// 	method:         http.MethodGet,
		// 	inputCategorys: func() {},
		// 	wantStatus:     http.StatusInternalServerError,
		// },
		"Success POST": {
			method:     http.MethodPost,
			inputPost:  models.Post{ID: 1, CategoryID: 1, Username: "User", Title: "Success", Content: "Content", Timestamp: time.Now().Format("2.Jan.2006, 15:04")},
			inputBody:  []byte(`{"id":1,"category_id":1,"username":"User","title":"Success","content":"Content"}`),
			wantStatus: http.StatusOK,
		},
		"Wait BadRequest with nil request body": {
			method:     http.MethodPost,
			inputPost:  models.Post{ID: 1, CategoryID: 1, Username: "User", Title: "Nil Body", Content: "Content", Timestamp: time.Now().Format("2.Jan.2006, 15:04")},
			inputBody:  nil,
			wantStatus: http.StatusBadRequest,
		},
		"Wait ErrWrongCategory with wrong categoryID": {
			method:     http.MethodPost,
			inputPost:  models.Post{ID: 1, CategoryID: 1, Username: "User", Title: "Nil Body", Content: "Content", Timestamp: time.Now().Format("2.Jan.2006, 15:04")},
			inputBody:  []byte(`{"id":1,"category_id":1,"username":"User","title":"Success","content":"Content"}`),
			wantError:  newErr.ErrWrongCategory,
			wantStatus: http.StatusBadRequest,
		},
		"Wait MethodNotAllowed": {
			method:     http.MethodDelete,
			inputPost:  models.Post{Username: "User"},
			inputBody:  nil,
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			switch test.method {
			case http.MethodPost:
				db.On("InsertPost", &test.inputPost).Return(test.inputPost.ID, test.wantError)
				db.On("GetUsernameByID", test.sessionID).Return(test.inputPost.Username, test.wantError)
				db.On("CheckCategoryID", test.inputPost.ID).Return(test.wantError)
				session.On("GetIDByCookie", mock.Anything).Return(test.sessionID, test.wantError)

				req, err := http.NewRequest(test.method, srv.URL+"/newpost", bytes.NewBuffer(test.inputBody))
				assert.Nil(t, err)

				req.AddCookie(&http.Cookie{
					Name: "session",
				})

				resp, err := http.DefaultClient.Do(req)
				assert.Nil(t, err)

				assert.Equal(t, test.wantStatus, resp.StatusCode)
			case http.MethodGet:
				db.On("GetCategories").Return(test.inputCategorys, test.wantError)

				req, err := http.NewRequest(test.method, srv.URL+"/newpost", bytes.NewBuffer(test.inputBody))
				assert.Nil(t, err)
				resp, err := http.DefaultClient.Do(req)
				assert.Nil(t, err)

				// assert.Equal(t, test.wantError, err.Error())
				assert.Equal(t, test.wantStatus, resp.StatusCode)
			default:
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
