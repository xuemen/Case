<!DOCTYPE html>
<html lang="zh-cmn-Hans">
	<head>
		<meta charset="UTF-8"> 
		<link rel="stylesheet" type="text/css" href="/static/css/sidebar.css" />
		<link rel="stylesheet" type="text/css" href="/static/css/pure-min.css" />
		<script type="text/javascript">
 			window.onload = init;

			 function init()
			 {
				var ajax;
				if (window.XMLHttpRequest)
				  {// code for IE7+, Firefox, Chrome, Opera, Safari
				  ajax=new XMLHttpRequest();
				  }
				else
				  {// code for IE6, IE5
				  ajax=new ActiveXObject("Microsoft.XMLHTTP");
				  }
				var rid;
				rid=GetQueryString("rid");
				ajax.open("GET","/case/info?readonly=true&rid="+rid,false);
				ajax.send();
				document.getElementById("caseinfo").innerHTML=ajax.responseText;
				
				var pid;
				pid=document.getElementById("pid").value;
				ajax.open("GET","/patient/info?pid="+pid,false);
				ajax.send();
				document.getElementById("patientinfo").innerHTML=ajax.responseText;
			 }
				
			
			function gotolast()
			{
				var rid;
				rid=GetQueryString("rid");
				
				var intrid=parseInt(rid);
				if(intrid > 0) {
					intrid=intrid-1;
					window.location="/case/detail?rid="+intrid
				}
			}
			
			function gotonext()
			{
				var rid;
				rid=GetQueryString("rid");
				
				var intrid=parseInt(rid);
				if((document.getElementById("MainComplaint") != null)||(intrid == 0)) {
					intrid=intrid+1;
					window.location="/case/detail?rid="+intrid
				}
			}	
			
			function addnew()
			{
				var pid;
				pid=document.getElementById("pid").value;
				
				window.location="/case/new?pid="+pid
			}
			
			function copynew()
			{
				var rid;
				rid=GetQueryString("rid");
				
				window.location="/case/new?rid="+rid
			}
			
			 function flip()
			 {
				MainComplaint = document.getElementById('MainComplaint');
				ExamReport = document.getElementById('ExamReport');
				Diag = document.getElementById('Diag');
				DRR = document.getElementById('DRR');
				Presciption = document.getElementById('Presciption');
				Notes = document.getElementById('Notes');
				
				 if (MainComplaint.getAttribute("disabled") == "disabled") {
					MainComplaint.removeAttribute("disabled");
					ExamReport.removeAttribute("disabled");
					Diag.removeAttribute("disabled");
					DRR.removeAttribute("disabled");
					Presciption.removeAttribute("disabled");
					Notes.removeAttribute("disabled");
            	}
            	else {
                	MainComplaint.setAttribute("disabled","disabled");
					ExamReport.setAttribute("disabled","disabled");
					Diag.setAttribute("disabled","disabled");
					DRR.setAttribute("disabled","disabled");
					Presciption.setAttribute("disabled","disabled");
					Notes.setAttribute("disabled","disabled");
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
		<div class="header"  align="center">
			<h1>病历详情</h1>
			<hr>
		</div>

		<form class="pure-form" id="casedetial" action="/case/detail" method="post">
			<div class="pure-g">
				<div class="pure-u-1-3">
					<div id="patientinfo"></div>
					<fieldset>
						<br>
						<input type="submit" class="pure-button pure-button-primary pure-input-1-3" name="b" value="提交"></input>
						<input type="button" id="edit" class="pure-button pure-button-primary pure-input-1-3" onclick="flip()" value="编辑"></input>
						<input type="button" id="next" class="pure-button pure-button-primary pure-input-1-3" onclick="gotonext()" value="下一份"></input>
						<input type="button" id="last" class="pure-button pure-button-primary pure-input-1-3" onclick="gotolast()" value="上一份"></input>
						<input type="button" id="new" class="pure-button pure-button-primary pure-input-1-3" onclick="addnew()" value="新增空白病历"></input>
						<input type="button" id="copy" class="pure-button pure-button-primary pure-input-1-3" onclick="copynew()" value="拷贝到新病历"></input>
					</fieldset>
				</div>
				<div class="pure-u-2-3" align="center">
					
					<fieldset>
						<div id="caseinfo"></div>
					</fieldset>
				
				</div>
			</div>
		</form>
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