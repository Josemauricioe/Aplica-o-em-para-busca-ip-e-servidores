package app

import (
	"log"
	"net"
	"fmt"
	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação a linha de comando
func Gerar() *cli.App{

	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca Ips e Nomes de Servidor na internet"
	flags :=  []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "amazon.com.br",
		},
	}	


	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPS de endereços na internet",
			Flags: flags,
			Action: buscas_ips,
		},
		{
			Name: "servidores",
			Usage: "Buscar o nome de servidores na internet",
			Flags: flags,
			Action: buscaServidores,
		},
		}	
		return app
	}


func buscas_ips(c *cli.Context)	{	
	host :=	c.String("host")

	Ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)	
	} 

	for _, ip := range Ips {
		fmt.Println(ip)
	}
}

func buscaServidores(c *cli.Context)  {
		host := c.String("host")

		servidores, erro := net.LookupNS(host)
		if erro != nil {
			log.Fatal(erro)	
		} 

		for _, servidor := range servidores {
			fmt.Println(servidor)
		}
}