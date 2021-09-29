# Seminario de GO - Entregable

## Integrantes
* David Nuñez
* Tomás Cepeda

## Problema
La descripción del problema se encuentra detallada en la siguiente [página de GitHub](https://github.com/juanpablopizarro/tudai2021)

## Solución
Para resolver el problema, se realizó una solución basada en expresiones regulares. 

Esta es la funcion principal que parsea el string de entrada de acuerdo al siguiente patron: 
> [Leta,Letra] [Numero,Numero] [...]
```go
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
```
Una vez parseado el string de forma exitosa, se asignan los valores a la `struct` correspondiente y luego se ejecutan diversas funciones para validar los resultados obtenidos

 
## Test de Unidad
Se alcanzó una cobertura del 100%, [Aquí](https://davidnunez950.github.io/TUDAI-Seminarios-Go/out.html#file0) se pueden observar los resultados.
