package cmd

import (
	"context"
	"fmt"
	"log"
	"mathOp_service/service"
	"strconv"

	valid "github.com/asaskevich/govalidator"
	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:   "sqrt",
	Short: "Calculates the square-root of a number",
	RunE: func(cmd *cobra.Command, args []string) error {
		var isValid, err = ValidateSqrtArgs(args)
		if isValid {
			sqrt(context.Background(), args)
			return nil
		}
		fmt.Println(err)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}

// This method will call the service to
// exceute the subtraction op on two numbers
func sqrt(ctx context.Context, args []string) {
	X, _ := strconv.ParseUint(args[0], 10, 64)

	req := &service.SqrtRequest{X: float64(X)}
	answer, err := client.Sqrt(ctx, req)
	if err != nil {
		log.Printf("Couldn't execute the square-root operation   #%v ", err)
	}

	fmt.Println(answer.Answer)
}

func ValidateSqrtArgs(args []string) (bool, string) {
	if len(args) != 1 {
		return false, "Enter 1 number only"
	}

	if valid.IsInt(args[0]) {
		return true, ""
	}

	return false, "Make sure the argument is a number"

}
