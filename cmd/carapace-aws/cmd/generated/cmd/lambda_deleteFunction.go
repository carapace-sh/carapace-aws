package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_deleteFunctionCmd = &cobra.Command{
	Use:   "delete-function",
	Short: "Deletes a Lambda function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_deleteFunctionCmd).Standalone()

	lambda_deleteFunctionCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function or version.")
	lambda_deleteFunctionCmd.Flags().String("qualifier", "", "Specify a version to delete.")
	lambda_deleteFunctionCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_deleteFunctionCmd)
}
