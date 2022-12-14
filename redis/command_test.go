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

package redis

import (
	"context"
	"strconv"
	"testing"
	"time"
)

func TestExistsAndSet(t *testing.T) {
	key := strconv.Itoa(time.Now().Nanosecond())
	exist, err := EventIDExists(key)
	if err != nil {
		t.Error(err)
		return
	}

	if exist {
		t.Errorf("%s should not exists", key)
		return
	}

	exist, err = EventIDExists(key)
	if err != nil {
		t.Error(err)
		return
	}

	if !exist {
		t.Errorf("%s should exists", key)
		return
	}
}

func TestHyperLogLogExists(t *testing.T) {
	initClient()
	client.Del(context.Background(), distinctPrefix+"test")
	exists, err := HyperLogLogExists("test", "a")
	if err != nil {
		t.Errorf("add error")
	}
	if exists {
		t.Errorf("not exist")
	}

	exists, err = HyperLogLogExists("test", "b")
	if err != nil {
		t.Errorf("add error")
	}
	if exists {
		t.Errorf("not exist")
	}

	exists, err = HyperLogLogExists("test", "a")
	if err != nil {
		t.Errorf("add error")
	}
	if !exists {
		t.Errorf("exist")
	}
}
