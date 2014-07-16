<!DOCTYPE html>
<html lang="zh-cn">
<head>
  <title>查找病人</title>
  <meta charset="UTF-8"> 
  <link rel="stylesheet" href="http://yui.yahooapis.com/pure/0.3.0/pure-min.css">
	<script type="text/javascript">
	function formReset()
	  {
	  document.getElementById("searchpatient").reset()
	  }
	</script>
</head>
<body>
<div class="header" align="center">
	<h1>查找病人</h1>
	<hr>
</div>

<div class="pure-g">
	<div class="pure-u-1"  align="center"> 
		<form class="pure-form pure-form-stacked" id="searchpatient" action="/patient/search#result" method="post">
			<fieldset>
				<legend>输入单项即可查询</legend>
				
				<div class="pure-g-r">
					<div class="pure-u-1-3">
		                <label for="id">编号</label>
						<input type="text" name="id" placeholder="请输入病人编号">
	            	</div>
					
					<div class="pure-u-1-3">
		                <label for="BOD">生日</label>
						<input type="text" name="BOD" placeholder="（格式：yyyy-mm-dd）">
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
					<td>{{.PatientID}}</td><td>{{.Name}}</td><td>{{.Sex}}</td><td>{{.BOD}}</td><td>{{.Address}}</td><td>{{.CreateTime}}</td><td>{{.Diag}}</td><td><a href="/patient/detail?id={{.PatientID}}">查看</a><a href="/case/new?pid={{.PatientID}}">新建</a></td>
				</tr>
			{{end}}
			</tbody>
		</table>
	</div>
</div>
</body>
</html>
