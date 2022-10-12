# Lab 05 Service Mesh

Setup 基本環境
```
oc apply -f 01-bookinfo.yml -f 02-bookinfo-gateway.yml -f 03-destination-rule-all.yml
```

建議事項
1. 測試 05 項目時建議刪除 detail service 物件以便更容易看出百分比