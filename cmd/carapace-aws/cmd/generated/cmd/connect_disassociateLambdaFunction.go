package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateLambdaFunctionCmd = &cobra.Command{
	Use:   "disassociate-lambda-function",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateLambdaFunctionCmd).Standalone()

	connect_disassociateLambdaFunctionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_disassociateLambdaFunctionCmd.Flags().String("function-arn", "", "The Amazon Resource Name (ARN) of the Lambda function being disassociated.")
	connect_disassociateLambdaFunctionCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_disassociateLambdaFunctionCmd.MarkFlagRequired("function-arn")
	connect_disassociateLambdaFunctionCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_disassociateLambdaFunctionCmd)
}
