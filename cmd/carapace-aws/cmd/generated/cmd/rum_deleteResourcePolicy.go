package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Removes the association of a resource-based policy from an app monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rum_deleteResourcePolicyCmd).Standalone()

		rum_deleteResourcePolicyCmd.Flags().String("name", "", "The app monitor that you want to remove the resource policy from.")
		rum_deleteResourcePolicyCmd.Flags().String("policy-revision-id", "", "Specifies a specific policy revision to delete.")
		rum_deleteResourcePolicyCmd.MarkFlagRequired("name")
	})
	rumCmd.AddCommand(rum_deleteResourcePolicyCmd)
}
