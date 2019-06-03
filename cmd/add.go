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

// Add command allows to add two numbers
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds two numbers",
	RunE: func(cmd *cobra.Command, args []string) error {
		var isValid, err = ValidateAdditionArgs(args)
		if isValid {
			add(context.Background(), args)
			return nil
		}
		fmt.Println(err)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// This method will call the service to
// exceute the addition op on two numbers
func add(ctx context.Context, args []string) {
	X, _ := strconv.ParseUint(args[0], 10, 64)
	Y, _ := strconv.ParseUint(args[1], 10, 64)

	req := &service.AddRequest{X: int64(X), Y: int64(Y)}
	answer, err := client.Add(ctx, req)
	if err != nil {
		log.Printf("Couldn't execute the addition operation   #%v ", err)
	}

	fmt.Println(answer.Answer)
}

// "Validate Addition arguments."
func ValidateAdditionArgs(args []string) (bool, string) {
	if len(args) != 2 {
		return false, "Enter two numbers"
	}

	if valid.IsInt(args[0]) && valid.IsInt(args[1]) {
		return true, ""
	}

	return false, "Make sure both arguments are numbers"

}
