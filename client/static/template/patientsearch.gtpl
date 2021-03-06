<!DOCTYPE html>
<html lang="zh-cmn-Hans">
	<head>
		<meta charset="UTF-8"> 
		<link rel="stylesheet" type="text/css" href="/static/css/sidebar.css" />
		<link rel="stylesheet" type="text/css" href="/static/css/pure-min.css" />
	</head>
	<body>
		<!--content begin-->
		<div class="header" align="center">
			<h1>查找病人</h1>
			<hr>
		</div>
		<div class="pure-g"  style="padding: 0px 20px">
			<div class="pure-u-1"  align="center"> 
				<form id="searchpatient" action="/patient/search#result" method="post">
					<fieldset>
						<legend>输入单项即可查询</legend>
						
						<div class="pure-g-r">
							<div class="pure-u-1-3">
								<label for="id">编号</label>
								<input type="text" name="id" placeholder="请输入病人编号">
							</div>
							
							<div class="pure-u-1-3">
								<label for="DOB">生日</label>
							<input type="text" name="DOB" placeholder="（格式：yyyy-mm-dd）">
							</div>
							
							<div class="pure-u-1-3">
								<label for="time">最近就诊</label>
								<input type="radio" name="time" checked="checked" value="All" /> 不限
								<input type="radio" name="time" value="24h" /> 24小时内
								<input type="radio" name="time" value="7d" /> 7天内
							</div>
							
							<div class="pure-u-1-3">
								<label for="name">姓名</label>
								<input type="text" name="name" placeholder="请输入病人姓名">
							</div>
							
							<div class="pure-u-1-3">
								<label for="sex">性别</label>
								<input type="radio" name="sex" checked="checked" value="All" /> 不限
								<input type="radio" name="sex" value="男" /> 男
								<input type="radio" name="sex" value="女" /> 女	<br>
							</div>
								<div class="pure-u-1-3">
								<input type="submit" class="pure-button pure-button-primary" name="b" value="查找"></input>
								<input type="button" class="pure-button pure-button-primary" onclick="formReset()" value="重置"></input>
							</div>
						</div>
					</fieldset>
				</form>
			</div>
			<hr>
			<a name="result"></a>
			<div class="pure-u-1">
				<table class="pure-table pure-table-bordered" align="center" width="100%">
					<caption>查找结果</caption>
					<thead>
						<tr>
							<th>编号</th><th>姓名</th><th>性别</th><th>生日</th><th>地址</th><th>最近就诊时间</th><th>以往诊断</th><th>病历操作</th>
						</tr>
					</thead>
					<tbody>
					{{range .}}
						<tr>
							<td>{{.PatientID}}</td><td>{{.Name}}</td><td>{{.Sex}}</td><td>{{.DOB}}</td><td>{{.Address}}</td><td>{{.CreateTime}}</td><td>{{.Diag}}</td><td><a href="/case/list?pid={{.PatientID}}">查看</a> | <a href="/case/new?pid={{.PatientID}}">新建</a></td>
						</tr>
					{{end}}
					</tbody>
				</table>
			</div>
		</div>
		<!--content end-->
	

		<div class="nav">
			<div class="home">
			  <span><img src="/favicon.ico"></img></span>
			</div>
			<div class="sidebar">
				<ul>
					<li onclick="window.location.href='/welcome'">Home</li>
					<li class="devider"></li>
					<li onclick="window.location.href='/patient/search'">查找病人</li>
					<li onclick="window.location.href='/patient/new'">新增病人</li>
					<li class="devider"></li>
					<li onclick="window.location.href='/case/list'">查找病历</li>
					<li class="devider"></li>
					<li onclick="window.open('http://www.processon.com/myteams/539577890cf21885c69f20b3#diagrams')">设计图样</li>
					<li onclick="window.open('http://git.oschina.net/hyg/Case/issues')">质量反馈</li>
				</ul>
			</div>
		</div>
	</body>
</html>