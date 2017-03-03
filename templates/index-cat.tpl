<html lang="cat" >
<head>
<title>Hola {{ .IP }}</title>
<style>
body{font-size:1em;text-align:center;}
table {font-size:1.3em;border:1px #ccc solid;margin:10px auto;}
</style>
</head>
<body>
<nav>
	<span>Language</span>
	<a href="/">English</a>
</nav>
<table rules="rows" cellspacing="10" border="8">
<tbody>
	<tr><td>IP</td><td>{{ .IP }}</td> </tr>
	<tr><td>País</td><td>{{ .Country }}</td> 
	<tr><td>Ciutat</td><td>{{ .City }}</td></tr>
</tbody>
</table>
</body>
</html>
