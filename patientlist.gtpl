<html>
<head>
<title>查找结果</title>
<meta charset="UTF-8"> 
</head>
<body>
<h3>查找结果</h3>

<table border="1" style="width:80%">
<tr><th>编号</th><th>姓名</th><th>性别</th><th>生日</th><th>操作</th></tr>
{{range .}}
<tr>
<td>{{.PatientID}}</td>
<td>{{.Name}}</td>
<td>{{.Sex}}</td>
<td>{{.BOD}}</td>
<td>
病人：<a href="./update?ID={{.PatientID}}">更新</a><br />
病案：<a href="../case/new?ID={{.PatientID}}">新增</a>
	&nbsp;|&nbsp;<a href="../case/list?ID={{.PatientID}}">浏览</a>
</td>
</tr>
{{end}}
</table>
</body>
</html>