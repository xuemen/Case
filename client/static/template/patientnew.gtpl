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

		<form class="pure-form pure-form-aligned" id="newpatient" action="/patient/new" method="post"  style="padding: 0px 20px">
			<div class="pure-g">
				<div class="pure-u-1-2" align="center">
					<fieldset>
						
						<div class="pure-control-group">
							<label for="name">姓名：</label>
							<input type="text" name="name" class="pure-input-2-3" placeholder="姓名" required="required">
						</div>
						<div class="pure-control-group">
							<label for="scbc">就诊次数：</label>
							<input type="text" name="scbc" class="pure-input-2-3" placeholder="就诊次数" required="required">
						</div>
						<div class="pure-control-group">
							<label for="sex">性别：</label>
							<div class="pure-u-2-3">
								<input type="radio" name="sex" checked="checked" value="未知" /> 未知
								<input type="radio" name="sex" value="男" /> 男
								<input type="radio" name="sex" value="女" /> 女
							</div>
						</div>
						<div class="pure-control-group">
							<label for="DOB">生日：</label>
							<input type="date" name="DOB" class="pure-input-2-3" placeholder="生日（格式为YYYY-MM-DD，月、日必须是两位数，十位数可以写0.）">
						</div>
						<div class="pure-control-group">
							<label for="weight">体重(公斤)：</label>
							<input type="text" name="weight" class="pure-input-2-3" placeholder="体重（公斤）">
						</div>
						<div class="pure-control-group">
							<label for="marital">婚姻状态：</label>
							<div class="pure-u-2-3">
								<input type="radio" name="marital" checked="checked" value="未婚" /> 未婚
								<input type="radio" name="marital" value="已婚" /> 已婚
								<input type="radio" name="marital" value="离异" /> 离异
							</div>
						</div>
						<div class="pure-control-group">
							<label for="career">职业：</label>
							<input type="text" name="career" class="pure-input-2-3" placeholder="职业">
						</div>
						<div class="pure-control-group">
							<label for="nationality">国籍：</label>
							<input type="text" name="nationality" class="pure-input-2-3" placeholder="国籍">
						</div>
						<div class="pure-control-group">
							<label for="race">民族：</label>
							<input type="text" name="race" class="pure-input-2-3" placeholder="民族">
						</div>
						<div class="pure-control-group">
							<label for="POB">出生地：</label>
							<input type="text" name="POB" class="pure-input-2-3" placeholder="出生地">
						</div>
						<div class="pure-control-group">
							<label for="phone">联系电话：</label>
							<input type="text" name="phone" class="pure-input-2-3" placeholder="联系电话">
						</div>
						<div class="pure-control-group">
							<label for="address">家庭地址：</label>
							<input type="text" name="address" class="pure-input-2-3" placeholder="地址">
						</div>
						<div class="pure-control-group">
							<label for="postcode">邮政编码：</label>
							<input type="text" name="postcode" class="pure-input-2-3" placeholder="邮政编码">
						</div>
						<div class="pure-control-group">
							<label for="servicetime">就诊时间：</label>
							<input type="date" name="servicetime" class="pure-input-2-3" placeholder="就诊时间">
						</div>
						
					</fieldset>
				</div>
				<div class="pure-u-1-2" align="center">
					<fieldset>
						<div class="pure-control-group">
							<label for="PMH">既往病史：</label>
							<textarea name="PMH" class="pure-input-2-3" rows="3" placeholder="既往病史" ></textarea>
						</div>
						<div class="pure-control-group">
							<label for="FMH">家族病史：</label>
							<textarea name="FMH" class="pure-input-2-3" rows="3" placeholder="家族病史" ></textarea>
						</div>
						<div class="pure-control-group">
							<label for="allergies">过敏史：</label>
							<textarea name="allergies" class="pure-input-2-3" rows="3" placeholder="过敏史" ></textarea>
						</div>
						<br>
						<input type="submit" class="pure-button pure-button-primary pure-input-1-3" name="b" value="提交"></input>
						<input type="button" class="pure-button pure-button-primary pure-input-1-3" onclick="formReset()" value="重置"></input>
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