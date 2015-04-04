##动态植入式需求
技术机制说明书

###概述
1. 包括普通用户的任何成员都可以制定需求、提交需求。
2. 用户可以在正式界面上预览新需求，参与讨论、修订和投票。

###角色
* 用户
* 需求分析者
* 需求审议者

###资产
* 页面需求提议
* 页面审议报告
* 版本方案
	* 需求提议
	* 术语表
	* 界面清单
	* 泳道图
	* 状态机
	* 用户协议
	* 基础数据清单
	* 工作计划及报酬
	* 任务承担者意见

* 版本审议报告

###接口
* 需求提议
* 需求显示
* 需求审议
* 版本封装
* 版本审议

###协作过程
* 需求提议
	1. 用户在页脚点击需求提议。
	2. 用户可以增加该页面的行为，behavior：
		* 引发行为的事件，event
		* 行为的操作脚本，script
		* 行为的下一页面，next
	3. 如果该行为的yaml字段能正确解析，客户端软件将它提交到服务器。
* 需求显示
	1. 客户端软件启动时，会试图从服务器获得各页面的信息。
	2. 所有用户在访问某页面时，它的behavior.event会列在页脚。
	3. 用户点击该behavior.event将执行behavir.script，然后跳转到behavior.next。
	4. 任何页脚都能看到生效的《版本方案》清单。
* 需求审议
	4. 用户可以在页脚提交《页面审议报告》。
	5. 用户可以在页脚对需求投票。

* 版本封装
	1. 由建模者签名的需求分析者，可以在任何页脚提交《版本方案》。
 
* 版本审议
	1. 用户可以在页脚提交《版本审议报告》。
	2. 开发者、测试者、录入者等角色必须提交《版本审议报告》。

###分配规则
* 对需求提议者、审议者的具体奖励由需求分析者在《版本方案》中规定。

###技术方案
####客户端
1. 在 /static/page/ 路径下设置同名同构的子文件夹。
2. 每个路径下保存page.yaml文件，记录页面和行为需求。
3. pageinit读取：
	* 本地 /static/page/  ：写入一个map[string]page
	* 服务端 /static/page/ ：写入另一个map[string]page
	* 服务端 /static/version/ ：写入一个map[string]version
4. 在各页脚显示：
	* 所有的version
	* 本路径下的两种page

####服务器
1. 在 /static/page/ 路径下设置同名同构的子文件夹。
2. 每个路径下保存:
	* index.yaml：保存现有的需求索引。
	* user.page.yaml：保存各用户提交的页面和行为需求。
3. 在 /static/version/ 路径下保存《版本方案》

