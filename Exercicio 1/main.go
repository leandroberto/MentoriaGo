package main

import "fmt"

func main() {
	var custo_conta float64
	var pago float64
	fmt.Print("Digite o valor da conta: R$")
	fmt.Scanln(&custo_conta)
	fmt.Print("Digite o valor pago pelo cliente: R$")
	fmt.Scanln(&pago)

	if pago < custo_conta {
		fmt.Printf("O valor pago é menor que o da conta, falta R$%.2f para o cliente pagar \n", custo_conta-pago)
	}
	if pago > custo_conta {
		fmt.Printf("O valor do troco é: R$%.2f \n", pago-custo_conta)
	}
	if pago == custo_conta {
		fmt.Print("A conta foi paga integralmente, não é necessário troco \n")
	}

}
