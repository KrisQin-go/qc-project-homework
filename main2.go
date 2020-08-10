package main

import (
	"flag"
	"fmt"
)

func main() {
	var cin string
	const ts = "1"
	const sc = "2"
	const help = `
程序名称：main
描述:输入1或2，输入1时输出唐诗，输入2时输出宋词
`
	flag.Usage = func() {
		fmt.Printf(help)
	}
	flag.Parse()
	fmt.Println("输入1或2:")
	if _, err := fmt.Scanf("%s", &cin); err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	if ts == cin {
		fmt.Printf("孤帆远影碧空尽，唯见长江天际流。")
		return
	}
	if sc == cin {
		fmt.Printf("明月几时有？把酒问青天。")
		return
	}
}
