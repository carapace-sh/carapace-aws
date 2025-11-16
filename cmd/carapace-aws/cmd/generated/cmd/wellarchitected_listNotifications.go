package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listNotificationsCmd = &cobra.Command{
	Use:   "list-notifications",
	Short: "List lens notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listNotificationsCmd).Standalone()

	wellarchitected_listNotificationsCmd.Flags().String("max-results", "", "The maximum number of results to return for this request.")
	wellarchitected_listNotificationsCmd.Flags().String("next-token", "", "")
	wellarchitected_listNotificationsCmd.Flags().String("resource-arn", "", "The ARN for the related resource for the notification.")
	wellarchitected_listNotificationsCmd.Flags().String("workload-id", "", "")
	wellarchitectedCmd.AddCommand(wellarchitected_listNotificationsCmd)
}
