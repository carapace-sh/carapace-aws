package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_createEndpointCmd = &cobra.Command{
	Use:   "create-endpoint",
	Short: "Creates a global endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_createEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_createEndpointCmd).Standalone()

		events_createEndpointCmd.Flags().String("description", "", "A description of the global endpoint.")
		events_createEndpointCmd.Flags().String("event-buses", "", "Define the event buses used.")
		events_createEndpointCmd.Flags().String("name", "", "The name of the global endpoint.")
		events_createEndpointCmd.Flags().String("replication-config", "", "Enable or disable event replication.")
		events_createEndpointCmd.Flags().String("role-arn", "", "The ARN of the role used for replication.")
		events_createEndpointCmd.Flags().String("routing-config", "", "Configure the routing policy, including the health check and secondary Region..")
		events_createEndpointCmd.MarkFlagRequired("event-buses")
		events_createEndpointCmd.MarkFlagRequired("name")
		events_createEndpointCmd.MarkFlagRequired("routing-config")
	})
	eventsCmd.AddCommand(events_createEndpointCmd)
}
