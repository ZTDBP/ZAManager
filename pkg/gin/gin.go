// Copyright 2022-present The ZTDBP Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gin

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/DeanThompson/ginpprof"
	"github.com/ZTDBP/ZAManager/pkg/confer"
	"github.com/gin-gonic/gin"
)

type OnShutdownF struct {
	f       func(cancel context.CancelFunc)
	timeout time.Duration
}

var (
	onShutdown []OnShutdownF
)

func RegisterOnShutdown(f func(cancel context.CancelFunc), timeout time.Duration) {
	onShutdown = append(onShutdown, OnShutdownF{
		f:       f,
		timeout: timeout,
	})
}

func NewGin() *gin.Engine {
	if confer.ConfigEnvIsDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	if confer.ConfigEnvIsDev() {
		ginpprof.Wrap(r)
		r.Use(gin.Logger())
	}
	return r
}

func ListenHTTP(httpPort string, r http.Handler, timeout int, f ...func()) {
	srv := &http.Server{
		Addr:    httpPort,
		Handler: r,
	}
	// 监听端口
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 注册关闭使用函数
	for _, v := range f {
		srv.RegisterOnShutdown(v)
	}
	// 监听信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	// 执行on shutdown 函数 - 同步
	for _, v := range onShutdown {
		var wg sync.WaitGroup
		wg.Add(1)
		ctx, cancel := context.WithCancel(context.TODO())
		go v.f(cancel)
		select {
		case <-time.After(v.timeout):
			log.Println("on shutdown timeout:", f)
			wg.Done()
		case <-ctx.Done():
			wg.Done()
		}
		wg.Wait()
	}
	// 执行shutdown
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Panic("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
