<!DOCTYPE html>
<html lang="zh-cmn-Hans">
	<head>
		<meta charset="UTF-8"> 
		<link rel="stylesheet" type="text/css" href="/static/css/sidebar.css" />
		<link rel="stylesheet" type="text/css" href="/static/css/pure-min.css" />
		<script type="text/javascript">
			
			function godetail(){
				var pid;
				if(GetQueryString("pid")!=null)
				{
					pid=GetQueryString("pid");	
					window.location.href="/case/new?pid=" + pid;
				}
			}
			
			function GetQueryString(name) {
			   var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)","i");
			   var r = window.location.search.substr(1).match(reg);
			   if (r!=null) return unescape(r[2]); return null;
			}
		</script>
	</head>
	<body>
		<!--content begin-->
		<div class="header" align="center">
			<h1>历史诊疗记录</h1>
			<hr>
		</div>
		<div class="pure-g"  style="padding: 0px 20px">

			<div class="pure-u-1">
				<table class="pure-table pure-table-bordered" align="center" width="100%">
					<caption>查找结果</caption>
					<thead>
						<tr>
							<th>编号</th><th>就诊时间</th><th>主诉</th><th>检查报告</th><th>诊断</th><th>操作</th>
						</tr>
					</thead>
					<tbody>
					{{range .}}
						<tr>
							<td>{{.RecordID}}</td><td>{{.CreateTime}}</td><td>{{.FourDiagInfo.StrA1}}</td><td>{{.ExamInfo}}</td><td>{{.DiagAndTreatmentInfo}}</td><td><a href="/case/new?rid={{.RecordID}}">四诊</a> | <a href="/case/exam?rid={{.RecordID}}">检查</a> | <a href="/case/dat?rid={{.RecordID}}">诊断</a></td>
						</tr>
					{{end}}
					</tbody>
				</table>
			</div>
			<hr>

			<div class="pure-u-1"  align="center"> 
				<button class="pure-button pure-button-primary" onclick="godetail()">添加新病例</button>
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