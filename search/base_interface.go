package search

type simpleTable interface {
	put()      //将键值对存入表中
	get()      // 获取key对应的值
	delete()   //从表中删除key
	contains() //key在表中是否有对应的值
	isEmpty()  //是否为空
	size()     //表中的键值对的数量
	keys()     //表中的所有键的集合
}

type sortTable interface {
	min()
	max()
	floor()   //向下取整
	ceiling() //向上取整
	rank()    // key键的排名
	_select() //排名为k的元素
	deleteMin()
	deleteMax()
	size()
}
