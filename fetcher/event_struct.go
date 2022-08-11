// Copyright 2022 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fetcher

import "time"

type Event struct {
    ID        string    `json:"id"`
    Type      string    `json:"type"`
    Actor     Actor     `json:"actor"`
    Repo      Repo      `json:"repo"`
    Payload   Payload   `json:"payload"`
    Public    bool      `json:"public"`
    CreatedAt time.Time `json:"created_at"`
}
type Actor struct {
    ID           int    `json:"id"`
    Login        string `json:"login"`
    DisplayLogin string `json:"display_login"`
    GravatarID   string `json:"gravatar_id"`
    URL          string `json:"url"`
    AvatarURL    string `json:"avatar_url"`
}
type Repo struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    URL  string `json:"url"`
}
type Author struct {
    Email string `json:"email"`
    Name  string `json:"name"`
}
type Commits struct {
    Sha      string `json:"sha"`
    Author   Author `json:"author"`
    Message  string `json:"message"`
    Distinct bool   `json:"distinct"`
    URL      string `json:"url"`
}
type Payload struct {
    PushID       int64     `json:"push_id"`
    Size         int       `json:"size"`
    DistinctSize int       `json:"distinct_size"`
    Ref          string    `json:"ref"`
    Head         string    `json:"head"`
    Before       string    `json:"before"`
    Commits      []Commits `json:"commits"`
}
