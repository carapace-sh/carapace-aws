package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_deleteAddonInstanceCmd = &cobra.Command{
	Use:   "delete-addon-instance",
	Short: "Deletes an Add On instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_deleteAddonInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_deleteAddonInstanceCmd).Standalone()

		mailmanager_deleteAddonInstanceCmd.Flags().String("addon-instance-id", "", "The Add On instance ID to delete.")
		mailmanager_deleteAddonInstanceCmd.MarkFlagRequired("addon-instance-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_deleteAddonInstanceCmd)
}
