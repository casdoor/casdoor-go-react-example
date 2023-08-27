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
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"net/http"
	"strings"
)

func redirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	url := casdoorsdk.GetSigninUrl("http://localhost:3000/callback")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
		"data":   url,
	})
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	fmt.Println("get code and state:", code, state)
	token, err := casdoorsdk.GetOAuthToken(code, state)

	if err != nil {
		fmt.Println("get token error", err)
		http.Error(w, "Error get token", http.StatusInternalServerError)
		return
	}

	casdoorsdk.ParseJwtToken(token.AccessToken)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
		"data":   token.AccessToken,
	})
	fmt.Println("finish")
}

func userinfoHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token := strings.Split(authHeader, "Bearer ")
	if len(token) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	chaims, err := casdoorsdk.ParseJwtToken(token[1])
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
		"data":   chaims.User,
	})
}
