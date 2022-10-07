package utils

//比如  我在utils包下面  新建一个函数  交getresult   您可以通过包来组织您的diama   名字您取有意义的 一眼就能看出来做什么的就可以
// 一个pack下面可以由多个文件   每个go文件下面可以由多个函数   这里我就使用这一个go文件   里面就写一个函数

import "github.com/google/keytransparency/core/crypto/vrf/p256"

// 这里示意一下  函数名字要大写 不然不对外公开  别人无法调用  在golang中     开头大写代表公开，否则不公开  外部不能直接调用   函数名字  您看您要实现什么
//功能把  我这里就这样写了
// 我们需要两个入参
func GetResult(d1, d2 string) (r1, r2 [32]byte) {
	//生成个key
	k, pk := p256.GenerateKey()
	//输入两个参数
	//这里做下判断
	if len(d1) == 0 {
		d1 = "This is a test" //给他个默认的   不给也可以  直接提示错误 结束函数， 看您需求把
	}
	//两种判空都可以
	if d2 == "" {
		d2 = "This is not a test"
	}

	//下面这两个咱们在函数中就不需要了   因为外界的入参是什么  咱们不用管   写死的 或者从命令行输入的  最终都是调用者来决定
	// //获取命令行输入的参数   （默认有一个参数为程序的名字，所以这里取切片从1开始取）
	// argCount := len(os.Args[1:])
	// if argCount > 0 {
	// 	d1 = os.Args[1]
	// }
	// if argCount > 1 {
	// 	d2 = os.Args[2]
	// }
	// //接下来就是如果没有输入参数  ，那么使用上面您声明的默认的这两个参数

	m1 := []byte(d1)
	m2 := []byte(d2)

	//接下来就是调用第三方库  做一些处理 ，以及输出结果了
	index1, proof1 := k.Evaluate(m1)
	index2, proof2 := k.Evaluate(m2)

	// fmt.Printf("== Creation of proofs ===\n")
	// fmt.Printf("Data: [%s] Index: %x Proof: %x\n", m1, index1, proof1)
	// fmt.Printf("Data: [%s] Index: %x Proof: %x\n", m2, index2, proof2)

	//fmt.Printf("\n== Verfication of proofs ===\n")
	newindex1, _ := pk.ProofToHash(m1, proof1)
	r1 = newindex1
	//fmt.Printf("Result 1: %x\n", newindex1)
	if index1 == newindex1 {
		//fmt.Printf("Proven\n")
		//不给赋值了  后续您可以在这里写您的处理方式
	}

	newindex2, _ := pk.ProofToHash(m2, proof2)
	//fmt.Printf("Result 2: %x\n", newindex2)
	r2 = newindex2
	if index2 == newindex2 {
		//fmt.Printf("Proven\n")
	}
	return r1, r2
}
