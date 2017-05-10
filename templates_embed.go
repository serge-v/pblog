package main
const (
main_template = `<!DOCTYPE html>
<html>

<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title>FTP viewer</title>
<style type="text/css">
body {
	line-height:1.6;
	font-size:18px;
	color:#444;
}

a {
	color:#375eab;
	text-decoration: none
}

a:hover {
	text-decoration: underline;
}
</style>
</head>

<body>

	<table border="0">
		<tr>
			<td width="80px">&nbsp;</td>
			<td width="50px">{prev_href}</td>
			<td width="50px">{next_href}</td>
			<td width="50px">&nbsp;</td>
			<td width="80px"><a href="toc.html">contents</a></td>
			<td width="80px">{zip_href}</td>
		</tr>
	</table>

	<h2>{title}</h2>

	{contents}

	<table border="0">
		<tr>
			<td width="80px">&nbsp;</td>
			<td width="50px">{prev_href}</td>
			<td width="50px">{next_href}</td>
			<td width="50px">&nbsp;</td>
			<td width="80px"><a href="toc.html">contents</a></td>
			<td width="80px">{zip_href}</td>
		</tr>
	</table>

</body>
</html>
`

toc_template = `<!DOCTYPE html>
<html>

<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title>Photo log</title>
<style type="text/css">
body {
	margin:10px auto;
	line-height:1.6;
	font-size:18px;
	color:#444;
}

h1,h2,h3,h4 {
	line-height:1.2
}

a {
	color:#375eab;
	text-decoration: none
}

a:hover {
	text-decoration: underline;
}
</style>
</head>

<body>

<h2>Contents</h2>

<ul>
{{range .}}<li><a href="{{.Link}}">{{.Date}} {{.Title}}</li>
{{end}}
</ul>

</body>
</html>
`

)
