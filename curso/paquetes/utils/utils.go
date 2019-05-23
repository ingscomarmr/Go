package utils

//Publica
// Obtiene solo los caracteres que son letras de un string
func GetLettersOnly(val string) (valStr string) {

	if IsNilOrEmpty(val) {
		valStr = ""
		return
	}

	for _, v := range val {
		//fmt.Printf("Tipo: %T, ValorInt:%v, ValorStr:%s\n", v, v, string(v))
		if isLetterUnicode(v) {
			valStr += string(v)
		}
	}

	return
}

//Publica
//Indica si un string es nulo o esta vacio
func IsNilOrEmpty(s string) bool { //las funciones publicas si se pueden ocupar fuera del paquete
	if len(s) == 0 || s == "" {
		return true
	}
	return false
}

//Privada
//indica si un int esta entre el rango de letras o espacio en blanco segun la tabla unicode
func isLetterUnicode(v int32) bool { //las funciones privadas solo se puden ocupar dentro del paqute
	if v == 32 || (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
		return true
	}
	return false
}
