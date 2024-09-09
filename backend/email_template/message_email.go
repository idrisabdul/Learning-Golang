package email_template

func Feedback(code string, author string, version string, rating string, by string, title string, comment string) string {
	message := `
	<h3>` + title + `</h3>
	<table>
		<tr>
			<td>Knowledge ID</td>
			<td>: ` + code + `</td>
		</tr>
		<tr>
			<td>Author</td>
			<td>: ` + author + `</td>
		</tr>
		<tr>
			<td>Version</td>
			<td>: ` + version + `.0</td>
		</tr>
		<tr>
			<td>Rating</td>
			<td>: ` + rating + `</td>
		</tr>
		<tr>
			<td>Comment</td>
			<td>: ` + comment + `.0</td>
		</tr>
		<tr>
			<td>Rated By</td>
			<td>: ` + by + `</td>
		</tr>
	</table>`
	content := BodyEmail(message)
	return content
}

func Report(code string, author string, version string, by string, title string) string {
	message := `
	<h3>` + title + `</h3>
	<table>
		<tr>
			<td>Knowledge ID</td>
			<td>: ` + code + `</td>
		</tr>
		<tr>
			<td>Author</td>
			<td>: ` + author + `</td>
		</tr>
		<tr>
			<td>Version</td>
			<td>: ` + version + `.0</td>
		</tr>
		<tr>
			<td>Reported By</td>
			<td>: ` + by + `</td>
		</tr>
	</table>`
	content := BodyEmail(message)
	return content
}
