package main

import (
	"fmt"
	"go-training/example18-write-testing-and-doc/car" // 导入 car 包
)

func main() {
	// 1. 创建 Car 实例（成功）
	myCar, err := car.New("奥迪A4", 32.5)
	if err != nil {
		fmt.Println("创建失败：", err)
		return
	}
	fmt.Println("初始名称：", myCar.Name) // 输出：初始名称：奥迪A4

	// 2. 修改名称（有效）
	newName := myCar.SetName("奥迪A6")
	fmt.Println("修改后名称：", newName) // 输出：修改后名称：奥迪A6

	// 3. 尝试设置空名称（无效）
	emptyName := myCar.SetName("")
	fmt.Println("空名称设置后：", emptyName) // 输出：空名称设置后：奥迪A6

	// 4. 创建失败场景（名称为空）
	invalidCar, err := car.New("", 20.0)
	if err != nil {
		fmt.Println("创建无效汽车失败：", err) // 输出：创建无效汽车失败：missing name
	}
	fmt.Println(invalidCar) // 输出：<nil>
}
