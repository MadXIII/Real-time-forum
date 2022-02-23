package server

import (
	"bytes"
	"errors"
	"forum/internal/database"
	"forum/internal/database/testdb"
	newErr "forum/internal/error"
	"forum/internal/models"
	s "forum/internal/sessions"
	"forum/internal/sessions/testsession"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePost(t *testing.T) {
	type args struct {
		post           models.Post
		sessionID      int
		wantError      error
		inputCategorys interface{}
	}
	tests := []struct {
		name       string
		args       args
		method     string
		inputBody  []byte
		wantStatus int
		callback   func(args) (database.Repository, s.Repository)
	}{
		{
			name: "Success GET",
			args: args{
				inputCategorys: []models.Categories{{ID: 1, Name: "All"}},
			},
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				db.On("GetCategories").Return(a.inputCategorys, a.wantError)
				return db, session
			},
		},
		{
			name: "Wait GetCategories error",
			args: args{
				inputCategorys: []models.Categories{{ID: 1, Name: "All"}},
				wantError:      errors.New("GetCategories, Scan: sql: expected 2 destination arguments in Scan, not 1"),
			},
			method:     http.MethodGet,
			wantStatus: http.StatusInternalServerError,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				db.On("GetCategories").Return(a.inputCategorys, a.wantError)
				return db, session
			},
		},
		{
			name: "Success POST",
			args: args{
				post: models.Post{ID: 1, CategoryID: 1, Username: "User", Title: "Success", Content: "Content", Timestamp: time.Now().Format("2.Jan.2006, 15:04")},
			},
			method:     http.MethodPost,
			inputBody:  []byte(`{"id":1,"category_id":1,"username":"User","title":"Success","content":"Content"}`),
			wantStatus: http.StatusOK,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				db.On("CheckCategoryID", a.post.CategoryID).Return(a.wantError).Once()
				session.On("GetIDByCookie", mock.Anything).Return(a.sessionID, a.wantError).Once()
				db.On("GetUsernameByID", a.sessionID).Return(a.post.Username, a.wantError).Once()
				db.On("InsertPost", &a.post).Return(a.post.ID, a.wantError).Once()
				return db, session
			},
		},
		{
			name: "Wait BadRequest with nil request body",
			args: args{
				post: models.Post{ID: 1, CategoryID: 1, Username: "User", Title: "Nil Body", Content: "Content", Timestamp: time.Now().Format("2.Jan.2006, 15:04")},
			},
			method:     http.MethodPost,
			inputBody:  nil,
			wantStatus: http.StatusBadRequest,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				return db, session
			},
		},
		{
			name: "Wait MethodNotAllowed",
			args: args{
				post: models.Post{Username: "User"},
			},
			method:     http.MethodDelete,
			inputBody:  nil,
			wantStatus: http.StatusMethodNotAllowed,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				return db, session
			},
		},
		{
			name: "Wait ErrWrongCategory with wrong categoryID",
			args: args{
				post:      models.Post{ID: 1, CategoryID: -1, Username: "User", Title: "Nil Body", Content: "Content", Timestamp: time.Now().Format("2.Jan.2006, 15:04")},
				wantError: newErr.ErrWrongCategory,
			},
			method:     http.MethodPost,
			inputBody:  []byte(`{"id":1,"category_id":-1,"username":"User","title":"Success","content":"Content"}`),
			wantStatus: http.StatusBadRequest,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				db.On("CheckCategoryID", a.post.CategoryID).Return(a.wantError).Once()
				return db, session
			},
		},
		{
			name: "Wait GetUsernameByID error",
			args: args{
				post:      models.Post{ID: 1, CategoryID: 1, Username: "User", Title: "Success", Content: "Content", Timestamp: time.Now().Format("2.Jan.2006, 15:04")},
				sessionID: -1,
				wantError: errors.New("handleCreatePost, getUsernameByCookie: GetUsernameByID, Scan: sql: no rows in result set"),
			},
			inputBody:  []byte(`{"id":1,"category_id":1,"username":"User","title":"Success","content":"Content"}`),
			method:     http.MethodPost,
			wantStatus: http.StatusInternalServerError,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				db.On("CheckCategoryID", a.post.CategoryID).Return(nil).Once()
				session.On("GetIDByCookie", mock.Anything).Return(a.sessionID, nil).Once()
				db.On("GetUsernameByID", a.sessionID).Return(a.post.Username, a.wantError).Once()
				return db, session
			},
		},
		{
			name: "Wait ErrPostTitle error",
			args: args{
				post:      models.Post{ID: 1, CategoryID: 1, Username: "User", Title: "", Content: "Content", Timestamp: time.Now().Format("2.Jan.2006, 15:04")},
				wantError: newErr.ErrPostTitle,
			},
			method:     http.MethodPost,
			inputBody:  []byte(`{"id":1,"category_id":1,"username":"User","title":"","content":"Content"}`),
			wantStatus: http.StatusBadRequest,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				db.On("CheckCategoryID", a.post.CategoryID).Return(nil).Once()
				return db, session
			},
		},
		{
			name: "Wait ErrPostContent error",
			args: args{
				post:      models.Post{ID: 1, CategoryID: 1, Username: "User", Title: "Title", Content: "", Timestamp: time.Now().Format("2.Jan.2006, 15:04")},
				wantError: newErr.ErrPostContent,
			},
			method:     http.MethodPost,
			inputBody:  []byte(`{"id":1,"category_id":1,"username":"User","title":"Title","content":""}`),
			wantStatus: http.StatusBadRequest,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				db.On("CheckCategoryID", a.post.CategoryID).Return(nil).Once()
				return db, session
			},
		},
		{
			name: "Wait InsertPost error",
			args: args{
				post:      models.Post{ID: 1, CategoryID: 1, Username: "User", Title: "Success", Content: "Content", Timestamp: time.Now().Format("2.Jan.2006, 15:04")},
				wantError: errors.New("handleCreatePost, InsertPost: InsertPost, Prepare: 5 values for 6 columns"),
			},
			method:     http.MethodPost,
			inputBody:  []byte(`{"id":1,"category_id":1,"username":"User","title":"Success","content":"Content"}`),
			wantStatus: http.StatusInternalServerError,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				db.On("CheckCategoryID", a.post.CategoryID).Return(nil).Once()
				session.On("GetIDByCookie", mock.Anything).Return(a.sessionID, nil).Once()
				db.On("GetUsernameByID", a.sessionID).Return(a.post.Username, nil).Once()
				db.On("InsertPost", &a.post).Return(a.post.ID, a.wantError).Once()
				return db, session
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, session := test.callback(test.args)
			mysrv := Init(db, session)
			mysrv.router.HandleFunc("/newpost", mysrv.CreatePost)

			req, err := http.NewRequest(test.method, "/newpost", bytes.NewBuffer(test.inputBody))
			assert.Nil(t, err)

			req.AddCookie(&http.Cookie{
				Name: "session",
			})

			recorder := httptest.NewRecorder()
			mysrv.CreatePost(recorder, req)
			resp := recorder.Result()

			assert.Equal(t, test.wantStatus, resp.StatusCode)
		})
	}
}
