package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Gets the resource-based permission policy that is set for the given environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_getResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_getResourcePolicyCmd).Standalone()

		migrationHubRefactorSpaces_getResourcePolicyCmd.Flags().String("identifier", "", "The Amazon Resource Name (ARN) of the resource associated with the policy.")
		migrationHubRefactorSpaces_getResourcePolicyCmd.MarkFlagRequired("identifier")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_getResourcePolicyCmd)
}
