package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_updateAliasCmd = &cobra.Command{
	Use:   "update-alias",
	Short: "Updates the configuration of a Lambda function [alias](https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_updateAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_updateAliasCmd).Standalone()

		lambda_updateAliasCmd.Flags().String("description", "", "A description of the alias.")
		lambda_updateAliasCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_updateAliasCmd.Flags().String("function-version", "", "The function version that the alias invokes.")
		lambda_updateAliasCmd.Flags().String("name", "", "The name of the alias.")
		lambda_updateAliasCmd.Flags().String("revision-id", "", "Only update the alias if the revision ID matches the ID that's specified.")
		lambda_updateAliasCmd.Flags().String("routing-config", "", "The [routing configuration](https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html#configuring-alias-routing) of the alias.")
		lambda_updateAliasCmd.MarkFlagRequired("function-name")
		lambda_updateAliasCmd.MarkFlagRequired("name")
	})
	lambdaCmd.AddCommand(lambda_updateAliasCmd)
}
