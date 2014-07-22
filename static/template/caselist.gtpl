<!DOCTYPE html>
<html lang="zh-cmn-Hans">
	<head>
		<meta charset="UTF-8"> 
		<link rel="stylesheet" type="text/css" href="sidebar.css" />
		<link rel="stylesheet" type="text/css" href="pure-min.css" />
		<script type="text/javascript">
 			window.onload = init;

			 function init()
			 {
				var pid;

				if(GetQueryString("pid")!=null)
				{
					pid=GetQueryString("pid");	
					
					var ajax;
					if (window.XMLHttpRequest)
					  {// code for IE7+, Firefox, Chrome, Opera, Safari
					  ajax=new XMLHttpRequest();
					  }
					else
					  {// code for IE6, IE5
					  ajax=new ActiveXObject("Microsoft.XMLHTTP");
					  }
					ajax.open("GET","/patient/brief?pid="+pid,false);
					ajax.send();
					document.getElementById("patientbrief").innerHTML=ajax.responseText;
				}
			 }
			
			function godetail(){
				var pid;
				if(GetQueryString("pid")!=null)
				{
					pid=GetQueryString("pid");	
					window.location.href="/case/detail?pid=" + pid;
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
			<div class="pure-u-1"  align="center"> 
				<div id="patientbrief"></div>
			</div>
			<hr>
			<a name="result"></a>
			<div class="pure-u-1">
				<table class="pure-table pure-table-bordered" align="center" width="100%">
					<caption>查找结果</caption>
					<thead>
						<tr>
							<th>编号</th><th>就诊时间</th><th>主诉</th><th>检查报告</th><th>诊断</th><th>医嘱</th><th>处方</th><th>操作</th>
						</tr>
					</thead>
					<tbody>
					{{range .Cases}}
						<tr>
							<td>{{.RecordID}}</td><td>{{.CreateTime}}</td><td>{{.MainComplaint}}</td><td>{{.ExamReport}}</td><td>{{.Diag}}</td><td>{{.DRR}}</td><td>{{.Presciption}}</td><td><a href="/case/detail?rid={{.RecordID}}">详情</a> | <a href="/case/detail?rid={{.RecordID}}">拷贝</a></td>
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
				</ul>
			</div>
		</div>
	</body>
</html>