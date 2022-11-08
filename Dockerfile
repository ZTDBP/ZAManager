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

FROM golang:1.17.8-alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.oneitfarm.com,https://goproxy.cn,direct

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 go build -o ZAManager .


FROM alpine AS final
WORKDIR /app
COPY --from=builder /build/ZAManager /app/
COPY --from=builder /build/config.yaml /app/
COPY --from=builder /build/db /app/db

ENTRYPOINT ["/app/ZAManager"]
