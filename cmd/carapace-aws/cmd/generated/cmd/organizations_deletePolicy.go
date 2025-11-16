package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_deletePolicyCmd = &cobra.Command{
	Use:   "delete-policy",
	Short: "Deletes the specified policy from your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_deletePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_deletePolicyCmd).Standalone()

		organizations_deletePolicyCmd.Flags().String("policy-id", "", "The unique identifier (ID) of the policy that you want to delete.")
		organizations_deletePolicyCmd.MarkFlagRequired("policy-id")
	})
	organizationsCmd.AddCommand(organizations_deletePolicyCmd)
}
