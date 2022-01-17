package exec

// Loop executa infinitamente a função passada por parametro
func Loop(code func()) {
	While(func() bool {
		code()
		return true
	})
}

// While executa a função passada por parametro enquanto o valor de retorno for verdadeiro
func While(code func() bool) {
	for {
		ok := while(code)
		if ok == false {
			break
		}
	}
}

func while(code func() bool) (b bool) {
	defer func() {
		err := recover()
		if err != nil {
			b = true
			return
		}
	}()

	return code()
}
