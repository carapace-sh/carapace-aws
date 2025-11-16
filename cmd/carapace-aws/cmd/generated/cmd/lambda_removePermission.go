package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_removePermissionCmd = &cobra.Command{
	Use:   "remove-permission",
	Short: "Revokes function-use permission from an Amazon Web Services service or another Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_removePermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_removePermissionCmd).Standalone()

		lambda_removePermissionCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function, version, or alias.")
		lambda_removePermissionCmd.Flags().String("qualifier", "", "Specify a version or alias to remove permissions from a published version of the function.")
		lambda_removePermissionCmd.Flags().String("revision-id", "", "Update the policy only if the revision ID matches the ID that's specified.")
		lambda_removePermissionCmd.Flags().String("statement-id", "", "Statement ID of the permission to remove.")
		lambda_removePermissionCmd.MarkFlagRequired("function-name")
		lambda_removePermissionCmd.MarkFlagRequired("statement-id")
	})
	lambdaCmd.AddCommand(lambda_removePermissionCmd)
}
