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

import "./App.css";

import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { AuthCallback } from "casdoor-react-sdk";
import * as Setting from "./Setting";
import HomePage from "./HomePage";

class App extends React.Component {
  authCallback = (
    <AuthCallback
      sdk={Setting.CasdoorSdk}
      serverUrl={Setting.ServerUrl}
      saveTokenFromResponse={(res) => {
        Setting.setToken(res.data);
        Setting.goToLink("/");
      }}
      isGetTokenSuccessful={(res) => res.status === "ok"}
    />
  );

  render() {
    return (
      <BrowserRouter>
        <Routes>
          <Route path="/callback" element={this.authCallback} />
          <Route path="/" element={<HomePage />} />
        </Routes>
      </BrowserRouter>
    );
  }
}

export default App;
