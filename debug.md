# Go言語でのデバックについて

	// fmt.Printlnを使用してデバッグ
	fmt.Println("strSlice:", strSlice)
	fmt.Println("counts:", counts)

	// spewパッケージを使用してデバッグ
	spew.Dump(strSlice)
	spew.Dump(counts)

	// logパッケージを使用してデバッグ
	log.Println("strSlice:", strSlice)
	log.Println("counts:", counts)

	// リフレクションを使用して型情報を表示
	fmt.Println("Type of strSlice:", reflect.TypeOf(strSlice))
	fmt.Println("Type of counts:", reflect.TypeOf(counts))