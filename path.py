#!/usr/bin/python
# -- coding: utf-8 --

import os

# 构建
pa = os.path.join("/home/coding/","workspace","mycode","README.md")
print(pa)
path = pa


# 拆分
sp = os.path.split(pa)
print(sp)
print(os.path.splitext(pa))
print(os.path.splitext(sp[1]))

# 查询判断
print(os.path.isfile(pa)) #是文件
print(os.path.isdir(pa)) #是目录？
print(os.path.isabs(pa)) #是绝对路径？
print(os.path.abspath(pa)) #返回pa的绝对路径
print(os.path.basename(path)) #返回文件名
print(os.path.dirname(path)) #返回目录路径
print(os.path.exists(path)) #路径存在？

