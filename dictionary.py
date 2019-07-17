#!/usr/bin/python
# -- coding: utf-8 --

# 定义&初始化
mydictionary = {}
mydictionary2 = {"name":"lihua", "age":22}

# 增
mydictionary["add"] = "myaddvalue"
mydictionary2["add2"] = "add2"

print(mydictionary)
print(mydictionary2)

# 删
del mydictionary2["add2"]
print(mydictionary2)
mydictionary2.clear()

# 改
mydictionary["add"] = "update"
print(mydictionary)

# 查
print(mydictionary.has_key("key"))
print(mydictionary.get("key", default="haha"))
print(mydictionary.keys())
print(mydictionary.values())
for key in mydictionary.keys():
    print(key + ":" + mydictionary[key])

# 比较
print(cmp(mydictionary, mydictionary2))

# 常用函数,len,str
print(len(mydictionary))
