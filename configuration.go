/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package nsxt

import (
	"net/http"
)

const ContextOAuth2 int = 1
const ContextBasicAuth int = 2
const ContextAccessToken int = 3
const ContextAPIKey int = 4

type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type APIKey struct {
	Key    string
	Prefix string
}

type ClientRetriesConfiguration struct {
	MaxRetries      int
	RetryMinDelay   int // milliseconds
	RetryMaxDelay   int // milliseconds
	RetryOnStatuses []int
}

type Configuration struct {
	BasePath                     string `json:"basePath,omitempty"`
	Host                         string `json:"host,omitempty"`
	Scheme                       string `json:"scheme,omitempty"`
	UserName                     string
	Password                     string
	RemoteAuth                   bool
	DefaultHeader                map[string]string `json:"defaultHeader,omitempty"`
	UserAgent                    string            `json:"userAgent,omitempty"`
	ClientAuthCertFile           string
	ClientAuthKeyFile            string
	CAFile                       string
	ClientAuthCertString         string
	ClientAuthKeyString          string
	CAString                     string
	Insecure                     bool
	RetriesConfiguration         ClientRetriesConfiguration
	HTTPClient                   *http.Client
	InsecureSkipServerNameVerify bool
	OverwriteProtection          bool
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "https://nsxmanager.your.domain/api/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
