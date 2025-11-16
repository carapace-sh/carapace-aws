package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_createPolicyCmd = &cobra.Command{
	Use:   "create-policy",
	Short: "Creates a policy of a specified type that you can attach to a root, an organizational unit (OU), or an individual Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_createPolicyCmd).Standalone()

	organizations_createPolicyCmd.Flags().String("content", "", "The policy text content to add to the new policy.")
	organizations_createPolicyCmd.Flags().String("description", "", "An optional description to assign to the policy.")
	organizations_createPolicyCmd.Flags().String("name", "", "The friendly name to assign to the policy.")
	organizations_createPolicyCmd.Flags().String("tags", "", "A list of tags that you want to attach to the newly created policy.")
	organizations_createPolicyCmd.Flags().String("type", "", "The type of policy to create.")
	organizations_createPolicyCmd.MarkFlagRequired("content")
	organizations_createPolicyCmd.MarkFlagRequired("description")
	organizations_createPolicyCmd.MarkFlagRequired("name")
	organizations_createPolicyCmd.MarkFlagRequired("type")
	organizationsCmd.AddCommand(organizations_createPolicyCmd)
}
