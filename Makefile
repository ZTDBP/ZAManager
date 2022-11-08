# Copyright 2022-present The ZTDBP Authors.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#     http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

GOPROXY=https://goproxy.oneitfarm.com,https://goproxy.cn,direct

PROG=bin/ZAManager


SRCS=.

# 安装目录
INSTALL_PREFIX=/usr/local/ZAManager

# 配置安装的目录
CONF_INSTALL_PREFIX=/usr/local/ZAManager

# git commit hash
COMMIT_HASH=$(shell git rev-parse --short HEAD || echo "GitNotFound")

# 编译日期
BUILD_DATE=$(shell date '+%Y-%m-%d %H:%M:%S')

# 编译条件
CFLAGS = -ldflags "-s -w -X \"main.BuildVersion=${COMMIT_HASH}\" -X \"main.BuildDate=$(BUILD_DATE)\""

all:
	if [ ! -d "./bin/" ]; then \
	mkdir bin; \
	fi
	GOPROXY=$(GOPROXY) CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  $(CFLAGS) -o $(PROG) $(SRCS)

# 编译race版本
race:
	if [ ! -d "./bin/" ]; then \
    	mkdir bin; \
    	fi
	go build $(CFLAGS) -race -o $(PROG) $(SRCS)

# release 版本
RELEASE_DATE = $(shell date '+%Y%m%d%H%M%S')
RELEASE_VERSION = $(shell git rev-parse --short HEAD || echo "GitNotFound")
RELEASE_DIR=release_bin
RELEASE_BIN_NAME=ZAManager
release:
	if [ ! -d "./$(RELEASE_DIR)/$(RELEASE_DATE)_$(RELEASE_VERSION)" ]; then \
	mkdir ./$(RELEASE_DIR)/$(RELEASE_DATE)_$(RELEASE_VERSION); \
	fi
	go build  $(CFLAGS) -o $(RELEASE_DIR)/$(RELEASE_DATE)_$(RELEASE_VERSION)/$(RELEASE_BIN_NAME)_linux_amd64 $(SRCS)

install:
	cp $(PROG) $(INSTALL_PREFIX)/bin

	if [ ! -d "${CONF_INSTALL_PREFIX}" ]; then \
	mkdir $(CONF_INSTALL_PREFIX); \
	fi

	cp -R config/* $(CONF_INSTALL_PREFIX)

clean:
	rm -rf ./bin

	rm -rf $(INSTALL_PREFIX)/bin/ZAManager

	rm -rf $(CONF_INSTALL_PREFIX)

run:
	go run main.go

run_race:
	go run --race main.go
