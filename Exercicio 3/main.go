package main

type Usuario struct {
	cpf   string
	nome  string
	idade int
	email string
	senha string
	admin bool
}

type Produto struct {
	codProduto    int
	nome          string
	unidadeMedida string // Unidade de Medida
	categoria     string
	precoAtacado  float64
	precoVarejo   float64
	quantidade    int
}

type Venda struct {
	codVenda    int
	nomeCliente string
	cpfCliente  string
	tipoVenda   string
	valorTotal  float64
	produtos    []string
}

func main() {

	usuarios := make([]Usuario, 0)

	//Cadastro de Usuarios
	usuario1 := Usuario{
		"123.456.789-00",
		"Maycom Jaqueson ",
		24,
		"mj2020@gmail.com",
		"Biridin@123",
		false,
	}
	usuarios = append(usuarios, usuario1)

	usuario2 := Usuario{
		"012.345.678-90",
		"Eliot Silva ",
		29,
		"msebig0@gmail.com",
		"Senha@rar",
		true,
	}
	usuarios = append(usuarios, usuario2)

	//Cadastro de Produtos
	produtos := make([]Produto, 0)

	novoProduto1 := Produto{
		1,
		"Askov Satanais 1L",
		"UN",
		"Bebidas",
		10.30,
		12.29,
		32,
	}
	produtos = append(produtos, novoProduto1)

	novoProduto2 := Produto{
		2,
		"Maca Gala",
		"Kg",
		"Frutas",
		2.5,
		2.62,
		100,
	}
	produtos = append(produtos, novoProduto2)

	type item struct {
		codProduto int
		quantidade float64
	}

	// //Login
	// var login, senha string
	// logado :=  false

	// fmt.Print("Digite seu login")
	// fmt.Scanln(&login)
	// fmt.Print("Digite sua senha")
	// fmt.Scanln(&senha)

	// for _, usuario := range usuarios {
	// 	if usuario.email == login && usuario.senha == senha {
	// 		logado = true
	// 	}
	// }

	// // carrinho := make([]Produto, 0)

}
