package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getAddonInstanceCmd = &cobra.Command{
	Use:   "get-addon-instance",
	Short: "Gets detailed information about an Add On instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getAddonInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_getAddonInstanceCmd).Standalone()

		mailmanager_getAddonInstanceCmd.Flags().String("addon-instance-id", "", "The Add On instance ID to retrieve information for.")
		mailmanager_getAddonInstanceCmd.MarkFlagRequired("addon-instance-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_getAddonInstanceCmd)
}
