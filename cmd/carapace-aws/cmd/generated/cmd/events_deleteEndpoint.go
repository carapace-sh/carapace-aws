package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_deleteEndpointCmd = &cobra.Command{
	Use:   "delete-endpoint",
	Short: "Delete an existing global endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_deleteEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_deleteEndpointCmd).Standalone()

		events_deleteEndpointCmd.Flags().String("name", "", "The name of the endpoint you want to delete.")
		events_deleteEndpointCmd.MarkFlagRequired("name")
	})
	eventsCmd.AddCommand(events_deleteEndpointCmd)
}
