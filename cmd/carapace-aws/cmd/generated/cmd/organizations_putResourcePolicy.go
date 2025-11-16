package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Creates or updates a resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_putResourcePolicyCmd).Standalone()

	organizations_putResourcePolicyCmd.Flags().String("content", "", "If provided, the new content for the resource policy.")
	organizations_putResourcePolicyCmd.Flags().String("tags", "", "A list of tags that you want to attach to the newly created resource policy.")
	organizations_putResourcePolicyCmd.MarkFlagRequired("content")
	organizationsCmd.AddCommand(organizations_putResourcePolicyCmd)
}
