package main

import "cliNew/myCli"

func main() {

	// 第一个初始化操作 运行并输出一些默认信息   ./hello --help
	//test.Test01()
	// 第二个传入一些参数使用  ./hello 参数一  参数二
	//test.Test02()

	//test.Test03()

	//test.Test04()

	// 自己的一些cli工具
	myCli.MyCli()

}

/**
form-data  参数使用
	./cliTest param   f -f a.txt


*/
