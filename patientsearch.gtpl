<!DOCTYPE html>
<html lang="zh-cn">
<head>
  <title>查找病人</title>
  <meta charset="UTF-8"> 
  <link rel="stylesheet" href="http://yui.yahooapis.com/pure/0.3.0/pure-min.css">
</head>
<body>
<div class="header">
	<h1>查找病人</h1>
	<hr>
</div>

<div class="pure-g">
	<div class="pure-u-1-3"> 
		<form class="pure-form pure-form-aligned" id="searchpatient" action="/patient/search" method="post">
			<fieldset>
				<div class="pure-control-group">
					<label for="id">编号</label>
					<input type="text" name="id" placeholder="请输入病人编号">
				</div>
				<div class="pure-control-group">
					<label for="name">姓名</label>
					<input type="text" name="name" placeholder="请输入病人姓名">
				</div>
				<div class="pure-control-group">
					<label for="sex">性别</label>
					<input type="radio" name="sex" checked="checked" value="All" /> 不限
					<input type="radio" name="sex" value="Male" /> 男
					<input type="radio" name="sex" value="Female" /> 女	<br>
				</div>
				<div class="pure-control-group">
					<label for="BOD">生日</label>
					<input type="text" name="BOD" placeholder="（格式：yyyy-mm-dd）">
				</div>
				<div class="pure-controls">	
					<label for="b">输入单项即可查询</label><br>			
				    <input type="submit" class="pure-button pure-button-active" name="b" value="查找"></input><br>
				</div>
					<hr>
				<div class="pure-controls">	
					<input type="submit" class="pure-button" name="b" value="24小时内就诊"></input><br>
					<input type="submit" class="pure-button" name="b" value="7天内就诊"></input><br>
				</div>
			</fieldset>
		</form>
	</div>
	<div class="pure-u-2-3">
		<table class="pure-table pure-table-bordered">
			<thead>
				<tr>
					<th>编号</th><th>姓名</th><th>性别</th><th>生日</th><th>地址</th><th>最近就诊时间</th><th>以往诊断</th><th>操作</th>
				</tr>
			</thead>
			<tbody>
			{{range .}}
				<tr>
					<td>{{.PatientID}}</td><td>{{.Name}}</td><td>{{.Sex}}</td><td>{{.BOD}}</td><td>{{.Address}}</td><td>{{.CreateTime}}</td><td>{{.Diag}}</td><td><a href="patient/detail?id={{.PatientID}}">查看</a></td>
				</tr>
			{{end}}
			</tbody>
		</table>
	</div>
</div>
</body>
</html>