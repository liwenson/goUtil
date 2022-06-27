# Date
<a href="https://github.com/noogo/date">![noogo-2020](https://img.shields.io/badge/github-noogo%2Fdate-brightgreen)</a> ![noogo-2020](https://img.shields.io/badge/noogo-2020-brightgreen)

Date是一个基于time包装的一个日期包，通过此包可以快速创建日期、获取时间戳、毫秒数及最重要的日期格式化，另外你还可以继续使用time包下的所有函数（除`time.Foramt(string)`外）你可以通过以下方法快速创建一个Date对象：
- ` Now()`
- `WithTime(t time.Time)`
- `WithTimestamp(timestamp int64)`
- `WithMillisecond(millisecond int64)`
- `WithDate(year, month, date, hour, minute, second int)`

**Note**:你可以通过`Date.Format(String,...bool)`方法来对日期进行格式化，日期格式化是按照Java风格实现的，免去了Golang中`非常规`的格式化方法，这对我们使用日期格式化增加了很大的便利，以下问Java日期格式化参考表:


|字母     |日期或时间元素 |		表示	|	示例|
|:----    |:----        |:----     |:----  |
|G	  |Era 标志符					|Text	|AD|
|y	  |年				|	Year	|1996; 96  |
|M	  |年中的月份			|	Month	|July; Jul; 07    |
|w	  |年中的周数			|	Number	|27     |
|W	  |月份中的周数			|Number	|2       |
|D	  |年中的天数			|	Number	|189  |
|d	  |月份中的天数			|Number	|10        |
|F	  |月份中的星期			|Number	|2            |
|E	  |星期中的天数			|Text	|Tuesday; Tue  v
|a	  |Am/pm 标记			|	Text|	PM       |
|H	  |一天中的小时数（0-23）	|Number	|0         |
|k	  |一天中的小时数（1-24）	|Number	|24         |
|K	  |am/pm 中的小时数（0-11）|Number	|0           |
|h	  |am/pm 中的小时数（1-12）|Number	|12           |
|m	  |小时中的分钟数			|Number	|30          |
|s	  |分钟中的秒数			|Number	|55         |
|S	  |毫秒数				|	Number|	978     |
|z	  |时区				|	General time zone|	Pacific Standard Time; PST; GMT-08:00   |
|Z	  |时区				|	RFC 822 time zone|	-0800                    |


## 开始

### 获取Date

`go get -u github.com/noogo/date`

### 使用Date

```go
// get date
d:=date.Now()
//d:=date.WithTime(time.Now())
//d:=date.WithTimestamp(1586448000)
//d:=date.WithMillisecond(1586448000000)
//d:=date.WithDate(2020,04,29,0,0,0)
// get milliseconds
//milliseconds:=date.Millisecond()
// get timestamp
//timestamp:=date.Timestamp()
// date format
ret,err:=d.Format("yyyy-MM-dd HH:mm:ss EEEE",true)
if err!=nil{
    log.Fatalln(err)
}
fmt.Println(ret)
```

### 运行结果

```
2020-04-29 00:13:12 星期三
```

### 格式化说明

- G:保留字段，不支持格式化
- 年:当y的连续个数小于4时则显示缩写后的年，如2008，则会格式化为08
- 月:当M的连续个数大于3时则显示英文单词月份，如果等于3则显示英文单词缩写，否则显示数字月份，位数不足用0填充。
- 对于表格中`表示`类型为Number类型的按照统一规则显示对应数值，其余多余的格式化字符用0填充，假如当前时间为2020年1月1日，08时08分08秒，那么`mm`格式化后的分钟则为`08`，`mmm`格式化后的分钟则为`008`依次类推

- 如果`Date.Format(string,...bool)`中第二个参数传true，代表中文模式，此参数控制am/pm及星期数，对应会被格式化为`上午/下午`和`星期一`格式。

### 格式化参结果
令：当前日期为2008-08-18 18:28:38.888

|layout     |result |
|:----    |:----        |
|y|08|
|yy|08|
|yyy|08|
|yyyy|2008|
|yyyyy|2008|
|M|08|
|MM|08|
|MMM|Aug|
|MMMM|August|
|MMMMM|August|
|w|34|
|ww|34|
|www|034|
|wwww|0034|
|wwwww|00034|
|W|4|
|WW|04|
|WWW|004|
|WWWW|0004|
|WWWWW|00004|
|D|231|
|DD|231|
|DDD|231|
|DDDD|0231|
|DDDDD|00231|
|d|18|
|dd|18|
|ddd|018|
|dddd|0018|
|ddddd|00018|
|F|3|
|FF|03|
|FFF|003|
|FFFF|0003|
|FFFFF|00003|
|E|星期一(chinese)|
|EE|星期一(chinese)|
|EEE|星期一(chinese)|
|EEEE|星期一(chinese)|
|EEEEE|星期一(chinese)|
|a|下午(chinese)|
|aa|下午(chinese)|
|aaa|下午(chinese)|
|aaaa|下午(chinese)|
|aaaaa|下午(chinese)|
|E|1(standard)|
|EE|01(standard)|
|EEE|Mon(standard)|
|EEEE|Monday(standard)|
|EEEEE|Monday(standard)|
|a|PM(standard)|
|aa|PM(standard)|
|aaa|PM(standard)|
|aaaa|PM(standard)|
|aaaaa|PM(standard)|
|H|18|
|HH|18|
|HHH|018|
|HHHH|0018|
|HHHHH|00018|
|k|18|
|kk|18|
|kkk|018|
|kkkk|0018|
|kkkkk|00018|
|K|6|
|KK|06|
|KKK|006|
|KKKK|0006|
|KKKKK|00006|
|h|6|
|hh|06|
|hhh|006|
|hhhh|0006|
|hhhhh|00006|
|m|28|
|mm|28|
|mmm|028|
|mmmm|0028|
|mmmmm|00028|
|s|38|
|ss|38|
|sss|038|
|ssss|0038|
|sssss|00038|
|S|888|
|SS|888|
|SSS|888|
|SSSS|0888|
|SSSSS|00888|
|z|CST|
|zz|CST|
|zzz|CST|
|zzzz|CST|
|zzzzz|CST|
|Z|+0800|
|ZZ|+0800|
|ZZZ|+0800|
|ZZZZ|+0800|
|ZZZZZ|+0800|

# License

MIT License

Copyright (c) ![noogo-2020](https://img.shields.io/badge/noogo-2020-orange)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
