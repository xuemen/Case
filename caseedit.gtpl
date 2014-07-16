<!DOCTYPE html>
<html lang="zh-cmn-Hans">
<head>
  <title>填写病历</title>
  <meta charset="UTF-8"> 
  <link rel="stylesheet" href="http://yui.yahooapis.com/pure/0.3.0/pure-min.css">
	<script type="text/javascript">
	function formReset()
	  {
	  document.getElementById("newpatient").reset()
	  }
	</script>
</head>
<body>
<div class="header"  align="center">
	<h1>填写病历</h1>
	<hr>
</div>

<form class="pure-form" id="newcase" action="/case/new" method="post">
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
				<input type="hidden" name="pid" value="{{.PatientID}}"></input>
		    	<input type="submit" class="pure-button pure-button-primary pure-input-1-3" name="b" value="提交"></input>
				<input type="button" class="pure-button pure-button-primary pure-input-1-3" onclick="formReset()" value="重置"></input>
			</fieldset>
		</div>
		<div class="pure-u-2-3" align="center">
			就诊日期：
			<script language=JavaScript> 
			today=new Date(); 
			document.write(
				today.getYear(),"年", 
				today.getMonth()+1,"月", 
				today.getDate(),"日"); 
			</script><br>
			<fieldset>
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
					<input type="text" name="Diag" class="pure-input-2-3" placeholder="诊断">
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
			</fieldset>
		
		</div>
	</div>
</form>

</body>
</html>
