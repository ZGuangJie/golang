## 1、数组

​	一组固定长度和特定数据类型的序列，在内存中连续存储。数组通过==make==声明，也可以通过 ==:=== 声明。golang中数组长度也是类型的一部分，的数组**长度**也是**数组类型**的一部分，数组的类型僵化了长度信息，因此并不常用。

1. **声明方式**

   ```go
   var myArray [5]int
   //两种方式
   myArray = [5]int{1, 2, 3, 4, 5}
   
   myArray := [...]int{1, 2, 3, 4, 5}
   
   // second, 返回新数组的指针，相当于 *[5]int
   var arr1 = new([5]int)
   ```
   
2. 使用方式：通过索引赋值、访问、修改

## 2、切片

​	切片（slice）是对**数组**一个连续片段的引用，切片可以看作一个结构体。

```go
type slice struct{
    Len int, //当前长度
    Cap int, //最大长度（底下数组的长度）
    Arrry[], //存数据的数组
}
```



（该数组我们称之为相关数组，通常是匿名的），所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型）。这个片段是**该数组**由起始和终止索引标识的一些项的子集。需要注意的是，终止索引标识的项不包括在切片内。

​	切片是一个引用类型，也就是一个指针，保存的是一个地址。所以直接

1. **声明方式**

   ```go
   // 使用make函数创建切片，指定元素类型、初始长度和容量
   mySlice := make([]int, 5, 10)
   // 直接声明并初始化切片
   mySlice := []int{1, 2, 3, 4, 5}
   
   // 通过数组创建切片
   myArray := [5]int{1, 2, 3, 4, 5}
   mySlice := myArray[1:4] // 创建一个包含myArray[1], myArray[2], myArray[3]的切片
   ```

   ##### 为什么两个slice不能直接用==判断是否含有相等的元素？

   

2. **优点**

   ​	因为切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中 切片比数组更常用.（不需要说明长度）。

3. **bytes 包**

   ​	类型 []byte 的切片十分常见，Go 语言有一个 bytes 包专门用来解决这种类型的操作方法。bytes 包和字符串包十分类似。而且它还包含一个十分有用的类型buffer：

   ```go
   import "bytes"
   
   type Buffer struct {
       ...
   }
   //这是一个长度可变的 bytes 的 buffer，提供 Read 和 Write 方法，读写长度未知的 bytes 最好使用 buffer。
   
   // buffer可以这样定义
   var buffer bytes.Buffer
   ```

   ​	我们创建一个 buffer，通过 `buffer.WriteString(s)` 方法将字符串 s 追加到后面，最后再通过 `buffer.String()` 方法转换为 string，这种实现方式比使用 `+=` 要更节省内存和 CPU，尤其是要串联的字符串数目特别多的时候。

   ```go
   var buffer bytes.Buffer
   for {
       if s, ok := getNextString(); ok { //method getNextString() not shown here
           buffer.WriteString(s)
       } else {
           break
       }
   }
   fmt.Print(buffer.String(), "\n")
   ```

## 3、map

​	map 是一种特殊的数据结构：一种**元素对（pair）的无序集合**，pair 的一个元素是==key==，对应的另一个元素是 ==value==，所以这个结构也称为==关联数组或字典==。这是一种快速寻找值的理想结构：给定 key，对应的 value 可以迅速定位。

1. 声明方式

   map 是 **引用类型** 的： 内存用 make 方法来分配。

   ```go
   // var map1 map[keytype]valuetype, eg:
   var map1 map[string]int
   
   var map1 = make(map[keytype]valuetype)
   ```

   **不要使用 new，永远用 make 来构造 map**

2. 测试键是否存在

   ```go
   _, ok := map1[key1] // 如果key1存在则ok == true，否则ok为false
   // 常结合if混合使用
   if _, ok := map1[key1]; ok {
       // ...
   }
   ```

3. 删除key

   ```go
   // 直接使用，即使不存在也不会产生错误
   delete(map1, key1)
   ```

   

4. for-range的配套用法

   使用for循环构造map：

   ```go
   // 第一个返回值 key 是 map 中的 key 值，第二个返回值则是该 key 对应的 value 值；
   for key, value := range map1 {
       ...
   }
   
   // 如果你只关心值，可以这么使用：
   for _, value := range map1 {
       ...
   }
   
   // 如果只想获取 key，你可以这么使用:
   for key := range map1 {
       fmt.Printf("key is: %d\n", key)
   }
   ```

   

5. 

   