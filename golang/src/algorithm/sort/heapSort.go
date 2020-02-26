/*
 堆排序，数组下标从1开始表示一根完全二叉树
*/
package sort

func adjust(arr []int, low int, high int) {
	i := low
	j := 2 * i
	temp := arr[i]

	for j <= high {
		if j < high && arr[j] < arr[j+1] {
			j++
		}

		if temp < arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
			i = j
			j = 2 * i
		} else {
			break
		}
	}

	arr[i] = temp

}

/*
	n为数组的大小减去一
	eg:
	var arr = []int{0, 10, 2, 16, 8, 9}
 	heapSort(arr, len(arr)-1)
 	fmt.Println(arr) //output:[0 2 8 9 10 16]
*/
func HeapSort(arr []int, n int) {
	for i := n / 2; i >= 1; i-- {
		adjust(arr, i, n)
	}

	for i := n; i > 1; i-- {
		arr[1], arr[i] = arr[i], arr[1]
		adjust(arr, 1, i-1)
	}
}
