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
				if(GetQueryString("rid")!=null)
				{
					rid=GetQueryString("rid");
					ajax.open("GET","/case/info?rid="+rid,false);
					ajax.send();
					document.getElementById("caseinfo").innerHTML=ajax.responseText;
				}
				
				var pid;
				if(GetQueryString("pid")!=null)
				{
					pid=GetQueryString("pid");	
				}else
				{
					pid=document.getElementById("pid").value;	
				}
				
				ajax.open("GET","/patient/info?pid="+pid,false);
				ajax.send();
				document.getElementById("patientinfo").innerHTML=ajax.responseText;
				
				var myDate = new Date();
				document.getElementById("recordtime").innerHTML = myDate.toLocaleString( );
				
				MainComplaint = document.getElementById('MainComplaint');
				ExamReport = document.getElementById('ExamReport');
				Diag = document.getElementById('Diag');
				DRR = document.getElementById('DRR');
				Presciption = document.getElementById('Presciption');
				Notes = document.getElementById('Notes');
				
				
				ExamReport.removeAttribute("disabled");
				Diag.removeAttribute("disabled");
				DRR.removeAttribute("disabled");
				Presciption.removeAttribute("disabled");
				Notes.removeAttribute("disabled");
			 }
			
			function formReset(){
 				document.getElementById("newcase").reset();
				
				document.getElementById("MainComplaint").value = "";
				document.getElementById("ExamReport").value = "";
				document.getElementById("Diag").value = "";
				document.getElementById("DRR").value = "";
				document.getElementById("Presciption").value = "";
				document.getElementById("Notes").value = "";
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
			<h1>填写病历</h1>
			<hr>
		</div>

		<form class="pure-form" id="newcase" action="/case/new" method="post">
			<div class="pure-g">
				<div class="pure-u-1-3">
					<div id="patientinfo"></div>
					<fieldset>
						<br>
						<input type="submit" class="pure-button pure-button-primary pure-input-1-3" name="b" value="提交"></input>
						<input type="button" class="pure-button pure-button-primary pure-input-1-3" onclick="formReset()" value="重置"></input>
					</fieldset>
				</div>
				<div class="pure-u-2-3" align="center">
					
					<fieldset>
						<div id="caseinfo">
							就诊日期：
							<script language=JavaScript> 
							today=new Date(); 
							document.write(
								today.getYear(),"年", 
								today.getMonth()+1,"月", 
								today.getDate(),"日"); 
							</script><br>
							<input type="hidden" id="r_pid" name="r_pid" value=""></input>
							<div class="pure-control-group">
								<label for="<MainComplaint">主&nbsp;&nbsp;&nbsp;&nbsp;诉</label>
								<textarea name="MainComplaint" class="pure-input-2-3" rows="3" placeholder="主诉" ></textarea>
							</div>
							<div class="pure-control-group">
								<label for="ExamReport">检查报告</label>
								<textarea name="ExamReport" class="pure-input-2-3" rows="3" placeholder="检查报告" ></textarea>
							</div>
							<div class="pure-control-group">
								<label for="Diag">诊&nbsp;&nbsp;&nbsp;&nbsp;断</label>
								<input type="text" name="Diag" class="pure-input-2-3" placeholder="诊断"></input>
							</div>
							<div class="pure-control-group">
								<label for="DRR">医&nbsp;&nbsp;&nbsp;&nbsp;嘱</label>
								<textarea name="DRR" class="pure-input-2-3" rows="3" placeholder="医嘱" ></textarea>
							</div>
							<div class="pure-control-group">
								<label for="Presciption">处&nbsp;&nbsp;&nbsp;&nbsp;方</label>
								<textarea name="Presciption" class="pure-input-2-3" rows="3" placeholder="处方" ></textarea>
							</div>
							<div class="pure-control-group">
								<label for="Notes">备&nbsp;&nbsp;&nbsp;&nbsp;注</label>
								<textarea name="Notes" class="pure-input-2-3" rows="3" placeholder="备注" ></textarea>
							</div>
						</div>
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
					<li class="devider"></li>
					<li onclick="window.open('http://www.processon.com/myteams/539577890cf21885c69f20b3#diagrams')">设计图样</li>
					<li onclick="window.open('http://git.oschina.net/hyg/Case/issues')">质量反馈</li>
				</ul>
			</div>
		</div>
	</body>
</html>