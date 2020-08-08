# suffle_the_array

## 題目解讀：

### 題目來源:

[shuffle-the-array](https://leetcode.com/problems/shuffle-the-array/)

### 原文:

Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].

Return the array in the form [x1,y1,x2,y2,...,xn,yn].

### 解讀：

給定一個正整數陣列nums 一個正整數n

符合以下條件

nums = [x1, x2,..., xn,y1,y2,...yn]

撰寫一個演算法

產生一個正整數陣列results

符合以下條件

results = [x1, y1, x2, y2, ....,xn,yn]

## 初步解法:
### 初步觀察:

給定一個正整數陣列nums

則這個n = nums.length/2

然後 把nums 分成 前半x 跟後半y

兩個array 在同步loop放置到新的array results

### 初步設計:

given an array nums

set n = half of the length nums

create an array results with length of 2n

loop for index 0 to n-1:

results[2*i] = nums[i]
results[2*i+1] = nums[n+i]

return results

## 遇到的困難

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間

## 測資的撰寫

```golang
package shuffle_array

import (
	"reflect"
	"testing"
)

func Test_shuffle(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{2, 5, 1, 3, 4, 7},
				n:    3,
			},
			want: []int{2, 3, 5, 4, 1, 7},
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 2, 3, 4, 4, 3, 2, 1},
				n:    4,
			},
			want: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			name: "Example3",
			args: args{
				nums: []int{1, 1, 2, 2},
				n:    2,
			},
			want: []int{1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shuffle(tt.args.nums, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
```
## my self record
[leetcode 30 golang day2](https://hackmd.io/g84oEi-LTRCD1zebDg_BUw?view)
## 參考文章
[golang test](https://ithelp.ithome.com.tw/articles/10204692)