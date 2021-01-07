# 哪吒面板

> 服务器状态监控，报警通知，被动接收，极省资源 64M 小鸡也能装 Agent。

|  哪吒面板    |   首页截图1   |   首页截图2   |
| ---- | ---- | ---- |
|   <img src="https://s3.ax1x.com/2020/12/08/DzHv6A.jpg" width="2333px" />   | ![首页截图1](https://s3.ax1x.com/2020/12/07/DvTCwD.jpg)     | <img src="https://s3.ax1x.com/2020/12/09/rPF4xJ.png" width="1600px" /> |

\>> [查看针友列表](https://www.google.com/search?q=%22powered+by+%E5%93%AA%E5%90%92%E9%9D%A2%E6%9D%BF%22&filter=0) (Google)

## 一键脚本

**建议使用 WatchTower 自动更新面板，Windows监控可以使用nssm配置自启动**

- 海外：

    ```shell
    curl -L https://raw.githubusercontent.com/naiba/nezha/master/script/install.sh -o nezha.sh && chmod +x nezha.sh
    sudo ./nezha.sh
    ```

- 国内加速：

    ```shell
    curl -L https://raw.sevencdn.com/naiba/nezha/master/script/install.sh -o nezha.sh && chmod +x nezha.sh
    sudo ./nezha.sh
    ```

## 使用说明
### 自定义代码

可以去版权、改LOGO、加统计代码等等。

- 默认主题更改进度条颜色示例

    ```
    <style>
    .ui.fine.progress> .bar {
        background-color: pink !important;
    }
    </style>
    ```
- 默认主题修改LOGO、移除版权示例（来自 [@iLay1678](https://github.com/iLay1678)，欢迎PR）

    ```
   <style>
    .right.menu>a{
    visibility: hidden;
    }
    .footer .is-size-7{
    visibility: hidden;
    }
    .item img{
    visibility: hidden;
    }
    </style>
    <script>
    window.onload = function(){
    var avatar=document.querySelector(".item img")
    var footer=document.querySelector("div.is-size-7")
    footer.innerHTML="Powered by 你的名字"
    footer.style.visibility="visible"
    avatar.src="你的方形logo地址"
    avatar.style.visibility="visible"
    }
    </script>
   ```

- hotaru 主题更改背景图片示例

    ```
    <style>
    .hotaru-cover {
        background: url(https://s3.ax1x.com/2020/12/08/DzHv6A.jpg) center;
    }
    </style>
    ```

### 报警通知

#### 灵活通知方式

`#NEZHA#` 是面板消息占位符，面板触发通知时会自动替换占位符到实际消息

Body 内容是`JSON` 格式的：**当请求类型为FORM时**，值为 `key:value` 的形式，`value` 里面可放置占位符，通知时会自动替换。**当请求类型为JSON时** 只会简进行字符串替换后直接提交到`URL`。

URL 里面也可放置占位符，请求时会进行简单的字符串替换。

参考下方的示例，非常灵活。

1. 添加通知方式

    - server酱示例
      - 备注：server酱
      - URL：https://sc.ftqq.com/SCUrandomkeys.send?text=#NEZHA#
      - 请求方式: GET
      - 请求类型: 默认
      - Body: 空
      
    - wxpusher示例，需要关注你的应用
      - 备注: wxpusher
      - URL：http://wxpusher.zjiecode.com/api/send/message
      - 请求方式: POST
      - 请求类型: JSON
      - Body: `{"appToken":"你的appToken","topicIds":[],"content":"#NEZHA#","contentType":"1","uids":["你的uid"]}`
   
2. 添加一个离线报警

    - 备注：离线通知
    - 规则：`[{"Type":"offline","Min":0,"Max":0,"Duration":10}]`
    - 启用：√

3. 添加一个监控 CPU 持续 10s 超过 50% **且** 内存持续 20s 占用低于 20% 的报警

    - 备注：CPU+内存
    - 规则：`[{"Type":"cpu","Min":0,"Max":50,"Duration":10},{"Type":"memory","Min":20,"Max":0,"Duration":20}]`
    - 启用：√

#### 报警规则说明

- Type
  - cpu、memory、swap、disk：Min/Max 数值为占用百分比
  - net_in_speed(入站网速)、net_out_speed(出站网速)、net_all_speed(双向网速)、transfer_in(入站流量)、transfer_out(出站流量)、transfer_all(双向流量)：Min/Max 数值为字节（1kb=1024，1mb = 1024*1024）
  - offline：不支持 Min/Max 参数
- Duration：持续秒数，监控比较简陋，取持续时间内的 70 采样结果

## 常见问题

### 数据备份恢复

数据储存在 `/opt/nezha` 文件夹中，迁移数据时打包这个文件夹，到新环境解压。然后执行一键脚本安装即可

### 启用 HTTPS

使用宝塔反代或者上CDN，建议 Agent配置 跟 访问管理面板 使用不同的域名，这样管理面板使用的域名可以直接套CDN，Agent配置的域名是解析管理面板IP使用的，也方便后面管理面板迁移（如果你使用IP，后面IP更换了，需要修改每个agent，就麻烦了）

### 实时通道断开(WebSocket反代)

使用反向代理时需要针对 `/ws` 路径的 WebSocket 进行特别配置以支持实时更新服务器状态。

- Nginx(宝塔)：在你的 nginx 配置文件中加入以下代码

    ```nginx
    server{

        #server_name blablabla...

        location /ws {
            proxy_pass http://ip:站点访问端口;
            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection "Upgrade";
            proxy_set_header Host $host;
        }

        #其他的 location blablabla...
    }
    ```

- CaddyServer v1（v2无需特别配置）

    ```Caddyfile
    proxy /ws http://ip:8008 {
        websocket
    }
    ```

## 社区文章

- [哪吒探针 - Windows 客户端安装](https://nyko.me/2020/12/13/nezha-windows-client.html)
- [哪吒面板，一个便携服务器状态监控面板搭建教程，不想拥有一个自己的探针吗？](https://haoduck.com/644.html)
- [哪吒面板：小鸡们的最佳探针](https://www.zhujizixun.com/2843.html) *（已过时）*
- [>>更多教程](https://www.google.com/search?q=%22%E5%93%AA%E5%90%92%E9%9D%A2%E6%9D%BF%22+%22%E6%95%99%E7%A8%8B%22) (Google)
