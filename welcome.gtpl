<html>
<head>
<title>欢迎</title>
<meta charset="UTF-8"> 
</head>
<body>
<h3>欢迎</h3>

病人：<br />
<a href="patient/search">查找</a>
<a href="patient/new">新增</a>
<br />
<br />
病案：<br />
{{range .}}
<a href="case/detial?CaseID={{.CaseID}}">{{.PatientName}} : [{{.CreateTime}}]</a> <br />
{{end}}
<br />

</body>
</html>