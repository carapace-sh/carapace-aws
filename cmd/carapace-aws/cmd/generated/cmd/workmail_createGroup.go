package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_createGroupCmd = &cobra.Command{
	Use:   "create-group",
	Short: "Creates a group that can be used in WorkMail by calling the [RegisterToWorkMail]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_createGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_createGroupCmd).Standalone()

		workmail_createGroupCmd.Flags().Bool("hidden-from-global-address-list", false, "If this parameter is enabled, the group will be hidden from the address book.")
		workmail_createGroupCmd.Flags().String("name", "", "The name of the group.")
		workmail_createGroupCmd.Flags().Bool("no-hidden-from-global-address-list", false, "If this parameter is enabled, the group will be hidden from the address book.")
		workmail_createGroupCmd.Flags().String("organization-id", "", "The organization under which the group is to be created.")
		workmail_createGroupCmd.MarkFlagRequired("name")
		workmail_createGroupCmd.Flag("no-hidden-from-global-address-list").Hidden = true
		workmail_createGroupCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_createGroupCmd)
}
