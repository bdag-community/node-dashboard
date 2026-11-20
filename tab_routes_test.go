package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemovedStatusTabsReturnNotFound(t *testing.T) {
	app := &App{templates: parseTemplates()}
	for _, path := range []string{
		"/tabs/node-sync",
		"/tabs/pool-status",
		"/tabs/warnings",
		"/tabs/asic-fleet",
		"/tabs/stratum-runtime",
		"/tabs/pool-management",
		"/tabs/chain-dag",
		"/tabs/pool-earnings",
		"/tabs/global",
	} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		rec := httptest.NewRecorder()

		app.handleTab(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Fatalf("%s status = %d, want 404", path, rec.Code)
		}
	}
}

func TestNodeDashboardTabsRemainAvailable(t *testing.T) {
	app := &App{
		templates: parseTemplates(),
		nodeLogProvider: func(context.Context, int) string {
			return "node line"
		},
		poolLogProvider: func(context.Context, int) string {
			return "pool line"
		},
	}
	for _, path := range []string{"/tabs/overview", "/tabs/peers", "/tabs/logs", "/tabs/node-logs", "/tabs/pool-logs"} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		rec := httptest.NewRecorder()

		app.handleTab(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("%s status = %d, want 200", path, rec.Code)
		}
	}
}
