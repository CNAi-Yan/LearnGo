// package main 定义了包名
package main

// 依赖
import (
	"fmt"
)

func main() {
	/* 这是我的第一个简单的程序 */ //注释
	age := 28
	age++
	fmt.Println("age++ 的值是:",age, "\n")
	fmt.Println("age++ 的值是否等于29:",age == 29, "\n")
}
/*
当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Println，
那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），
这被称为导出（像面向对象语言中的 public）；
标识符如果以小写字母开头，则对包外是不可见的，
但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）
*/

/*
需要注意的是 { 不能单独放在一行，所以以下代码在运行时会产生错误：
# command-line-arguments
.\main.go:7:1: syntax error: unexpected semicolon or newline before {
*/

/*
同一个文件夹下的文件只能有一个包名，否则编译报错。
main.go:6:2: "./myMath" is relative, but relative import paths are not supported in module mode
*/
