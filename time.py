#!/usr/bin/python
# -- coding: utf-8 --

import datetime
print datetime.datetime.now()
# 2018-05-08 16:53:30.101000

# 格式化输出
print datetime.datetime.now().strftime("%Y-%m-%d %H:%M")
# 2018-05-08 16:54

# 加一天，然后格式化输出
print (datetime.datetime.now()+datetime.timedelta(days=1)).strftime("%Y-%m-%d %H:%M:%S")
# 2018-05-09 16:56:07

# 减一天，然后格式化输出
print (datetime.datetime.now()+datetime.timedelta(days=-1)).strftime("%Y-%m-%d %H:%M:%S")
# 2018-05-07 16:56:59

