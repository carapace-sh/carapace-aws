package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_createAliasCmd = &cobra.Command{
	Use:   "create-alias",
	Short: "Creates an [alias](https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html) for a Lambda function version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_createAliasCmd).Standalone()

	lambda_createAliasCmd.Flags().String("description", "", "A description of the alias.")
	lambda_createAliasCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_createAliasCmd.Flags().String("function-version", "", "The function version that the alias invokes.")
	lambda_createAliasCmd.Flags().String("name", "", "The name of the alias.")
	lambda_createAliasCmd.Flags().String("routing-config", "", "The [routing configuration](https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html#configuring-alias-routing) of the alias.")
	lambda_createAliasCmd.MarkFlagRequired("function-name")
	lambda_createAliasCmd.MarkFlagRequired("function-version")
	lambda_createAliasCmd.MarkFlagRequired("name")
	lambdaCmd.AddCommand(lambda_createAliasCmd)
}
