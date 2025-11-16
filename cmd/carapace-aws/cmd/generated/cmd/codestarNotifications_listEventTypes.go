package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_listEventTypesCmd = &cobra.Command{
	Use:   "list-event-types",
	Short: "Returns information about the event types available for configuring notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_listEventTypesCmd).Standalone()

	codestarNotifications_listEventTypesCmd.Flags().String("filters", "", "The filters to use to return information by service or resource type.")
	codestarNotifications_listEventTypesCmd.Flags().String("max-results", "", "A non-negative integer used to limit the number of returned results.")
	codestarNotifications_listEventTypesCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
	codestarNotificationsCmd.AddCommand(codestarNotifications_listEventTypesCmd)
}
