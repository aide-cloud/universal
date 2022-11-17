# aidekit
用于创建AIDE 初始化项目的工具包，项目来源于aide-family-layout模板。
仓库地址如下：
```go
const (
    githubRepo = "https://github.com/aide-cloud/aide-family-layout.git"
    giteeRepo  = "https://gitee.com/aide-cloud/aide-family-layout.git"
)
```

---

## 命令获取
```bash
 go install github.com/aide-cloud/universal/kit/aidekit@latest
```

## 参数说明
```bash
-mod string
     mod name
-p string
     package path
-r string
     layout origin repo, github or gitee (default "github")
-v    version
-w string
     worker mode (default "help")
```

### 使用示例
```bash
# 创建项目
aidekit -w new -p myApp -mod 'github.com/aide-cloud/myApp'
# 网络不好的情况可以使用gitee
aidekit -w new -p myApp -mod "github.com/aide-cloud/myApp" -r gitee
```
