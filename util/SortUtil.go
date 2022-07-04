/*
 * @Autor: 郭彬
 * @Description:
 * @Date: 2022-06-17 10:03:42
 * @LastEditTime: 2022-06-17 18:29:17
 * @FilePath: \Test\util\SortUtil.go
 */
package util

// 冒泡排序 O(n2)
// 思想： 相连的元素不断比较，就像冒泡一样，大的元素逐步上升到尾部
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 选择排序 O(n2)
// i后面的，元素依次和i比较，如果 arr[j] < arr[i] 则交换位置
func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// 直接插入排序
// 思想: 每一步将一个待排序的数据插入到前面已经排好序的有序序列中，直到插完所有元素为止。
func InSertSort(arr []int) {
	n := len(arr)
	// 进行 N-1 轮迭代
	for i := 1; i < n; i++ {
		deal := arr[i] // 待排序的数
		j := i - 1     // 待排序的数左边的第一个数的位置
		for ; j >= 0 && deal < arr[j]; j-- {
			arr[j+1] = arr[j] // 某数后移，给待排序留空位
		}
		arr[j+1] = deal // 结束了，待排序的数插入空位
	}
}

/*
 希尔排序##
（1）希尔排序是 插入排序的一种， 这个排序方法又称为缩小增量排序，是直接插入排序的改进版本，实际上是一种分组插入排序法。
     该方法的基本思想是：设待排序元素序列有n个元素，首先取一个整数increment（小于n）作为间隔将全部元素分为increment个子序列，所有距离为increment的元素放在同一个子序列中，在每一个子序列中分别实行直接插入排序。
	 然后缩小间隔increment，重复上述子序列划分和排序工作。
	 直到最后取increment=1，将所有元素放在同一个子序列中排序为止。
（2）由于开始时，increment的取值较大，每个子序列中的元素较少，排序速度较快，到排序后期increment取值逐渐变小，子序列中元素个数逐渐增多，但由于前面工作的基础，大多数元素已经基本有序，所以排序速度仍然很快。
*/

func ShellSort(arr []int) {
	len := len(arr)
	gap := 1
	for gap < len/3 {
		gap = 3*gap + 1
	}
	for gap >= 1 {
		for i := gap; i < len; i++ {
			e := arr[i]
			var j int
			for j = i; j >= gap && arr[j-gap] > e; j = j - gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = e
		}
		gap /= 3
	}
}

// 快速排序
// 快速排序使用分治的思想，通过一趟排序将待排序列分割成两部分，其中一部分记录的关键字均比另一部分记录的关键字小。
// 之后分别对这两部分记录继续进行排序，以达到整个序列有序的目的。
// 需要用到递归
func QuikSort(arr []int, begin, end int) {
	if begin > end {
		return
	}
	tmp := arr[begin]
	i := begin
	j := end

	for i != j {
		for arr[j] >= tmp && j > i {
			j--
		}
		for arr[i] <= tmp && j > i {
			i++
		}
		if j > i {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[begin], arr[i] = arr[i], arr[begin]
	QuikSort(arr, begin, i-1)
	QuikSort(arr, i+1, end)
}
