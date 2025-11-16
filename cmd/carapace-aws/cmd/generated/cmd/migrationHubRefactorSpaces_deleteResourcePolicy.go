package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes the resource policy set for the environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_deleteResourcePolicyCmd).Standalone()

		migrationHubRefactorSpaces_deleteResourcePolicyCmd.Flags().String("identifier", "", "Amazon Resource Name (ARN) of the resource associated with the policy.")
		migrationHubRefactorSpaces_deleteResourcePolicyCmd.MarkFlagRequired("identifier")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_deleteResourcePolicyCmd)
}
