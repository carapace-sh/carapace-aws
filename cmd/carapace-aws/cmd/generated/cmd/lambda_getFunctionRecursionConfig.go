package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getFunctionRecursionConfigCmd = &cobra.Command{
	Use:   "get-function-recursion-config",
	Short: "Returns your function's [recursive loop detection](https://docs.aws.amazon.com/lambda/latest/dg/invocation-recursion.html) configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getFunctionRecursionConfigCmd).Standalone()

	lambda_getFunctionRecursionConfigCmd.Flags().String("function-name", "", "")
	lambda_getFunctionRecursionConfigCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_getFunctionRecursionConfigCmd)
}
