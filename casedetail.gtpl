<!DOCTYPE html>
<html lang="zh-cmn-Hans">
	<head>
		<meta charset="UTF-8"> 
		<link rel="stylesheet" type="text/css" href="/static/css/sidebar.css" />
		<link rel="stylesheet" type="text/css" href="/static/css/pure-min.css" />
		<script type="text/javascript">
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
		</script>	
	</head>
	<body>
		<!--content begin-->
		<div class="header"  align="center">
			<h1>病历详情</h1>
			<hr>
		</div>

		<form class="pure-form" id="casedetial" action="/case/new" method="post">
			<div class="pure-g">
				<div class="pure-u-1-3">
					编号：{{.PatientID}}<br>
					姓名：{{.Name}}<br>
					性别：{{.Sex}}<br>
					生日：{{.BOD}}<br>
					地址：{{.Address}}<br><hr>
					既往病史：{{.PMH}}<br><hr>
					家族病史：{{.FMH}}<br><hr>
					过敏史：{{.Allergies}}<br><hr>
					<fieldset>
						<br>
						<input type="hidden" name="rid" value="{{.RecordID}}"></input>
						<input type="submit" class="pure-button pure-button-primary pure-input-1-3" name="b" value="提交"></input>
						<input type="button" id="edit" class="pure-button pure-button-primary pure-input-1-3" onclick="flip()" value="编辑"></input>
					</fieldset>
				</div>
				<div class="pure-u-2-3" align="center">
					就诊日期：{{.CreateTime}}<br>
					<fieldset>
						<div class="pure-control-group">
							<label for="<MainComplaint">主&nbsp;&nbsp;&nbsp;&nbsp;诉</label>
							<textarea id="MainComplaint" name="MainComplaint" class="pure-input-2-3" rows="3" placeholder="主诉">{{.MainComplaint}}</textarea>
						</div>
						<div class="pure-control-group">
							<label for="ExamReport">检查报告</label>
							<textarea id="ExamReport" name="ExamReport" class="pure-input-2-3" rows="3" placeholder="检查报告" >{{.ExamReport}}</textarea>
						</div>
						<div class="pure-control-group">
							<label for="Diag">诊&nbsp;&nbsp;&nbsp;&nbsp;断</label>
							<input type="text" id="Diag" name="Diag" class="pure-input-2-3" placeholder="诊断" value="{{.Diag}}"></input>
						</div>
						<div class="pure-control-group">
							<label for="DRR">医&nbsp;&nbsp;&nbsp;&nbsp;嘱</label>
							<textarea id="DRR" name="DRR" class="pure-input-2-3" rows="3" placeholder="医嘱" >{{.DRR}}</textarea>
						</div>
						<div class="pure-control-group">
							<label for="Presciption">处&nbsp;&nbsp;&nbsp;&nbsp;方</label>
							<textarea id="Presciption" name="Presciption" class="pure-input-2-3" rows="3" placeholder="处方" >{{.Presciption}}</textarea>
						</div>
						<div class="pure-control-group">
							<label for="Notes">备&nbsp;&nbsp;&nbsp;&nbsp;注</label>
							<textarea id="Notes" name="Notes" class="pure-input-2-3" rows="3" placeholder="备注" >{{.Notes}}</textarea>
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
					<li onclick="window.location.href='/case/search'">查找病历</li>
				</ul>
			</div>
		</div>
	</body>
</html>