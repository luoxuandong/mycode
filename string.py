#!/usr/bin/python
# -- coding: utf-8 --

#  增
print("no1:"+"no2\n")

# 删
# Python没有删除子串的方法，只有字符串截取
# 那么要实现删除子串，查出来删就好了
str1 = "hello, golang"
str2 = "golang"
str_index = str1.find(str2)
str3 = str1[:str_index] + str1[str_index+len(str2):]
print(str3)
# 还可以用index方法查，但是这个方法如果没有子串则会报错


# 改
print("hello, golang".replace("golang", "php"))

# 查

