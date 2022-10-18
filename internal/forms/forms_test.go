package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}

}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b")
	if form.Valid() {
		t.Error("got valid when should have been invalid")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b")
	if !form.Valid() {
		t.Error("got invalid when should have been valid")
	}

}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	if form.Has("a") {
		t.Error("show up when there is none")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	if !form.Has("a") {
		t.Error("should have a")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "abc")
	form := New(postedData)
	if form.MinLength("a", 3) {
		t.Error("should report error when less than three")
	}
	if !form.MinLength("b", 3) {
		t.Error("should not report error when greater than or equal to three")
	}
	if form.MinLength("c", 3) {
		t.Error("should report error when item does not exist")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "3170102864@zju.edu.cn")
	form := New(postedData)
	form.IsEmail("a")
	if form.Valid() {
		t.Error("should report error it is not an email")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "3170102864@zju.edu.cn")
	form = New(postedData)
	form.IsEmail("b")
	if !form.Valid() {
		t.Error("should not report error when it is an email")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "3170102864@zju.edu.cn")
	form = New(postedData)
	form.IsEmail("c")
	if form.Valid() {
		t.Error("should report error when item not exist")
	}
}
