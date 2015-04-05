<!DOCTYPE html>
<html lang="zh-cmn-Hans">
	<head>
		<meta charset="UTF-8"> 
		<title>用户注册</title>	
		<link rel="stylesheet" type="text/css" href="static/css/sidebar.css" />
		<link rel="stylesheet" type="text/css" href="static/css/pure-min.css" />
		<script type="text/javascript" src="static/js/openpgp.min.js"></script>
		<script type="text/javascript">
			function createkey()
			 {
				var name;
				name=document.getElementById("realname").value;
				var nickname;
				nickname=document.getElementById("username").value;
				var email;
				email=document.getElementById("email").value;
				
				var openpgp = window.openpgp;
				alert(openpgp);
				var userid = name + " (" + nickname + ") <" + email +">"
				var passphrase = document.getElementById("keypassphrase").value;;
				var opt = {numBits: 2048, userId: userid, passphrase: passphrase};
				
				alert("准备创建密钥对，可能需要几十秒。");
				openpgp.generateKeyPair(opt).then(function(key) {
					document.getElementById("pub").value = key.publicKeyArmored ;
					document.getElementById("sec").value = key.publicKeyArmored ;
					});
				
				//document.getElementById("reg").submit();  
				document.getElementById("register").disabled=false;
			 }
		</script>
	</head>
	<body>
		<!--content begin-->
		<div class="header"  align="center">
			<h1>用户注册</h1>
			<hr>
		</div>

		<form class="pure-form pure-form-aligned" id="reg" action="/register" method="post"  style="padding: 0px 20px">
			<div class="pure-g">
				<div class="pure-u-1" align="center">		
					<fieldset>
						<div class="pure-control-group">
							<label for="username">用户名</label>
							<input type="text" id="username" name="username" class="pure-input-1-2" placeholder="英文字母、数字组成，系统使用。" required="required">
						</div>
						<div class="pure-control-group">
							<label for="email">Email</label>
							<input type="text" id="email" name="email" class="pure-input-1-2" placeholder="电子邮件地址" required="required">
						</div>
						<div class="pure-control-group">
							<label for="cellphone">手机号</label>
							<input type="text" name="cellphone" class="pure-input-1-2" placeholder="手机号码" required="required">
						</div>
						<div class="pure-control-group">
							<label for="password">密码</label>
							<input type="text" name="password" class="pure-input-1-2" placeholder="用户密码" required="required">
						</div>
						<div class="pure-control-group">
							<label for="keypassphrase">密钥口令</label>
							<input type="text" id="keypassphrase" name="keypassphrase" class="pure-input-1-2" placeholder="密钥口令" required="required">
						</div>
						<div class="pure-control-group">
							<label for="realname">真实姓名</label>
							<input type="text" id="realname" name="realname" class="pure-input-1-2" placeholder="真实姓名">
						</div>
						<div class="pure-control-group">
							<label for="org">工作单位</label>
							<input type="text" name="org" class="pure-input-1-2" placeholder="工作单位">
						</div>
						<div class="pure-control-group">
							<label for="statement">自我介绍</label>
							<textarea name="statement" class="pure-input-1-2" rows="5" placeholder="自我介绍" ></textarea>
						</div>
							<input type="hidden" id="pub" name="pub"></input>
							<input type="hidden" id="sec" name="sec"></input>
						<br>
						<div class="pure-controls">
						<input type="submit" class="pure-button pure-button-primary pure-input-1-3" id="register" disabled="true" value="提交"></input>
						<input type="button" class="pure-button pure-button-primary pure-input-1-4" onclick="createkey()" value="创建密钥"></input>
						</div>
					</fieldset>
				</div>
			</div>
		</form>
		<!--content end-->
	

		<div class="nav">
			<div class="home">
			  <span><img src="favicon.ico"></img></span>
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