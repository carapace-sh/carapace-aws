package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_deleteAliasCmd = &cobra.Command{
	Use:   "delete-alias",
	Short: "Deletes a Lambda function [alias](https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_deleteAliasCmd).Standalone()

	lambda_deleteAliasCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_deleteAliasCmd.Flags().String("name", "", "The name of the alias.")
	lambda_deleteAliasCmd.MarkFlagRequired("function-name")
	lambda_deleteAliasCmd.MarkFlagRequired("name")
	lambdaCmd.AddCommand(lambda_deleteAliasCmd)
}
