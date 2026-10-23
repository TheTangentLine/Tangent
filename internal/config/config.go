// Package config loads Tangent settings from environment variables.
// See docs/configuration.md for the full list.
package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

// Mode is the subcommand the binary runs as. Required settings differ per mode.
type Mode string

const (
	ModeAPI     Mode = "api"
	ModeWorker  Mode = "worker"
	ModeMigrate Mode = "migrate"
)

type Config struct {
	DatabaseURL string
	HTTPAddr    string
	BaseDomain  string
	LogLevel    slog.Level

	OIDCIssuer   string
	OIDCAudience string

	GitHubAppID          int64
	GitHubPrivateKeyPath string
	GitHubOrg            string
	GitOpsRepo           string
	ArgoCDURL            string
	ArgoCDTokenPath      string
	WorkerConcurrency    int

	OpenCostURL  string
	OTLPEndpoint string
}

// Load reads the environment and validates the settings that mode needs.
func Load(mode Mode) (*Config, error) {
	return load(mode, os.Getenv)
}

func load(mode Mode, getenv func(string) string) (*Config, error) {
	var errs []error
	get := func(key, def string) string {
		if v := strings.TrimSpace(getenv(key)); v != "" {
			return v
		}
		return def
	}

	cfg := &Config{
		DatabaseURL:          get("TANGENT_DATABASE_URL", ""),
		HTTPAddr:             get("TANGENT_HTTP_ADDR", ":8080"),
		BaseDomain:           get("TANGENT_BASE_DOMAIN", ""),
		OIDCIssuer:           get("TANGENT_OIDC_ISSUER", ""),
		OIDCAudience:         get("TANGENT_OIDC_AUDIENCE", ""),
		GitHubPrivateKeyPath: get("TANGENT_GITHUB_APP_PRIVATE_KEY_PATH", ""),
		GitHubOrg:            get("TANGENT_GITHUB_ORG", ""),
		GitOpsRepo:           get("TANGENT_GITOPS_REPO", ""),
		ArgoCDURL:            get("TANGENT_ARGOCD_URL", ""),
		ArgoCDTokenPath:      get("TANGENT_ARGOCD_TOKEN_PATH", ""),
		OpenCostURL:          get("TANGENT_OPENCOST_URL", ""),
		OTLPEndpoint:         get("OTEL_EXPORTER_OTLP_ENDPOINT", ""),
	}

	if err := cfg.LogLevel.UnmarshalText([]byte(get("TANGENT_LOG_LEVEL", "info"))); err != nil {
		errs = append(errs, fmt.Errorf("TANGENT_LOG_LEVEL: %w", err))
	}

	n, err := strconv.Atoi(get("TANGENT_WORKER_CONCURRENCY", "4"))
	if err != nil || n < 1 {
		errs = append(errs, errors.New("TANGENT_WORKER_CONCURRENCY: must be a positive integer"))
	}
	cfg.WorkerConcurrency = n

	if v := get("TANGENT_GITHUB_APP_ID", ""); v != "" {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			errs = append(errs, errors.New("TANGENT_GITHUB_APP_ID: must be an integer"))
		}
		cfg.GitHubAppID = id
	}

	require := func(key, val string) {
		if val == "" {
			errs = append(errs, fmt.Errorf("%s is required", key))
		}
	}

	require("TANGENT_DATABASE_URL", cfg.DatabaseURL)

	switch mode {
	case ModeMigrate:
		// Only the database is needed.
	case ModeAPI, ModeWorker:
		require("TANGENT_BASE_DOMAIN", cfg.BaseDomain)
		require("TANGENT_OIDC_ISSUER", cfg.OIDCIssuer)
		require("TANGENT_OIDC_AUDIENCE", cfg.OIDCAudience)
		if mode == ModeAPI {
			require("TANGENT_OPENCOST_URL", cfg.OpenCostURL)
		}
		if mode == ModeWorker {
			if cfg.GitHubAppID == 0 {
				require("TANGENT_GITHUB_APP_ID", get("TANGENT_GITHUB_APP_ID", ""))
			}
			require("TANGENT_GITHUB_APP_PRIVATE_KEY_PATH", cfg.GitHubPrivateKeyPath)
			require("TANGENT_GITHUB_ORG", cfg.GitHubOrg)
			require("TANGENT_GITOPS_REPO", cfg.GitOpsRepo)
			require("TANGENT_ARGOCD_URL", cfg.ArgoCDURL)
			require("TANGENT_ARGOCD_TOKEN_PATH", cfg.ArgoCDTokenPath)
		}
	default:
		errs = append(errs, fmt.Errorf("unknown mode %q", mode))
	}

	if err := errors.Join(errs...); err != nil {
		return nil, err
	}
	return cfg, nil
}
