package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getAliasCmd = &cobra.Command{
	Use:   "get-alias",
	Short: "Returns details about a Lambda function [alias](https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getAliasCmd).Standalone()

	lambda_getAliasCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_getAliasCmd.Flags().String("name", "", "The name of the alias.")
	lambda_getAliasCmd.MarkFlagRequired("function-name")
	lambda_getAliasCmd.MarkFlagRequired("name")
	lambdaCmd.AddCommand(lambda_getAliasCmd)
}
