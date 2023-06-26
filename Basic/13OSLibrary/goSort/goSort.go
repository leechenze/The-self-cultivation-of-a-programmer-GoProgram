package goSort

import (
	"fmt"
	"sort"
)

func GoSort() {
	println("========================OS Library========================")
	println()

	/** 基础 Int类型 排序 */
	println()
	basicSort()

	/** 自定义排序 */
	println()
	customSort()

	/** float64排序 */
	println()
	float64Sort()

	/** string排序 */
	println()
	stringSort()

	/** 复杂结构 [][]int 二维数组 */
	println()
	twoDimensionalSort()

	/** 字典排序 */
	println()
	mapSort()

	/** 结构体排序 */
	println()
	structSort()

	println()
	println("========================OS Library========================")

}

/** 结构体排序 */
type People struct {
	Name string
	Age  int
}

type StructSlice []People

func (ss StructSlice) Len() int           { return len(ss) }
func (ss StructSlice) Swap(i, j int)      { ss[i], ss[j] = ss[j], ss[i] }
func (ss StructSlice) Less(i, j int) bool { return ss[i].Age < ss[j].Age }

func structSort() {
	ss := StructSlice{
		{Name: "n2", Age: 11},
		{Name: "n1", Age: 12},
		{Name: "n3", Age: 10},
	}
	fmt.Printf("ss: %v\n", ss)
	sort.Sort(ss)
	fmt.Printf("ss: %v\n", ss)
}

/** 字典排序 */
type MapSlice []map[string]float64

func (ms MapSlice) Len() int           { return len(ms) }
func (ms MapSlice) Swap(i, j int)      { ms[i], ms[j] = ms[j], ms[i] }
func (ms MapSlice) Less(i, j int) bool { return ms[i]["a"] < ms[j]["a"] }

func mapSort() {
	ms := MapSlice{
		{"a": 4, "b": 12},
		{"a": 3, "b": 11},
		{"a": 5, "b": 10},
	}
	fmt.Printf("ms: %v\n", ms)
	sort.Sort(ms)
	fmt.Printf("ms: %v\n", ms)
}

/** 复杂结构 [][]int 二维数组 */
type TwoDimensionalInt [][]int

func (tdi TwoDimensionalInt) Len() int           { return len(tdi) }
func (tdi TwoDimensionalInt) Swap(i, j int)      { tdi[i], tdi[j] = tdi[j], tdi[i] }
func (tdi TwoDimensionalInt) Less(i, j int) bool { return tdi[i][0] < tdi[j][0] }

func twoDimensionalSort() {
	twoDimensionalInt := TwoDimensionalInt{
		{1, 4},
		{9, 3},
		{7, 5},
	}
	fmt.Printf("twoDimensionalInt: %v\n", twoDimensionalInt)
	sort.Sort(twoDimensionalInt)
	fmt.Printf("twoDimensionalInt: %v\n", twoDimensionalInt)
}

/** string排序 */
func stringSort() {
	// sort.StringSlice 声明方式等同于 []string{"100", "42", "41", "3", "2"}
	// 比较会先比较第一位，如果相同再比较第二位。
	ls := sort.StringSlice{"100", "42", "41", "3", "2"}
	fmt.Printf("ls 排序前: %v\n", ls)
	sort.Strings(ls)
	fmt.Printf("ls 排序后: %v\n", ls)

	// 汉字排序，依次比较byte大小
	ls1 := sort.StringSlice{"你", "好", "世", "界"}
	fmt.Printf("ls1 排序前: %v\n", ls1)
	sort.Strings(ls1)
	fmt.Printf("ls1 排序后: %v\n", ls1)
	// 遍历每个汉字为字节，查看汉字字节的大小
	for _, val := range ls1 {
		fmt.Printf("[]byte(val): %v %v \n", val, []byte(val))
	}
}

/** float64排序 */
func float64Sort() {
	f := []float64{1.1, 4.4, 5.5, 2.2, 3.3}
	sort.Float64s(f)
	fmt.Printf("f: %v \n", f)
}

/** 自定义排序 */
type NewInts []int

// 返回集合元素长度
func (n NewInts) Len() int {
	return len(n)
}

// 返回索引i的元素是否比索引j的元素小
func (n NewInts) Less(i, j int) bool {
	fmt.Printf("Less: %v, %v, %v, %v \n", i, j, n[i] < n[j], n)
	return n[i] < n[j]
}

// 交换 i 和 j 的值
func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func customSort() {
	ints := []int{1, 3, 2}
	sort.Sort(NewInts(ints))
	fmt.Printf("ints: %v\n", ints)
}

/** 基础 Int类型 排序 */
func basicSort() {
	ints := []int{2, 4, 1, 3}
	sort.Ints(ints)
	fmt.Printf("ints: %v\n", ints)
}
