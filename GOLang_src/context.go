package main

import (
	"context"
	"log"
	"sync"
	"time"
)

/* Golang中使用context作为goroutine之间的控制器, 即想要在一个goroutinue中使用另一个goroutinue */
func UseContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // 检测context的状态。ctx.Done()返回一个channel，如果这个channel已经被执行完毕，就执行下面的语句,
			// 同时，这也是其缺点，即每次使用钱都要手动检测一下channel是否已经关闭
			log.Printf("context is done with error %s", ctx.Err())
			return
		default: // 否则走默认分支
			log.Printf("nothing just loop...")
			time.Sleep(time.Second * time.Duration(1))
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go UseContext(ctx)

	time.Sleep(time.Second * time.Duration(1))
	cancel()
	time.Sleep(time.Second * time.Duration(2))
}

/* 执行结果
2020/11/07 10:19:14 nothing just loop...
2020/11/07 10:19:15 context is done with error context canceled
*/

/* 如此，便可以在main函数里告知UserContext所在的goroutinue，主函数已经想要退出。
本质很简单，就是王main力传递一个变量，UserContext不断地取检查变量而已，
只不过在Go里，用channel来实现 */

/* 细看context的源码 */
type Context interface {
	Deadline() (Deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

//  实现这个接口，例如用context.Background()
func Background() Context {
	return background
}

// background是emptyCtx的一个实例
var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

// WithCancel的实现
// WithCancel returns a copy of parent with a new Done channel. The returned
// context's Done channel is closed when the returned cancel function is called
// or when the parent context's Done channel is closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	c := newCancelCtx(parent)
	propagateCancel(parent, &c)
	return &c, func() { c.cancel(true, Canceled) }
}

// A cancelCtx can be canceled. When canceled, it also cancels any children
// that implement canceler.
type cancelCtx struct {
	Context

	mu       sync.Mutex            // protects following fields
	done     chan struct{}         // created lazily, closed by first cancel call
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first cancel call
}

// 以后可以研究Context中的传值和取值是如何实现的
