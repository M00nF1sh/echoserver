# Copyright 2017[<0;55;12M The Kubernetes Authors. All rights reserved.
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

FROM golang:alpine AS build-env
WORKDIR /go/src/github.com/M00nF1sh/echoserver
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=build-env /go/src/github.com/M00nF1sh/echoserver/server /app/
ENTRYPOINT /app/server