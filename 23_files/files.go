package main

func main() {
	//basic file infos
	/*
		f, err := os.Open("test.txt")
		if err != nil {
			//log the error
			panic(err)
		}

		fileInfo, err := f.Stat()

		if err != nil {
			panic(err)
		}
		fmt.Println("File Name", fileInfo.Name())
		fmt.Println("Is Folder", fileInfo.IsDir())
		fmt.Println("File Size", fileInfo.Size())
		fmt.Println("File permission", fileInfo.Mode())
		fmt.Println("File modified at", fileInfo.ModTime())

	*/

	//read file : method 1
	/*
		f, err := os.Open("test.txt")
		if err != nil {
			panic(err)
		}

		defer f.Close()

		buf := make([]byte, 13)
		d, err := f.Read(buf)
		if err != nil {
			panic(err)
		}

		for i := 0; i < len(buf); i++ {
			fmt.Println("Data", d, string(buf[i]))
		}

	*/

	//read file : method2 (simplest)
	//not efficient only good for small files
	//this method load whole file in memory at once
	/*
		data, err := os.ReadFile("test.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	*/

	//read folders
	/*
		dir, err := os.Open("../")
		if err != nil {
			panic(err)
		}

		defer dir.Close()
		fileInfo, err := dir.ReadDir(-1)
		if err != nil {
			panic(err)
		}

		for _, fi := range fileInfo {
			fmt.Println(fi.Name(), fi.IsDir())
		}
	*/

	//create a file
	/*
		f, err := os.Create("example2.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		// f.WriteString("Hi Go")
		// f.WriteString(" nice language")

		bytes := []byte("Hello Golang")
		f.Write(bytes)

	*/

	//read and write to another files(streming fashion)
	/*
		sourceFile, err := os.Open("test.txt")
		if err != nil {
			panic(err)
		}
		defer sourceFile.Close()

		destFile, err := os.Create("example2.txt")
		if err != nil {
			panic(err)
		}

		defer destFile.Close()

		reader := bufio.NewReader(sourceFile)
		writer := bufio.NewWriter(destFile)

		for {
			b, err := reader.ReadByte()
			if err != nil {
				if err.Error() != "EOF" {
					panic(err)
				}
				break
			}

			e := writer.WriteByte(b)
			if e != nil {
				panic(e)
			}
		}

		writer.Flush()
		fmt.Println("Written to new file Successfully")
	*/

	//delete a file
	/*
		err := os.Remove("example2.txt")
		if err != nil {
			panic(err)
		}

		fmt.Println("File Deleted Successfully")
	*/
}
