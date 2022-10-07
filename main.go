package main

import (
	"fmt"

	"github.com/fengzhongda/vrf/utils"
)

func main() {

	// //生成个key
	// k, pk := p256.GenerateKey()
	// //输入两个参数
	// d1 := "This is a test"
	// d2 := "This is not a test"

	// //获取命令行输入的参数   （默认有一个参数为程序的名字，所以这里取切片从1开始取）
	// argCount := len(os.Args[1:])
	// if argCount > 0 {
	// 	d1 = os.Args[1]
	// }
	// if argCount > 1 {
	// 	d2 = os.Args[2]
	// }
	// //接下来就是如果没有输入参数  ，那么使用上面您声明的默认的这两个参数

	// m1 := []byte(d1)
	// m2 := []byte(d2)

	// //接下来就是调用第三方库  做一些处理 ，以及输出结果了
	// index1, proof1 := k.Evaluate(m1)
	// index2, proof2 := k.Evaluate(m2)

	// fmt.Printf("== Creation of proofs ===\n")
	// fmt.Printf("Data: [%s] Index: %x Proof: %x\n", m1, index1, proof1)
	// fmt.Printf("Data: [%s] Index: %x Proof: %x\n", m2, index2, proof2)

	// fmt.Printf("\n== Verfication of proofs ===\n")
	// newindex1, _ := pk.ProofToHash(m1, proof1)
	// fmt.Printf("Result 1: %x\n", newindex1)
	// if index1 == newindex1 {
	// 	fmt.Printf("Proven\n")
	// }

	// newindex2, _ := pk.ProofToHash(m2, proof2)
	// fmt.Printf("Result 2: %x\n", newindex2)
	// if index2 == newindex2 {
	// 	fmt.Printf("Proven\n")
	// }

	//咱们要把自己的代码弄到git上面  供别人调用   首先要把主要逻辑提取出来   封装成函数

	d1 := "This is a test"
	d2 := "This is not a test"
	// r1, r2 := GetResult(d1, d2)//函数和main在同一包下  直接调用就行    在不同包中   要写包名
	r1, r2 := utils.GetResult(d1, d2)

	fmt.Printf("Result 1: %x\n", r1)
	fmt.Printf("Result 2: %x\n", r2)
}
