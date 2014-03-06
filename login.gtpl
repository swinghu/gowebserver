<html>
<head>
<title></title>
</head>
<body>

<input type="checkbox" name = "interest" value = "football">足球
<input type="checkbox" name = "interest" value= "basketball">篮球
<input type="checkbox" name = "interest" value ="tennis">网球
<form action="/login" method="post">
    用户名:<input type="text" name="username">
    密码:<input type="password" name="password">
    <input type = "hidden" name = "token" value = "{{.}}">
    <input type="submit" value="登陆">
</form>

<select name = "fruit">
<option value = "apple">apple</option>
<option value = "pear"> pear</option>
<option value = "banana">banana</option>
</select>
</body>
</html>