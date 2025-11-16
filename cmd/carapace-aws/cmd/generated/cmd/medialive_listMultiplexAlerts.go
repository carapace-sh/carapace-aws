package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listMultiplexAlertsCmd = &cobra.Command{
	Use:   "list-multiplex-alerts",
	Short: "List the alerts for a multiplex with optional filtering based on alert state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listMultiplexAlertsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_listMultiplexAlertsCmd).Standalone()

		medialive_listMultiplexAlertsCmd.Flags().String("max-results", "", "The maximum number of items to return")
		medialive_listMultiplexAlertsCmd.Flags().String("multiplex-id", "", "The unique ID of the multiplex")
		medialive_listMultiplexAlertsCmd.Flags().String("next-token", "", "The next pagination token")
		medialive_listMultiplexAlertsCmd.Flags().String("state-filter", "", "Specifies the set of alerts to return based on their state.")
		medialive_listMultiplexAlertsCmd.MarkFlagRequired("multiplex-id")
	})
	medialiveCmd.AddCommand(medialive_listMultiplexAlertsCmd)
}
