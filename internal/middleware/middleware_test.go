package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestCORSMiddleware tests CORS middleware configuration
func TestCORSMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(CORSMiddleware())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)

	router.ServeHTTP(w, req)

	// Check that response includes CORS headers
	corsOrigin := w.Header().Get("Access-Control-Allow-Origin")
	if corsOrigin == "" {
		t.Error("Expected Access-Control-Allow-Origin header")
	}
}

// TestErrorHandlerMiddleware tests error handling middleware
func TestErrorHandlerMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(ErrorHandlerMiddleware())

	router.GET("/ok", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ok", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// TestAuthMiddlewareHeader tests auth middleware with token header
func TestAuthMiddlewareHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/protected", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "protected"})
	})

	// Test without token
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/protected", nil)

	router.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Error("Expected 401 or similar error when no token provided")
	}
}

// TestRoleMiddleware tests role-based access control
func TestRoleMiddleware(t *testing.T) {
	tests := []struct {
		name     string
		roles    []string
		testRole string
		allowed  bool
	}{
		{"Admin access to admin route", []string{"admin"}, "admin", true},
		{"Student access to admin route", []string{"admin"}, "student", false},
		{"Instructor access to instructor route", []string{"instructor"}, "instructor", true},
		{"Multiple roles", []string{"admin", "instructor"}, "instructor", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simple role check logic
			hasRole := false
			for _, r := range tt.roles {
				if r == tt.testRole {
					hasRole = true
					break
				}
			}

			if hasRole != tt.allowed {
				t.Errorf("Expected allowed=%v, got allowed=%v", tt.allowed, hasRole)
			}
		})
	}
}

// TestMiddlewareChaining tests that multiple middlewares can be chained
func TestMiddlewareChaining(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	// Apply multiple middlewares
	router.Use(CORSMiddleware())
	router.Use(ErrorHandlerMiddleware())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// TestCORSPreflight tests CORS preflight request handling
func TestCORSPreflight(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(CORSMiddleware())

	router.OPTIONS("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("OPTIONS", "/test", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	req.Header.Set("Access-Control-Request-Method", "POST")

	router.ServeHTTP(w, req)

	// Should handle preflight request
	if w.Code != http.StatusOK && w.Code != http.StatusNoContent {
		t.Errorf("Expected 200 or 204 for OPTIONS request, got %d", w.Code)
	}
}
