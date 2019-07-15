#!/usr/bin/python
# -- coding: utf-8 --

#  增
testlist = []
testlist.append(333)
testlist.append(444)
print("append:", str(testlist))
# [333,444]

testlist.insert(1,111)
print("insert:", str(testlist))
# [333,111,444]

# 拼接
extendlist = ["dd", "ww"]
testlist2 = ["nihao", "how", "are", "you","are"]
extendlist.extend(testlist2)
print(extendlist)

# 删
extendlist.remove("dd")
print(extendlist)

# 判断某个item是否在list中
objindex = extendlist.index("you")
print(objindex)
# objindex = extendlist.index("dd") # yichang
# print(objindex)

if "you" in extendlist:
  print("you in")
else:
  print("you not in")


# 判断两个list是否相等
cmplist1 = [22,33]
cmplist2 = [33,44]
cmplist3 = [11,22]
com1 = cmp(cmplist1, cmplist2)
com2 = cmp(cmplist2, cmplist3)
com3 = cmp(cmplist1, cmplist1)
print("com:", str(com1), str(com2), str(com3))
# 相等为0 不等为非0


# max，min,len
print(max(cmplist1), min(cmplist1), len(cmplist1))


