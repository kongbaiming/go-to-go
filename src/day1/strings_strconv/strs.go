package main

import (
	"fmt"
	"strings"
)

func main() {
	//HasPrefix 判断字符串 s 是否以 prefix 开头
	//HasSuffix 判断字符串 s 是否以 suffix 结尾：
	str, str1 := "im king", "This is an example of a string"
	fmt.Printf("字符串 \"s\" 是否是 \"%s\" 的结尾\n", str)
	fmt.Printf("%t\n", strings.HasSuffix(str, "s"))
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?\n ", str1, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str1, "Th"))
	//Contains 判断字符串 s 是否包含 substr
	fmt.Printf("字符串\"%s\"是否包含字符串\"%s\"\n", str, str1)
	fmt.Printf("%t\n", strings.Contains(str1, str))
	fmt.Printf("字符串\"%s\"是否包含字符串\"%s\"\n", str1, "is")
	fmt.Printf("%t\n", strings.Contains(str1, "is"))
	/*判断子字符串或字符在父字符串中出现的位置（索引）
	Index 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：

	strings.Index(s, str string) int
	LastIndex 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：

	strings.LastIndex(s, str string) int
	如果 ch 是非 ASCII 编码的字符，建议使用以下函数来对字符进行定位：

	strings.IndexRune(s string, r rune) int*/
	fmt.Printf("在字符串\"%s\"中,\"example\"第一次出现的位置:\n", str1)
	fmt.Printf("%d\n", strings.Index(str1, "example"))
	fmt.Printf("在字符串\"%s\"中,\"is\"最后出现的位置:\n", str1)
	fmt.Printf("%d\n", strings.LastIndex(str1, "is"))

	/*Replace 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，
	如果 n = -1 则替换所有字符串 old 为字符串 new*/
	fmt.Printf("将字符串\"%s\"的前%d个字符替换为\"%s\"\n", str1, 3, "ttt")
	fmt.Println(strings.Replace(str1, "Thi", "ttt", 1))
	fmt.Printf("将字符串\"%s\"中的所有\"s\"替换为\"%s\"\n", str1, "A")
	fmt.Println(strings.Replace(str1, "s", "A", -1))
	//Count 用于计算字符串 str 在字符串 s 中出现的非重叠次数
	//Repeat 用于重复 count 次字符串 s 并返回一个新的字符串
	var newstr string
	newstr = strings.Repeat(str, 3)
	fmt.Printf(newstr)
	//ToLower 将字符串中的 Unicode 字符全部转换为相应的小写字符
	//ToUpper 将字符串中的 Unicode 字符全部转换为相应的大写字符
	var lower, upper string = " ", " "
	fmt.Printf("The original string is: %s\n", str1)
	lower = strings.ToLower(str1)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(str)
	fmt.Printf("The uppercase string is: %s\n", upper)
	/*你可以使用 strings.TrimSpace(s) 来剔除字符串开头和结尾的空白符号；
	如果你想要剔除指定字符，则可以使用 strings.Trim(s, "cut") 来将开头和结尾的 cut 去除掉。
	该函数的第二个参数可以包含任何字符，如果你只想剔除开头或者结尾的字符串，
	则可以使用 TrimLeft 或者 TrimRight 来实现*/
	var str3, new_str3 string = "  Even if I know this is clearly impossible. it's amazing    ", " "
	fmt.Printf("%s\n", str3)
	new_str3 = strings.TrimSpace(str3)
	fmt.Printf("%s\n", new_str3)
	// 剔除开始的 Even
	fmt.Printf("%s\n", strings.TrimLeft(str3, "  Even"))

	/*strings.Fields(s) 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，
	并返回一个 slice，如果字符串只包含空白符号，则返回一个长度为 0 的 slice。
	strings.Split(s, sep) 用于自定义分割符号来对指定字符串进行分割，同样返回 slice*/
	fmt.Println(strings.Fields(str1))
	fmt.Println(len(strings.Fields(str1)))
	fmt.Println(strings.Split(str1, "example"))
	fmt.Println(len(strings.Split(str1, "example")))
	//Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
	str_1 := "There's a girl but I let her get away"
	sl := strings.Fields(str_1)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s * ", val)
	}
	fmt.Println()
	str_2 := "It's all my fault cause pride got in the way"
	sl2 := strings.Split(str_2, "'")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str_3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str_3)
}
