# v2node
A v2board backend base on moddified xray-core.
一个基于修改版xray内核的V2board节点服务端。

**注意： 本项目需要搭配[修改版V2board](https://github.com/wyx2685/v2board)**

## 软件安装

### 一键安装

```
wget -N https://raw.githubusercontent.com/LOVEYIKANUOSI/v2node/main/script/install.sh && bash install.sh
```

如果你后面又 fork 到别的仓库，也可以显式指定来源仓库：

```bash
export V2NODE_GITHUB_REPO="yourname/v2node"
export V2NODE_GITHUB_BRANCH="main"
wget -N https://raw.githubusercontent.com/yourname/v2node/main/script/install.sh && bash install.sh
```

## 构建
``` bash
GOEXPERIMENT=jsonv2 go build -v -o build_assets/v2node -trimpath -ldflags "-X 'github.com/wyx2685/v2node/cmd.version=$version' -s -w -buildid="
```

## 统一限速

可以在节点配置里增加 `GlobalSpeedLimitMbps`，给该节点下的所有用户设置统一的最大带宽上限，单位是 Mbps。

为兼容旧配置，也仍然支持 `SpeedLimit`；如果两个字段同时存在，程序优先使用 `GlobalSpeedLimitMbps`。

```json
{
  "Nodes": [
    {
      "ApiHost": "https://example.com/",
      "NodeID": 1,
      "ApiKey": "your-api-key",
      "Timeout": 15,
      "GlobalSpeedLimitMbps": 30
    }
  ]
}
```

## Stars 增长记录

[![Stargazers over time](https://starchart.cc/wyx2685/v2node.svg?variant=adaptive)](https://starchart.cc/wyx2685/v2node)
