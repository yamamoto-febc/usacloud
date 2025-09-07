# Usacloud

![usacloud_logo_h.png](usacloud_logo_h.png)

[`usacloud`](https://github.com/sacloud/usacloud)は[さくらのクラウド](http://cloud.sakura.ad.jp/index.html)用の公式CLIクライアントです。  

![Test Status](https://github.com/sacloud/usacloud/workflows/Tests/badge.svg)
[![Discord](https://img.shields.io/badge/Discord-SAKURA%20Users-blue)](https://discord.gg/yUEDN8hbMf)
[![License](https://img.shields.io/github/license/sacloud/usacloud)](LICENSE.txt)
[![Version](https://img.shields.io/github/v/tag/sacloud/usacloud)](https://github.com/sacloud/usacloud/releases/latest)
![Downloads](https://img.shields.io/github/downloads/sacloud/usacloud/total)
[![Documents](https://img.shields.io/badge/docs-Documents%20%20for%20Usacloud-green)](https://docs.usacloud.jp/usacloud)

## Installation / インストール

[Documents: https://docs.usacloud.jp/usacloud/installation/start_guide](https://docs.usacloud.jp/usacloud/installation/start_guide)

### Quick Start

- [GitHub Releases](https://github.com/sacloud/usacloud/releases/latest)から自身のプラットフォーム向けのファイルをダウンロード&展開
- [さくらのクラウド ドキュメント: APIキーの新規作成・編集](https://manual.sakura.ad.jp/cloud/api/apikey.html#id3) を参照してAPIキーを作成
- `usacloud profile`コマンドでAPIキーを設定

## Usage / 基本的な使い方

コマンドは以下の書式で指定します。

    usacloud <リソース> <サブコマンド> [オプション] [対象リソースのID or 名前(部分一致) or タグ]

リソースやサブコマンド、オプションは`usacloud -h`、`usacloud <リソース名> -h`、または`usacloud <リソース名> <サブコマンド> -h`で確認できます。

#### コマンドの例

```bash
# 全ゾーンのサーバ一覧を取得
$ usacloud server list --zone=all

# 石狩第1ゾーンで名前に"example-"を含むサーバをすべてシャットダウン(オプションの位置は引数の後ろでもOK)
$ usacloud server shutdown "example-" --zone=is1a
```

### その他の使い方

Usacloud ドキュメントを参照してください。
> [Usacloud ドキュメント](https://docs.usacloud.jp/usacloud)

### コマンド一覧

```shell
usacloud -h
CLI to manage to resources on the SAKURA Cloud

Available Commands:
 === Configuration ===
    config             

 === Authentication ===
    auth-status        

 === Computing ===
    private-host       
    server             

 === Storage ===
    archive            
    auto-backup        
    cdrom              
    disk               

 === Networking ===
    bridge             
    internet           
    local-router       
    packet-filter      
    switch             

 === Networking(SubResources) ===
    interface          
    ipaddress          
    ipv6addr           
    ipv6net            
    subnet             

 === Appliance ===
    database           
    load-balancer      
    nfs                
    vpc-router         

 === SecureMobile ===
    mobile-gateway     
    sim                

 === Common service items ===
    dns                
    gslb               
    proxy-lb           
    simple-monitor     

 === Billing ===
    bill               
    coupon             

 === Lab ===
    container-registry 
    esme               

 === WebAccelerator ===
    web-accelerator    

 === Other services ===
    icon               
    license            
    note               
    ssh-key            

 === Region/Zone information ===
    region             
    zone               

 === Service/Product information ===
    disk-plan          
    internet-plan      
    license-info       
    private-host-plan  
    server-plan        
    service-class      

 === Other commands ===
    rest               
    self               
    completion         Generate completion script
    mcp                Model Context Protocol サーバを起動
```

## Model Context Protocol (MCP) サーバ

UsacloudにはModel Context Protocol (MCP) サーバ機能が組み込まれています。
MCPは、AI言語モデル（Claude、ChatGPT等）がアプリケーション固有のツール、データソース、プロンプトテンプレートへアクセスできるようにするための統一プロトコルです。

### MCPサーバの起動

```bash
usacloud mcp
```

このコマンドは標準入出力を使用してJSON-RPC形式で通信します。

### 利用可能なツール

MCPサーバは以下のツールを提供します：

#### 1. list_commands
- **説明**: usacloudの利用可能な全コマンドの一覧を取得
- **パラメータ**: 
  - `include_hidden` (bool, optional): 隠されたコマンドも含めるか

#### 2. get_command_help
- **説明**: 特定のコマンドの詳細なヘルプ情報を取得
- **パラメータ**:
  - `command` (string, required): ヘルプを取得したいコマンド (例: 'server list', 'disk create')

#### 3. execute_command
- **説明**: usacloudコマンドを安全に実行（読み取り専用操作のみ）
- **パラメータ**:
  - `command` (string, required): 実行するusacloudコマンド (例: 'server list', 'zone list')
  - `args` (array, optional): 追加の引数
  - `flags` (array, optional): 追加のフラグ
- **セキュリティ制限**: 破壊的操作（create, delete, shutdown等）は実行できません

#### 4. get_config
- **説明**: usacloudの設定情報を取得
- **パラメータ**:
  - `name` (string, optional): 取得する設定名（空の場合は全設定）

#### 5. list_zones
- **説明**: 利用可能なゾーン一覧を取得
- **パラメータ**:
  - `output_format` (string, optional): 出力形式 (json/yaml/table)

### Claude Desktopでの使用方法

Claude Desktop等のMCPクライアントでUsacloudを使用する場合は、設定ファイルに以下のような設定を追加してください：

```json
{
  "mcpServers": {
    "usacloud": {
      "command": "/path/to/usacloud",
      "args": ["mcp"]
    }
  }
}
```

## License

 `usacloud` Copyright (C) 2017-2025 The Usacloud Authors.

  This project is published under [Apache 2.0 License](LICENSE.txt).
