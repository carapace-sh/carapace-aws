package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_updateEndpointCmd = &cobra.Command{
	Use:   "update-endpoint",
	Short: "Update an existing endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_updateEndpointCmd).Standalone()

	events_updateEndpointCmd.Flags().String("description", "", "A description for the endpoint.")
	events_updateEndpointCmd.Flags().String("event-buses", "", "Define event buses used for replication.")
	events_updateEndpointCmd.Flags().String("name", "", "The name of the endpoint you want to update.")
	events_updateEndpointCmd.Flags().String("replication-config", "", "Whether event replication was enabled or disabled by this request.")
	events_updateEndpointCmd.Flags().String("role-arn", "", "The ARN of the role used by event replication for this request.")
	events_updateEndpointCmd.Flags().String("routing-config", "", "Configure the routing policy, including the health check and secondary Region.")
	events_updateEndpointCmd.MarkFlagRequired("name")
	eventsCmd.AddCommand(events_updateEndpointCmd)
}
