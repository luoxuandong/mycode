#!/usr/bin/python
# -- coding: utf-8 --

# 定义
myset = {}
myset = {"dd","ddd","dd","dd","d"}
print(myset)

myset2 = set("22ab")
myset3 = set("aaaabbbbccc")
print(myset2)
print(myset3)
testlist = ["22ab", "aaaabbbbccc"]
myset4 = set(testlist)
print(myset4)

# set函数接收的是一个列表，所以字符串就会被解析成单个的字母作为元素


# 增
myset.add("dddd")
print(myset)
myset.update(["nihao","hello"])
print(myset)

# 删
myset.remove("dd") # key 不存在则抛异常
myset.discard("dd") # 如果存在则删除，不存在不会抛异常
print(myset)

# 改
# 先删再增就行

# 查
if "ddd" in myset:
  print("ddd in myset")
for x in myset:
  print(x)

# 常用运算
set1 = set(["ab", "ac", "ad"])
set2 = set(["ac", "ad", "ae"])
print(set1 - set2)
print(set1 | set2)
print(set1 & set2)