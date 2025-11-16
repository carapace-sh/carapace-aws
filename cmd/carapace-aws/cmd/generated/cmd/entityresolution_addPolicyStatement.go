package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_addPolicyStatementCmd = &cobra.Command{
	Use:   "add-policy-statement",
	Short: "Adds a policy statement object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_addPolicyStatementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_addPolicyStatementCmd).Standalone()

		entityresolution_addPolicyStatementCmd.Flags().String("action", "", "The action that the principal can use on the resource.")
		entityresolution_addPolicyStatementCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource that will be accessed by the principal.")
		entityresolution_addPolicyStatementCmd.Flags().String("condition", "", "A set of condition keys that you can use in key policies.")
		entityresolution_addPolicyStatementCmd.Flags().String("effect", "", "Determines whether the permissions specified in the policy are to be allowed (`Allow`) or denied (`Deny`).")
		entityresolution_addPolicyStatementCmd.Flags().String("principal", "", "The Amazon Web Services service or Amazon Web Services account that can access the resource defined as ARN.")
		entityresolution_addPolicyStatementCmd.Flags().String("statement-id", "", "A statement identifier that differentiates the statement from others in the same policy.")
		entityresolution_addPolicyStatementCmd.MarkFlagRequired("action")
		entityresolution_addPolicyStatementCmd.MarkFlagRequired("arn")
		entityresolution_addPolicyStatementCmd.MarkFlagRequired("effect")
		entityresolution_addPolicyStatementCmd.MarkFlagRequired("principal")
		entityresolution_addPolicyStatementCmd.MarkFlagRequired("statement-id")
	})
	entityresolutionCmd.AddCommand(entityresolution_addPolicyStatementCmd)
}
