package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_putRumEventsCmd = &cobra.Command{
	Use:   "put-rum-events",
	Short: "Sends telemetry events about your application performance and user behavior to CloudWatch RUM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_putRumEventsCmd).Standalone()

	rum_putRumEventsCmd.Flags().String("alias", "", "If the app monitor uses a resource-based policy that requires `PutRumEvents` requests to specify a certain alias, specify that alias here.")
	rum_putRumEventsCmd.Flags().String("app-monitor-details", "", "A structure that contains information about the app monitor that collected this telemetry information.")
	rum_putRumEventsCmd.Flags().String("batch-id", "", "A unique identifier for this batch of RUM event data.")
	rum_putRumEventsCmd.Flags().String("id", "", "The ID of the app monitor that is sending this data.")
	rum_putRumEventsCmd.Flags().String("rum-events", "", "An array of structures that contain the telemetry event data.")
	rum_putRumEventsCmd.Flags().String("user-details", "", "A structure that contains information about the user session that this batch of events was collected from.")
	rum_putRumEventsCmd.MarkFlagRequired("app-monitor-details")
	rum_putRumEventsCmd.MarkFlagRequired("batch-id")
	rum_putRumEventsCmd.MarkFlagRequired("id")
	rum_putRumEventsCmd.MarkFlagRequired("rum-events")
	rum_putRumEventsCmd.MarkFlagRequired("user-details")
	rumCmd.AddCommand(rum_putRumEventsCmd)
}
