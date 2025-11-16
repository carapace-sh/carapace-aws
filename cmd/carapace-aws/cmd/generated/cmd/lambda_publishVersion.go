package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_publishVersionCmd = &cobra.Command{
	Use:   "publish-version",
	Short: "Creates a [version](https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html) from the current code and configuration of a function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_publishVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_publishVersionCmd).Standalone()

		lambda_publishVersionCmd.Flags().String("code-sha256", "", "Only publish a version if the hash value matches the value that's specified.")
		lambda_publishVersionCmd.Flags().String("description", "", "A description for the version to override the description in the function configuration.")
		lambda_publishVersionCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_publishVersionCmd.Flags().String("revision-id", "", "Only update the function if the revision ID matches the ID that's specified.")
		lambda_publishVersionCmd.MarkFlagRequired("function-name")
	})
	lambdaCmd.AddCommand(lambda_publishVersionCmd)
}
