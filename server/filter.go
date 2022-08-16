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

package main

import (
    "encoding/json"
    "strings"
)

// FilterMessageToMap using filter to compare message
func FilterMessageToMap(msg []byte, filter []string) (map[string]interface{}, error) {
    message := make(map[string]interface{})
    err := json.Unmarshal(msg, &message)
    if err != nil {
        return nil, err
    }

    result := make(map[string]interface{})
    for _, filterItem := range filter {
        currentNode := message
        var currentValue interface{}
        for _, level := range strings.Split(filterItem, ".") {
            if value, exist := currentNode[level]; exist {
                currentValue = value
                if node, ok := value.(map[string]interface{}); ok {
                    currentNode = node
                }
            }
        }

        result[filterItem] = currentValue
    }

    return result, nil
}
