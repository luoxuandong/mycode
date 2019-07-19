#!/usr/bin/python
# -- coding: utf-8 --

with open("/home/coding/workspace/mycode/list.py", "r") as f:
  lines = f.readlines()

print(lines)
with open("/home/coding/workspace/mycode/test.txt", "w") as f:
  i=0
  while i<len(lines):
    f.write(lines[i])
    i+=1