// Copyright 2023 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

func signinHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	token, err := casdoorsdk.GetOAuthToken(code, state)
	if err != nil {
		fmt.Println("GetOAuthToken() error", err)
		http.Error(w, "GetOAuthToken() error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
		"data":   token.AccessToken,
	})
}

func userinfoHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "authHeader is empty", http.StatusUnauthorized)
		return
	}

	token := strings.Split(authHeader, "Bearer ")
	if len(token) != 2 {
		http.Error(w, "token is not valid Bearer token", http.StatusUnauthorized)
		return
	}

	claims, err := casdoorsdk.ParseJwtToken(token[1])
	if err != nil {
		http.Error(w, "ParseJwtToken() error", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
		"data":   claims.User,
	})
}
