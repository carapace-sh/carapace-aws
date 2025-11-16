package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listAlertsCmd = &cobra.Command{
	Use:   "list-alerts",
	Short: "List the alerts for a channel with optional filtering based on alert state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listAlertsCmd).Standalone()

	medialive_listAlertsCmd.Flags().String("channel-id", "", "The unique ID of the channel")
	medialive_listAlertsCmd.Flags().String("max-results", "", "The maximum number of items to return")
	medialive_listAlertsCmd.Flags().String("next-token", "", "The next pagination token")
	medialive_listAlertsCmd.Flags().String("state-filter", "", "Specifies the set of alerts to return based on their state.")
	medialive_listAlertsCmd.MarkFlagRequired("channel-id")
	medialiveCmd.AddCommand(medialive_listAlertsCmd)
}
