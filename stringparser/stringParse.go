package stringparser

import (
	"errors"
	"regexp"
	"strconv"
)

// Función que añade los tipos
func ParseString(s string) (*Result, error) {
	t := parser{validTypes: make(map[string]string)}
	t.validTypes["TX"] = "[A-Z]"
	t.validTypes["NN"] = "[0-9]"
	r, err := t.getResult(s)
	return r, err
}

type Result struct {
	Type   string
	Length int
	Value  string
}

// Se utiliza una estructura para almacenar los tipos aceptados, el string de entrada, y el resultado que se irá validando
// a medida que se ejecute el código
type parser struct {
	validTypes  map[string]string
	inputString string
	Result
}

// Se define una signatura común para los métodos de las validaciones
type validParseredInputString func() (bool, error)

// Se ejecutan las distintas funciones para parsear el string y validar el resultado
func (r *parser) getResult(input string) (*Result, error) {
	r.inputString = input
	for _, f := range [3]validParseredInputString{r.parseInputString, r.validateType, r.validateValue} {
		ok, err := f()
		if !ok {
			return nil, err
		}
	}
	return &r.Result, nil
}

// Se parsean el String de entrafda utilizando el patrón de carácteres: letra, letra, número, número, número o letra u otra caracter...
// Y se asignan los valores al resultado
func (r *parser) parseInputString() (bool, error) {
	re, _ := regexp.Compile(`^([A-Z]{2})(\d{2})(.*)`)
	s := re.FindAllStringSubmatch(r.inputString, -1)
	if s == nil {
		return false, errors.New("ERROR: Invalid input string. Expected a string with Letters (L), Numbers (N), Number or Letter (X), like: LLNNXX..., but found:  " + r.inputString)
	}
	r.Type, r.Length, r.Value = func(s []string) (string, int, string) {
		length, _ := strconv.Atoi(s[2]) // Se convierte de string a número
		return s[1], length, s[3]
	}(s[0])
	return true, nil
}

// Se verifica que este permitido
func (r *parser) validateType() (bool, error) {
	_, ok := r.validTypes[r.Type]
	if !ok {
		return false, errors.New("ERROR: Invalid String Type. Unkowing type: " + r.Type)
	}
	return true, nil
}

// Se valida el valor del resultado, es decir, que cumpla con la longitud y con el tipo del string
func (r *parser) validateValue() (bool, error) {
	regexp, _ := regexp.Compile(r.getValidationFoValue())
	matched := regexp.MatchString(r.Value)
	if !matched || r.Length != len([]byte(r.Value)) {
		return false, errors.New(`ERROR: Invalid String Value. Expexted string of type "` + r.Type + `" with a length of "` + strconv.FormatInt(int64(r.Length), 10) + `", but found: "` + r.Value + `"`)
	}
	return true, nil
}

// Retorna un string que contiene el valor de la expresión regular que valida el valor del string del resultadp
func (r *parser) getValidationFoValue() string {
	t := r.validTypes[r.Type]
	return t + string("{"+strconv.FormatInt(int64(r.Length), 10)+"}")
}
