package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// 测试根路径逻辑
func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)

	handler.ServeHTTP(rr, req)

	// 检查状态码
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// 检查返回内容
	expected := "hello from go backend"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// 测试 POST /messages 时 content 为空的情况
func TestMessagesHandler_EmptyContent(t *testing.T) {
	// 构造一个 content 为空的 POST 请求
	req, err := http.NewRequest("POST", "/messages?content=", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(messagesHandler)

	handler.ServeHTTP(rr, req)

	// 此时应该返回 400 Bad Request
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected status 400, got %v", status)
	}

	expected := "content is empty"
	if rr.Body.String() != expected {
		t.Errorf("expected body %v, got %v", expected, rr.Body.String())
	}
}

// 注意：涉及真实数据库操作的测试（如完整的 POST/GET）
// 通常需要启动一个容器化数据库（如 Testcontainers）或使用 Mock 框架。
// 这里仅展示逻辑层的单元测试。
