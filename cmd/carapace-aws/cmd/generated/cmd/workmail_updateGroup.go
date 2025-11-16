package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_updateGroupCmd = &cobra.Command{
	Use:   "update-group",
	Short: "Updates attributes in a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_updateGroupCmd).Standalone()

	workmail_updateGroupCmd.Flags().String("group-id", "", "The identifier for the group to be updated.")
	workmail_updateGroupCmd.Flags().String("hidden-from-global-address-list", "", "If enabled, the group is hidden from the global address list.")
	workmail_updateGroupCmd.Flags().String("organization-id", "", "The identifier for the organization under which the group exists.")
	workmail_updateGroupCmd.MarkFlagRequired("group-id")
	workmail_updateGroupCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_updateGroupCmd)
}
