package main

// /* 先看如何使用flag */
// import (
// 	"flag"
// 	"log"
// )

// var (
// 	useWorker = flag.Bool("userWorker", true, "Hello, Are You There?")
// )

// func main() {
// 	flag.Parse()
// 	log.Printf("useWorker: %t", *useWorker)
// }

// /* 可以发现：
// 1.seWorker的类型是 *bool，而且是flag.Bool返回的
// 2.必须要执行 flag.Parse() 才能解析命令行 */

// /* 输出
// go run flag.go
//  */

//  /* Bool和Parse的源码 */
//  func Bool(name string, value bool, usage string) *bool {
// 	return CommandLine.Bool(name, value, usage)
// }

// func (f *FlagSet) Bool(name string, value bool, usage string) *bool {
// 	p := new(bool)
// 	f.BoolVar(p, name, value, usage)
// 	return p
// }

// func Parse() {
// 	// Ignore errors; CommandLine is set for ExitOnError.
// 	CommandLine.Parse(os.Args[1:])
// }

// func (f *FlagSet) Parse(arguments []string) error {
// 	f.parsed = true
// 	f.args = arguments
// 	for {
// 		seen, err := f.parseOne()
// ...

// func (f *FlagSet) parseOne() (bool, error) {
// 	if len(f.args) == 0 {
// 		return false, nil
// 	}
// ...
// 	if fv, ok := flag.Value.(boolFlag); ok && fv.IsBoolFlag() { // special case: doesn't need an arg
// 		if hasValue {
// 			if err := fv.Set(value); err != nil {
// 				return false, f.failf("invalid boolean value %q for -%s: %v", value, name, err)
// 			}
// 		} else {
// 			if err := fv.Set("true"); err != nil {
// 				return false, f.failf("invalid boolean flag %s: %v", name, err)
// 			}
// 		}
//     ...

// 	/* 可以理解为，flag.Bool新建一个bool类型的变量，然后赋给它默认值，再返回这个bool类型变量的指针。
// 	 在Parse函数里，对这个指针所指向的值进行更新。*/

/* 下面自己实现一个简单的flag，若要运行，别忘记注销上面所有的代码 */

import (
	"log"
	"os"
	"strings"
)

type MyFlagger interface {
	Set(v interface{})
}

type MyFlag struct {
	mapper map[string]MyFlagger
}

var myFlags = MyFlag{mapper: make(map[string]MyFlagger)}

type boolFlag struct {
	p *bool
}

func (b *boolFlag) Set(v interface{}) {
	*(b.p) = v.(bool)
}

func (m *MyFlag) Bool(name string, defaultValue bool) *bool {
	p := new(bool)
	*p = defaultValue
	m.mapper[name] = &boolFlag{p}
	return p
}

func (m *MyFlag) Parse() {
	if len(os.Args) == 1 {
		return
	}

	arg := os.Args[1]
	if !strings.HasPrefix(arg, "--") {
		log.Panicf("bad usage: ./test --blabla")
	}

	if len(arg) < 3 {
		log.Panicf("bad usage: ./test --blabla")
	}

	realArg := arg[2:]
	flag, ok := m.mapper[realArg]
	if !ok {
		log.Panicf("%s not found", realArg)
	}

	flag.Set(true)
}

func main() {
	useWorker := myFlags.Bool("useWorker", false)
	log.Printf("before parse: useWorker: %t", *useWorker)
	myFlags.Parse()
	log.Printf("after parse: useWorker: %t", *useWorker)
}
