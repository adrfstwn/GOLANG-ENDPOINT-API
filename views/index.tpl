<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Data Pegawai</title>
</head>
<body>
	<h1>Data Pegawai</h1>
	<table border="1">
		<tr>
			<th>ID</th>
			<th>Nama</th>
			<th>Alamat</th>
			<th>Jenis Kelamin</th>
			<th>Agama</th>
			<th>Status Pegawai</th>
		</tr>
		{{range .PegawaiList}}
		<tr>
			<td>{{.Id}}</td>
			<td>{{.Nama}}</td>
			<td>{{.Alamat}}</td>
			<td>{{.JenisKelamin.Nama}}</td>
			<td>{{.Agama.Nama}}</td>
			<td>{{.StatusPegawai.Nama}}</td>
		</tr>
		{{else}}
		<tr>
			<td colspan="6">Tidak ada data Pegawai.</td>
		</tr>
		{{end}}
	</table>
</body>
</html>
