package config

import (
	"log/slog"
	"strings"
	"testing"
)

func env(m map[string]string) func(string) string {
	return func(k string) string { return m[k] }
}

var apiEnv = map[string]string{
	"TANGENT_DATABASE_URL":  "postgres://localhost/tangent",
	"TANGENT_BASE_DOMAIN":   "localtest.me",
	"TANGENT_OIDC_ISSUER":   "http://keycloak.localtest.me/realms/tangent",
	"TANGENT_OIDC_AUDIENCE": "tangent-api",
	"TANGENT_OPENCOST_URL":  "http://opencost:9003",
}

func TestLoadDefaults(t *testing.T) {
	cfg, err := load(ModeAPI, env(apiEnv))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.HTTPAddr != ":8080" {
		t.Errorf("HTTPAddr = %q, want :8080", cfg.HTTPAddr)
	}
	if cfg.WorkerConcurrency != 4 {
		t.Errorf("WorkerConcurrency = %d, want 4", cfg.WorkerConcurrency)
	}
	if cfg.LogLevel != slog.LevelInfo {
		t.Errorf("LogLevel = %v, want info", cfg.LogLevel)
	}
}

func TestLoadMigrateNeedsOnlyDatabase(t *testing.T) {
	_, err := load(ModeMigrate, env(map[string]string{"TANGENT_DATABASE_URL": "postgres://x"}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLoadWorkerReportsAllMissing(t *testing.T) {
	_, err := load(ModeWorker, env(apiEnv))
	if err == nil {
		t.Fatal("expected error")
	}
	for _, key := range []string{
		"TANGENT_GITHUB_APP_ID",
		"TANGENT_GITHUB_APP_PRIVATE_KEY_PATH",
		"TANGENT_GITHUB_ORG",
		"TANGENT_GITOPS_REPO",
		"TANGENT_ARGOCD_URL",
		"TANGENT_ARGOCD_TOKEN_PATH",
	} {
		if !strings.Contains(err.Error(), key) {
			t.Errorf("error does not mention %s: %v", key, err)
		}
	}
}

func TestLoadInvalidValues(t *testing.T) {
	tests := map[string]string{
		"TANGENT_LOG_LEVEL":          "loud",
		"TANGENT_WORKER_CONCURRENCY": "0",
	}
	for key, val := range tests {
		t.Run(key, func(t *testing.T) {
			m := map[string]string{}
			for k, v := range apiEnv {
				m[k] = v
			}
			m[key] = val
			_, err := load(ModeAPI, env(m))
			if err == nil || !strings.Contains(err.Error(), key) {
				t.Errorf("expected error mentioning %s, got %v", key, err)
			}
		})
	}
}

func TestLoadUnknownMode(t *testing.T) {
	if _, err := load("nope", env(apiEnv)); err == nil {
		t.Fatal("expected error")
	}
}
