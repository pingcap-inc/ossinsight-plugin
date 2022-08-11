# Copyright 2022 PingCAP, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: redis config

redis:
	docker run -itd --name redis-test -p 6379:6379 redis

clean:
	docker stop redis-test
	docker rm redis-test

config:
	go install github.com/Icemap/yaml2go-cli@latest
	yaml2go-cli -p config -s Config -i config/config.yaml -o config/config_struct.go