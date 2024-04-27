package main

func CountBytes(file []byte) int {
	return len(file)
}

func CountLines(file []byte) (lineCount int) {
	for i := 0; i < len(file); i++ {
		if file[i] == '\n' {
			lineCount++
		}
		if file[i] == '\r' {
			if i+1 < len(file) {
				if file[i+1] == '\n' {
					lineCount++
					i++
					continue
				}
			}
			lineCount++
		}
	}
	return
}
