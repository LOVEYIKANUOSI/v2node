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

安装脚本默认会从你自己的 fork Releases 下载二进制包，不会在服务器上源码编译。

如果你想手动指定二进制来源，也可以设置：

```bash
export V2NODE_RELEASE_REPO="LOVEYIKANUOSI/v2node"
```

## Fork 自己发版

这个仓库已经包含适用于 fork 的 GitHub Actions 发版流程。

你只需要在 GitHub 仓库里推送一个 `v*` 标签，例如：

```bash
git tag v0.4.2
git push origin v0.4.2
```

工作流会自动：

- 构建各平台压缩包
- 创建对应的 GitHub Release
- 上传 `v2node-linux-64.zip` 等安装脚本需要的二进制文件

这样后续安装脚本也可以直接从你自己的 fork Release 下载，而且不会再回退到服务器本地编译。

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
