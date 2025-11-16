package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_updatePolicyCmd = &cobra.Command{
	Use:   "update-policy",
	Short: "Updates an existing policy with a new name, description, or content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_updatePolicyCmd).Standalone()

	organizations_updatePolicyCmd.Flags().String("content", "", "If provided, the new content for the policy.")
	organizations_updatePolicyCmd.Flags().String("description", "", "If provided, the new description for the policy.")
	organizations_updatePolicyCmd.Flags().String("name", "", "If provided, the new name for the policy.")
	organizations_updatePolicyCmd.Flags().String("policy-id", "", "The unique identifier (ID) of the policy that you want to update.")
	organizations_updatePolicyCmd.MarkFlagRequired("policy-id")
	organizationsCmd.AddCommand(organizations_updatePolicyCmd)
}
