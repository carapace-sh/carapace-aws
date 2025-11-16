package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getPolicyCmd = &cobra.Command{
	Use:   "get-policy",
	Short: "Returns the [resource-based IAM policy](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html) for a function, version, or alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getPolicyCmd).Standalone()

	lambda_getPolicyCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function, version, or alias.")
	lambda_getPolicyCmd.Flags().String("qualifier", "", "Specify a version or alias to get the policy for that resource.")
	lambda_getPolicyCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_getPolicyCmd)
}
