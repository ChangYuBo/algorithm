package main

import (
	"fmt"
	"time"
)

/*把  1 到 16  填到一个 4 * 4 的矩阵里面
[16]byte
使其横竖斜,  加起来都等于34

幻方 穷举回溯法求解
*/
var order = 4
var sum = 34
var resultArr = [17]int{}

var c = 0
var total = 0
var starArr = [17]int{}

func main() {
	start := time.Now()
	//MagicSquares(0)
	fmt.Println("执行次数:", c)
	elapsed := time.Now().Sub(start)
	fmt.Println("总耗时:", elapsed)
	fmt.Println(resultArr, starArr)
}

func MagicSquares(m int) {
	c++
	if m == order*order-1 {
		//最后一个数字必须满足行列以及对角线和都为34
		t1 := sum - (resultArr[12] + resultArr[13] + resultArr[14]) //行
		t2 := sum - (resultArr[3] + resultArr[7] + resultArr[11])   //列
		t3 := sum - (resultArr[0] + resultArr[5] + resultArr[10])   //对角线
		if t1 == t2 && t2 == t3 && starArr[t1] == 0 {
			resultArr[m] = t1
			total++
			fmt.Printf("第%v种:\n", total)
			lastArr := resultArr[:16]
			for k, v := range lastArr {
				if k == 4 || k == 8 || k == 12 || k == 16 {
					fmt.Println()
				}
				fmt.Print(" ", v)
			}
			fmt.Println()
		}
		starArr[resultArr[m-1]] = 0
		return
	}
	//行剪枝
	if m > 0 && m%4 == 0 && resultArr[(m/order-1)*order]+resultArr[(m/order-1)*order+1]+resultArr[(m/order-1)*order+2]+resultArr[(m/order-1)*order+3] != sum {
		starArr[resultArr[m-1]] = 0
		return
	}
	//中心剪枝
	if m == 11 && resultArr[5]+resultArr[6]+resultArr[9]+resultArr[10] != sum {
		starArr[resultArr[m-1]] = 0
		return
	}
	//对角线剪枝和边行剪枝
	if m == 13 && (resultArr[3]+resultArr[6]+resultArr[9]+resultArr[12] != sum || resultArr[4]+resultArr[8]+resultArr[7]+resultArr[11] != sum) {
		starArr[resultArr[m-1]] = 0
		return
	}
	//列剪枝
	if (m == 13 || m == 14) && resultArr[m-13]+resultArr[m-13+4]+resultArr[m-13+8]+resultArr[m-13+12] != sum {
		starArr[resultArr[m-1]] = 0
		return
	}
	//数组长度不够时
	for i := 1; i <= 16; i++ {
		if starArr[i] == 0 {
			resultArr[m] = i
			starArr[i] = 1
			MagicSquares(m + 1)
		}
	}
	if m > 0 {
		starArr[resultArr[m-1]] = 0
	}
}
