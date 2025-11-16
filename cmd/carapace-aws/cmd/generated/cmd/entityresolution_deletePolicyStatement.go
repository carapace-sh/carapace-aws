package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_deletePolicyStatementCmd = &cobra.Command{
	Use:   "delete-policy-statement",
	Short: "Deletes the policy statement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_deletePolicyStatementCmd).Standalone()

	entityresolution_deletePolicyStatementCmd.Flags().String("arn", "", "The ARN of the resource for which the policy need to be deleted.")
	entityresolution_deletePolicyStatementCmd.Flags().String("statement-id", "", "A statement identifier that differentiates the statement from others in the same policy.")
	entityresolution_deletePolicyStatementCmd.MarkFlagRequired("arn")
	entityresolution_deletePolicyStatementCmd.MarkFlagRequired("statement-id")
	entityresolutionCmd.AddCommand(entityresolution_deletePolicyStatementCmd)
}
