package main

import (
	"expvar"
	"net/http"
	"time"
)

// 1. ecpvar库用于暴露一些公共的变量，返回形式为JSON

/* func main() {
	http.ListenAndServe(":8080", expvar.Handler())
} */

/*使用 http :8080命令进行帧听，返回如下
{
"cmdline": ["/tmp/go-build794785460/b001/exe/use_expvar"],
"memstats": {"Alloc":301856,"TotalAlloc":301856,"Sys":71650312,"Lookups":0,"Mallocs":741,"Frees":13,"HeapAlloc":301856,"HeapSys":66748416,"HeapIdle":65904640,"HeapInuse":843776,"HeapReleased":65904640,"HeapObjects":728,"StackInuse":360448,"StackSys":360448,"MSpanInuse":20128,"MSpanSys":32768,"MCacheInuse":13888,"MCacheSys":16384,"BuckHashSys":3915,"GCSys":3801768,"OtherSys":686613,"NextGC":4473924,"LastGC":0,"PauseTotalNs":0,"PauseNs":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"PauseEnd":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"NumGC":0,"NumForcedGC":0,"GCCPUFraction":0,"EnableGC":true,"DebugGC":false,"BySize":[{"Size":0,"Mallocs":0,"Frees":0},{"Size":8,"Mallocs":16,"Frees":0},{"Size":16,"Mallocs":288,"Frees":0},{"Size":32,"Mallocs":63,"Frees":0},{"Size":48,"Mallocs":105,"Frees":0},{"Size":64,"Mallocs":38,"Frees":0},{"Size":80,"Mallocs":7,"Frees":0},{"Size":96,"Mallocs":18,"Frees":0},{"Size":112,"Mallocs":4,"Frees":0},{"Size":128,"Mallocs":9,"Frees":0},{"Size":144,"Mallocs":2,"Frees":0},{"Size":160,"Mallocs":19,"Frees":0},{"Size":176,"Mallocs":4,"Frees":0},{"Size":192,"Mallocs":0,"Frees":0},{"Size":208,"Mallocs":28,"Frees":0},{"Size":224,"Mallocs":3,"Frees":0},{"Size":240,"Mallocs":0,"Frees":0},{"Size":256,"Mallocs":10,"Frees":0},{"Size":288,"Mallocs":4,"Frees":0},{"Size":320,"Mallocs":3,"Frees":0},{"Size":352,"Mallocs":10,"Frees":0},{"Size":384,"Mallocs":23,"Frees":0},{"Size":416,"Mallocs":7,"Frees":0},{"Size":448,"Mallocs":0,"Frees":0},{"Size":480,"Mallocs":1,"Frees":0},{"Size":512,"Mallocs":0,"Frees":0},{"Size":576,"Mallocs":2,"Frees":0},{"Size":640,"Mallocs":3,"Frees":0},{"Size":704,"Mallocs":2,"Frees":0},{"Size":768,"Mallocs":0,"Frees":0},{"Size":896,"Mallocs":8,"Frees":0},{"Size":1024,"Mallocs":15,"Frees":0},{"Size":1152,"Mallocs":3,"Frees":0},{"Size":1280,"Mallocs":2,"Frees":0},{"Size":1408,"Mallocs":1,"Frees":0},{"Size":1536,"Mallocs":1,"Frees":0},{"Size":1792,"Mallocs":5,"Frees":0},{"Size":2048,"Mallocs":1,"Frees":0},{"Size":2304,"Mallocs":2,"Frees":0},{"Size":2688,"Mallocs":2,"Frees":0},{"Size":3072,"Mallocs":0,"Frees":0},{"Size":3200,"Mallocs":0,"Frees":0},{"Size":3456,"Mallocs":0,"Frees":0},{"Size":4096,"Mallocs":7,"Frees":0},{"Size":4864,"Mallocs":0,"Frees":0},{"Size":5376,"Mallocs":1,"Frees":0},{"Size":6144,"Mallocs":1,"Frees":0},{"Size":6528,"Mallocs":0,"Frees":0},{"Size":6784,"Mallocs":0,"Frees":0},{"Size":6912,"Mallocs":0,"Frees":0},{"Size":8192,"Mallocs":1,"Frees":0},{"Size":9472,"Mallocs":0,"Frees":0},{"Size":9728,"Mallocs":0,"Frees":0},{"Size":10240,"Mallocs":8,"Frees":0},{"Size":10880,"Mallocs":0,"Frees":0},{"Size":12288,"Mallocs":0,"Frees":0},{"Size":13568,"Mallocs":0,"Frees":0},{"Size":14336,"Mallocs":0,"Frees":0},{"Size":16384,"Mallocs":0,"Frees":0},{"Size":18432,"Mallocs":0,"Frees":0},{"Size":19072,"Mallocs":0,"Frees":0}]}
}
*/

/*
这是因为expvar在实现的时候，会自动带上cmdline 和 memstats这两节：
func cmdline() interface{}{
	return os.Args
}

func memstats() interface{} {
	stats := new(runtime.MemStats)
	runtime.ReadMemState(stats)
	return *stats
}

func init() {
        http.HandleFunc("/debug/vars", expvarHandler)
        Publish("cmdline", Func(cmdline))
        Publish("memstats", Func(memstats))
}

*/

/* 2. 可以加上自己想要展示的公共变量
 */

func main() {
	lastAccess := expvar.NewString("lastAccess")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
		lastAccess.Set(time.Now().String())
	})

	http.ListenAndServe(":8080", nil)
}

/* 3. 下面看expvar的源码实现 */
// 从NewString入手
lastAccess := expvar.NewString("lastAccess")

func NewString(name string) *String {
	v := new(String)
	Publish(name, v)
	return v
}

// String的定义
type String struct {
	// atomic.Value此类型的值相当于一个容器，可以被用来“原子地”存储和加载任意类型的值
	s atomic.Value // String
}

// publich的定义
func Publish(name string, v Var) {
	if _, dup := vars.LoaderOrStore(name, v); dup {
		log.Panicln("Reuse of exported var name:", name)
	}
	varKeyMu.Lock()
	defer varKeysMu.Unlock()
	varKeys = append(varKeys, name)
	sort.Strings(varKeys)
}

// 可以看到 Var 是一个接口
// Var is an abstract type for all exported variables.
type Var interface {
	// String returns a valid JSON value for the variable.
	// Types with String methods that do not return valid JSON
	// (such as time.Time) must not be used as a Var.
	String() string
}

// vars 和 varKeys 的定义是这样的
// All published variables.
var (
	vars      sync.Map // map[string]Var
	varKeysMu sync.RWMutex
	varKeys   []string // sorted
)

// 所以可以看到，逻辑就是每次把数值存储到 `vars`，`vars` 是一个map[string]Var类型的map。
// 因为map迭代时是无序的，
// 所以有 `varKeys` 用来排序，这样输出的时候，就可以每次都有序输出

// 来看看如何暴露这些变量
func expvarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "{\n")
	first := true
	Do(func(kv KeyValue) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
	})
	fmt.Fprintf(w, "\n}\n")
}


// 可以看到，很粗暴，先输出 `{\n`，然后输出key value的字符串，最后输出 `\n}\n`，相当于代码拼接JSON
// Do这个函数用来迭代
// Do calls f for each exported variable.
// The global variable map is locked during the iteration,
// but existing entries may be concurrently updated.
func Do(f func(KeyValue)) {
	varKeysMu.RLock()
	defer varKeysMu.RUnlock()
	for _, k := range varKeys {
		val, _ := vars.Load(k)
		f(KeyValue{k, val.(Var)})
	}
}

// 就如我们前面所说，按 varKeys 的顺序来迭代，然后依次执行传入的函数

