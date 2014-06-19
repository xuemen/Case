<html>
<head>
<title>查找病人</title>
<meta charset="UTF-8"> 
</head>
<body>
<h3>查找病人</h3>


<form id="searchpatient" action="/patient/search" method="post">
	编号：<input type="text" name="id"><br />
    姓名：<input type="text" name="name"><br />
	性别：<input type="radio" name="sex" value="Male" /> 男
		  <input type="radio" name="sex" value="Female" /> 女	<br />
	生日：<input type="date" name="BOD" />（格式：yyyy-mm-dd）<br />
	（TBD：可以增加其它字段）<br /><br />

    <input type="submit" value="查找">
</form>

</body>
</html>