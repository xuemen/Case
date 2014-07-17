<!DOCTYPE html>
<html lang="zh-cmn-Hans">
	<head>
		<meta charset="UTF-8"> 
		<link rel="stylesheet" type="text/css" href="/static/css/sidebar.css" />
		<link rel="stylesheet" type="text/css" href="/static/css/pure-min.css" />
		<script type="text/javascript">
 			function formReset(){
 				document.getElementById("newpatient").reset()
			}
		</script>
	</head>
	<body>
		<!--content begin-->
		<div class="header"  align="center">
			<h1>新增病人</h1>
			<hr>
		</div>

		<form class="pure-form" id="newpatient" action="/patient/new" method="post">
			<div class="pure-g">
				<div class="pure-u-1-2" align="center">
					<fieldset>
						<div class="pure-control-group">
							<label for="id">编号</label>
							<input type="text" name="id" class="pure-input-2-3" placeholder="编号" required="required">
						</div>
						<div class="pure-control-group">
							<label for="name">姓名</label>
							<input type="text" name="name" class="pure-input-2-3" placeholder="姓名" required="required">
						</div>
						<div class="pure-control-group">
							<label for="BOD">生日</label>
							<input type="text" name="BOD" class="pure-input-2-3" placeholder="生日（格式为YYYY-MM-DD，月、日必须是两位数，十位数可以写0.）">
						</div>
						<div class="pure-control-group">
							<label for="sex">性别</label>
							<div class="pure-u-2-3">
								<input type="radio" name="sex" checked="checked" value="未知" /> 未知
								<input type="radio" name="sex" value="男" /> 男
								<input type="radio" name="sex" value="女" /> 女
							</div>
						</div>
						<div class="pure-control-group">
							<label for="Address">地址</label>
							<input type="email" name="Address" class="pure-input-2-3" placeholder="地址">
						</div>
						<br>
						<input type="submit" class="pure-button pure-button-primary pure-input-1-3" name="b" value="提交"></input>
						<input type="button" class="pure-button pure-button-primary pure-input-1-3" onclick="formReset()" value="重置"></input>
					</fieldset>
				</div>
				<div class="pure-u-1-2" align="center">
					<fieldset>
						<div class="pure-control-group">
							<label for="PMH">既往病史</label>
							<textarea name="PMH" class="pure-input-2-3" rows="3" placeholder="既往病史" ></textarea>
						</div>
						<div class="pure-control-group">
							<label for="FMH">家族病史</label>
							<textarea name="FMH" class="pure-input-2-3" rows="3" placeholder="家族病史" ></textarea>
						</div>
						<div class="pure-control-group">
							<label for="Allergies">过敏史：</label>
							<textarea name="Allergies" class="pure-input-2-3" rows="3" placeholder="过敏史" ></textarea>
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