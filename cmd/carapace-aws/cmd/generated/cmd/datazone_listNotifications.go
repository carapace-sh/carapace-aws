package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listNotificationsCmd = &cobra.Command{
	Use:   "list-notifications",
	Short: "Lists all Amazon DataZone notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listNotificationsCmd).Standalone()

	datazone_listNotificationsCmd.Flags().String("after-timestamp", "", "The time after which you want to list notifications.")
	datazone_listNotificationsCmd.Flags().String("before-timestamp", "", "The time before which you want to list notifications.")
	datazone_listNotificationsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
	datazone_listNotificationsCmd.Flags().String("max-results", "", "The maximum number of notifications to return in a single call to `ListNotifications`.")
	datazone_listNotificationsCmd.Flags().String("next-token", "", "When the number of notifications is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of notifications, the response includes a pagination token named `NextToken`.")
	datazone_listNotificationsCmd.Flags().String("subjects", "", "The subjects of notifications.")
	datazone_listNotificationsCmd.Flags().String("task-status", "", "The task status of notifications.")
	datazone_listNotificationsCmd.Flags().String("type", "", "The type of notifications.")
	datazone_listNotificationsCmd.MarkFlagRequired("domain-identifier")
	datazone_listNotificationsCmd.MarkFlagRequired("type")
	datazoneCmd.AddCommand(datazone_listNotificationsCmd)
}
