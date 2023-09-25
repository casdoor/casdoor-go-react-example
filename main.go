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
	"fmt"
	"net/http"

	"casdoor-go-react-sdk-example/config"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func initAuthConfig() {
	casdoorsdk.InitConfig(
		config.GlobalConfig.Server.Endpoint,
		config.GlobalConfig.Server.ClientID,
		config.GlobalConfig.Server.ClientSecret,
		config.GlobalConfig.Certificate,
		config.GlobalConfig.Server.Organization,
		config.GlobalConfig.Server.Application,
	)
}

func main() {
	err := config.LoadConfig("config/app.yaml")
	if err != nil {
		panic(err)
	}

	initAuthConfig()

	router := mux.NewRouter()
	router.HandleFunc("/api/signin", signinHandler)
	router.HandleFunc("/api/userinfo", userinfoHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	fmt.Println("Server listening at: http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
