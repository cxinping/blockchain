# HLTV 爬虫系统

## 安装第三方模块
```
爬虫和页面解析
go get -u github.com/gocolly/colly/...
go get github.com/PuerkitoBio/goquery

调度 
go get github.com/robfig/cron/v3@v3.0.0

存储
go get -u github.com/jinzhu/gorm
go get -u gorm.io/driver/mysql

创建UUID的模块
go get github.com/go-basic/uuid

viper配置
go get github.com/spf13/viper
```

查询赛事
```
curl https://www.hltv.org/matches | grep 'events-container'

curl https://www.hltv.org/matches | grep 'event'
```






