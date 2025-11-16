package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateLambdaFunctionCmd = &cobra.Command{
	Use:   "associate-lambda-function",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateLambdaFunctionCmd).Standalone()

	connect_associateLambdaFunctionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_associateLambdaFunctionCmd.Flags().String("function-arn", "", "The Amazon Resource Name (ARN) for the Lambda function being associated.")
	connect_associateLambdaFunctionCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_associateLambdaFunctionCmd.MarkFlagRequired("function-arn")
	connect_associateLambdaFunctionCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_associateLambdaFunctionCmd)
}
