package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_listNotificationChannelsCmd = &cobra.Command{
	Use:   "list-notification-channels",
	Short: "Returns a list of notification channels configured for DevOps Guru.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_listNotificationChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_listNotificationChannelsCmd).Standalone()

		devopsGuru_listNotificationChannelsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	})
	devopsGuruCmd.AddCommand(devopsGuru_listNotificationChannelsCmd)
}
