package main
const (
main_template = `<!DOCTYPE html>
<html>

<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title>FTP viewer</title>
<style type="text/css">
body {
	margin:40px auto;
	max-width:650px;
	line-height:1.6;
	font-size:18px;
	color:#444;
	padding:0
	10px
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

span.sp {
	width: 100px
	display: block;
}
</style>
</head>

<body>

<div style="text-align:center">

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

</div>

</body>
</html>
`

toc_template = `<!DOCTYPE html>
<html>

<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title>FTP viewer</title>
<style type="text/css">
body {
	margin:40px auto;
	max-width:650px;
	line-height:1.6;
	font-size:18px;
	color:#444;
	padding:0
	10px
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

span.sp {
	width: 100px
	display: block;
}
</style>
</head>

<body>

<div style="text-align:center">

<h2>Contents</h2>

<div style="text-align:left">
{contents}
</div>
</div>

</body>
</html>
`

)
