package main
import(
	"log"
	"os"
	"fmt"

	"github.com/urfave/cli/v2"
)

func main(){
	app:=&cli.App{
		Name:"HealthChecker",
		Usage:"A tool that checks the give domain is up or down",
		Flags:[]cli.Flag{
			&cli.StringFlag{
				Name:"domain",
				Aliases:[]string{"d"},
				Usage:"Domain Name to check website Health",
				Required: true,
			},
			&cli.StringFlag{
				Name:"port",
				Aliases:[]string{"p"},
				Usage:"port number to check website health",
				Required: false,
			},
		},
		Action: func(c*cli.Context) error{
			port:=c.String("port")
			if c.String("port")==""{
				port="80"
			}
			status:=Check(c.String("domain"),port)
			fmt.Println(status)
			return nil
		},
	}
	err:=app.Run(os.Args)
	if err!=nil{
		log.Fatal(err)
	}
}