package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_listTargetsCmd = &cobra.Command{
	Use:   "list-targets",
	Short: "Returns a list of the notification rule targets for an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_listTargetsCmd).Standalone()

	codestarNotifications_listTargetsCmd.Flags().String("filters", "", "The filters to use to return information by service or resource type.")
	codestarNotifications_listTargetsCmd.Flags().String("max-results", "", "A non-negative integer used to limit the number of returned results.")
	codestarNotifications_listTargetsCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
	codestarNotificationsCmd.AddCommand(codestarNotifications_listTargetsCmd)
}
