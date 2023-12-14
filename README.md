# 桓创错误处理包
## 简介
错误由class派生，具有code，msg，cause，stack属性。

### code
表示错误类型，一般由class定义。

### msg
具体的错误，`.Error()`和`.Message`返回相同。

### cause
上级错误原因。

### stack
错误堆栈。

## 默认的class
默认的class是base，可以使用`errors.Errorf`和`errors.New`快速创建。

## 案例
见`sample`目录下
