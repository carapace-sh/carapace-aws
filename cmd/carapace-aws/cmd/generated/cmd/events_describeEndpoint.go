package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_describeEndpointCmd = &cobra.Command{
	Use:   "describe-endpoint",
	Short: "Get the information about an existing global endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_describeEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_describeEndpointCmd).Standalone()

		events_describeEndpointCmd.Flags().String("home-region", "", "The primary Region of the endpoint you want to get information about.")
		events_describeEndpointCmd.Flags().String("name", "", "The name of the endpoint you want to get information about.")
		events_describeEndpointCmd.MarkFlagRequired("name")
	})
	eventsCmd.AddCommand(events_describeEndpointCmd)
}
