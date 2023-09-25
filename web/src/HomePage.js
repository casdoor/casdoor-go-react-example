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

import React from "react";
import { isSilentSigninRequired, SilentSignin } from "casdoor-react-sdk";
import * as Setting from "./Setting";
import LoginPage from "./LoginPage";

class HomePage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      account: undefined,
    };
  }

  componentDidMount() {
    if (Setting.isLoggedIn()) {
      Setting.getUserinfo().then((res) => {
        if (res?.status === "ok") {
          this.setState({
            account: {
              username: res.data.displayName,
              avatar: res.data.avatar,
            },
          });
        } else {
          Setting.showMessage(res?.status);
        }
      });
    }
  }

  logout() {
    Setting.logout();
    Setting.goToLink("/");
  }

  render() {
    if (Setting.isLoggedIn()) {
      if (this.state.account) {
        return (
          <div
            style={{
              marginTop: 200,
              textAlign: "center",
              alignItems: "center",
            }}
          >
            <img
              width={100}
              height={100}
              src={this.state.account.avatar}
              alt={this.state.account.username}
            />
            <br />
            <p>{this.state.account.username}</p>
            <br />
            <button onClick={this.logout}>Logout</button>
          </div>
        );
      } else {
        return <p>Loading...</p>;
      }
    }

    if (isSilentSigninRequired()) {
      return (
        <div
          style={{ marginTop: 200, textAlign: "center", alignItems: "center" }}
        >
          <SilentSignin
            sdk={Setting.CasdoorSdk}
            isLoggedIn={Setting.isLoggedIn}
            handleReceivedSilentSigninSuccessEvent={() => Setting.goToLink("/")}
            handleReceivedSilentSigninFailureEvent={() => {}}
          />
          <p>Logging in...</p>
        </div>
      );
    }

    return (
      <div style={{ marginTop: 200 }}>
        <LoginPage />
      </div>
    );
  }
}

export default HomePage;
