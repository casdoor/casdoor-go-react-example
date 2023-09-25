// Copyright 2022 The Casdoor Authors. All Rights Reserved.
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

import Sdk from "casdoor-js-sdk";

export const ServerUrl = "http://localhost:8080";

const sdkConfig = {
  serverUrl: "https://door.casdoor.com",
  clientId: "b800a86702dd4d29ec4d",
  organizationName: "casbin",
  appName: "app-example",
  redirectPath: "/callback",
};

export const CasdoorSdk = new Sdk(sdkConfig);

export const isLoggedIn = () => {
  const token = localStorage.getItem("token");
  return token !== null && token.length > 0;
};

export const setToken = (token) => {
  localStorage.setItem("token", token);
};

export const goToLink = (link) => {
  window.location.href = link;
};

export function getSigninUrl() {
  return CasdoorSdk.getSigninUrl();
}

export const getRedirectUrl = () => {
  return fetch(`${ServerUrl}/api/redirect-url`, {
    method: "GET",
  }).then((res) => res.json());
};

export const getUserinfo = () => {
  return fetch(`${ServerUrl}/api/userinfo`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  }).then((res) => res.json());
};

export const logout = () => {
  localStorage.removeItem("token");
};

export const showMessage = (message) => {
  alert(message);
};
