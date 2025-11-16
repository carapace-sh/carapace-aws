package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Attaches a resource-based permission policy to the Amazon Web Services Migration Hub Refactor Spaces environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_putResourcePolicyCmd).Standalone()

		migrationHubRefactorSpaces_putResourcePolicyCmd.Flags().String("policy", "", "A JSON-formatted string for an Amazon Web Services resource-based policy.")
		migrationHubRefactorSpaces_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to which the policy is being attached.")
		migrationHubRefactorSpaces_putResourcePolicyCmd.MarkFlagRequired("policy")
		migrationHubRefactorSpaces_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_putResourcePolicyCmd)
}
