package tcp

/**
 * TCP服务器实现
 *
 * 本文件提供了一个通用的TCP服务器框架，包括：
 * 1. 监听TCP端口并接收客户端连接
 * 2. 使用goroutine处理并发连接
 * 3. 支持优雅关闭服务器
 * 4. 处理各种错误情况
 */

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/hdt3213/godis/interface/tcp"
	"github.com/hdt3213/godis/lib/logger"
)

// Config 存储TCP服务器配置属性
type Config struct {
	Address    string        `yaml:"address"`     // 服务器监听地址，格式为"IP:Port"
	MaxConnect uint32        `yaml:"max-connect"` // 最大并发连接数
	Timeout    time.Duration `yaml:"timeout"`     // 连接超时时间
}

// ClientCounter 记录当前Godis服务器的客户端连接数量
var ClientCounter int32

// ListenAndServeWithSignal 绑定端口并处理请求，阻塞直到接收到停止信号
// 参数:
// - cfg: 服务器配置
// - handler: 处理客户端连接的处理器
// 返回:
// - error: 启动过程中遇到的错误
func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error {
	// 创建关闭通知通道
	closeChan := make(chan struct{})

	// 设置信号监听，捕获操作系统发送的终止信号
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	// 启动goroutine监听系统信号
	go func() {
		sig := <-sigCh
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			// 收到终止信号后，发送关闭通知
			closeChan <- struct{}{}
		}
	}()

	// 开始监听TCP端口
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}

	//cfg.Address = listener.Addr().String()
	logger.Info(fmt.Sprintf("bind: %s, start listening...", cfg.Address))

	// 调用主服务函数
	ListenAndServe(listener, handler, closeChan)
	return nil
}

// ListenAndServe 绑定端口并处理请求，阻塞直到服务关闭
// 参数:
// - listener: 网络监听器
// - handler: 处理客户端连接的处理器
// - closeChan: 接收关闭信号的通道
func ListenAndServe(listener net.Listener, handler tcp.Handler, closeChan <-chan struct{}) {
	// 创建错误通道，用于监听Accept过程中的错误
	errCh := make(chan error, 1)
	defer close(errCh)

	// 启动goroutine监听关闭信号或错误
	go func() {
		select {
		case <-closeChan:
			// 收到关闭信号
			logger.Info("get exit signal")
		case er := <-errCh:
			// 接收连接过程中出现错误
			logger.Info(fmt.Sprintf("accept error: %s", er.Error()))
		}
		logger.Info("shutting down...")
		// 关闭监听器，会导致Accept()立即返回错误
		_ = listener.Close()
		// 关闭所有现有连接
		_ = handler.Close()
	}()

	// 创建上下文，并使用等待组同步所有客户端goroutine
	ctx := context.Background()
	var waitDone sync.WaitGroup

	// 主循环：接受新的客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			// 参考net/http/serve.go#Serve()的错误处理策略
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				// 如果是临时错误(如资源暂时不可用)，稍后重试
				logger.Infof("accept occurs temporary error: %v, retry in 5ms", err)
				time.Sleep(5 * time.Millisecond)
				continue
			}
			// 其他错误(如listener关闭)，发送到错误通道并退出循环
			errCh <- err
			break
		}

		// 新连接建立成功，增加计数并启动goroutine处理
		logger.Info("accept link")
		ClientCounter++
		waitDone.Add(1)

		// 为每个客户端启动一个独立的goroutine
		go func() {
			defer func() {
				// 客户端处理完毕，更新计数
				waitDone.Done()
				atomic.AddInt32(&ClientCounter, -1)
			}()
			// 调用实际的处理器处理连接
			handler.Handle(ctx, conn)
		}()
	}

	// 等待所有客户端连接处理完毕
	// 这确保了服务器在关闭前完成所有正在进行的客户端交互
	waitDone.Wait()
}
