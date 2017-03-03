<html lang="en">
<head>
<title>Hello {{ .IP }}! </title>
<style>
body{font-size:1em;text-align:center;}
table {font-size:1.3em;border:1px #ccc solid;margin:10px auto;}
</style>
</head>
<body>
<nav>
	<span>Language</span>
	<a href="/?lang=cat">Calata</a>
</nav>
<table rules="rows" cellspacing="10" border="8">
<tbody>
	<tr><td>IP</td><td>{{ .IP }}</td> </tr>
	<tr><td>Country</td><td>{{ .Country }}</td> 
	<tr><td>City</td><td>{{ .City }}</td></tr>
</tbody>
</table>
</body>
</html>
