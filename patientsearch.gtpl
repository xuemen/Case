<html>
<head>
<title>查找病人</title>
<meta charset="UTF-8"> 
</head>
<body>
<h3>查找病人</h3>

<table>
<tr>
<td>
<form id="searchpatient" action="/patient/search" method="post">
	编号：<input type="text" name="id"><br />
    姓名：<input type="text" name="name"><br />
	性别：<input type="radio" name="sex" value="Male" /> 男
		  <input type="radio" name="sex" value="Female" /> 女	<br />
	生日：<input type="date" name="BOD" />（格式：yyyy-mm-dd）<br />
	<br /><br />

    <input type="submit" name="s" value="查找"></input><br />
	输入单项即可查询<br /><br />
	<input type="submit" name="s" value="24小时内就诊"></input><br />
	<input type="submit" name="s" value="7天内就诊"></input><br />
	<br /><br />
	<a href="/patient/search?s=24h"><button>24小时内就诊</button></a><br />
	<a href="/patient/search?s=7d"><button>7天内就诊</button></a><br />
</form>
</td>
<td>
	<table>
	<tr>
	<th>编号</th><th>姓名</th><th>性别</th><th>生日</th><th>地址</th><th>最近就诊时间</th><th>以往诊断</th><th>操作</th>
	</tr>
	<tr>
	{{range .}}
	<td>{{.PatientID}}</td><td>{{.Name}}</td><td>{{.Sex}}</td><td>{{.BOD}}</td><td>{{.Address}}</td><td>{{.CreateTime}}</td><td>{{.Diag}}</td><td><a href="patient/detail?id={{.PatientID}}">查看</a></td>
	{{end}}
	</tr>
	</table>
</td>
</tr>
</table>
</body>
</html>