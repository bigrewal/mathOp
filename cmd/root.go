package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"mathOp_service/service"
	"os"

	"github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

// "MathOp service Url"
type conf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mathOp",
	Short: "mathsOps does basic math operations",
	Long:  "mathsOps does basic math operations, right now it only supports addition and subtraction.",
}

var client service.MathOpServiceClient

func init() {
	var c conf
	c.getConf()

	conn, err := grpc.Dial(c.Host+c.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the MathOp service: %v", err)
		os.Exit(1)
	}
	client = service.NewMathOpServiceClient(conn)
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
