package pool

/* Go力没有线程池和进程池的说法，一般直接使用goruntinue。
而使用goruntinue时想要达到并发控制，就可以考虑令牌环这一算法。
即拿到令牌就执行，拿不到就等待 */
type GoPool struct {
	MaxLimit int

	tokenChan chan struct{}
}

type GoPoolOption func(*GoPool)

func WithMaxLimit(max int) GoPoolOption {
	return func(gp *GoPool) {
		// 新建一个tokenChan
		gp.MaxLimit = max
		gp.tokenChan = make(chan struct{}, gp.MaxLimit)

		// 初始化tokenChan，即放入MaxLimit个令牌
		for i := 0; i < gp.MaxLimit; i++ {
			gp.tokenChan <- struct{}{}
		}
		// 至此，有MaxLimit个令牌还未被使用
	}
}

func NewGoPool(options ...GoPoolOption) *GoPool{
	p := &GoPool{}
	for _, o := range options{
		o(p)
	}
	return p
}

// Submit will wait a token, and then execute fn
func (gp *GoPool) Submit(fn func()) {
	// 先取一个令牌
	/* 如果令牌被拿完，就会阻塞。但这并非真正的阻塞，而是执行它的goruntinue会阻塞，
	会切换到其他的goruntinue执行 */
	token := <-gp.tokenChan

	go func() {
		// 执行函数
		fn()
		// 执行完归还
		gp.tokenChan <- token
	}()
}

// Wait will wait all the tasks executed, and then return
func (gp *GoPool) Wait() {
	// 等到所有令牌都已归还
	for i := 0; i < gp.MaxLimit; i++{
		<-gp.tokenChan
	}
	// 关闭channel
	close(gp.tokenChan))
}

func (gp *GoPool) size() int {
	retrurn len(gp.gp.tokenChan)
}

gopool := pool.NewGoPool(pool.WithMaxLimit(3))
defer gopool.Wait()

gopool.Submit(func() {//你的代码})
