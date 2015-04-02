<!DOCTYPE html>
<html lang="zh-cn">
<head>
  <title>欢迎</title>
  <meta charset="UTF-8"> 
  <link rel="stylesheet" type="text/css" href="/static/css/pure-min.css" />
</head>
<body>
<div class="header" align="center">
	<h1>欢迎</h1>
	<h2>病案管理系统（单机版）</h2>
	<hr>
</div>
<div class="pure-g">

    <div class="pure-u-1-3"  align="center"> 
		病人：<br>
		<p>
		<a href="patient/search">查找</a>
		<a href="patient/new">新增</a>
		</p>
	</div>
	
    <div class="pure-u-2-3"  align="center">
		病历：<br>
		<p>
		{{range .}}
		<a href="case/detial?CaseID={{.CaseID}}">{{.PatientName}} : [{{.CreateTime}}]</a> <br>
		{{end}}
		</p>
	</div>
</div>
</body>
</html>