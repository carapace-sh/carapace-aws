package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listDestinationsCmd = &cobra.Command{
	Use:   "list-destinations",
	Short: "List all notification destinations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listDestinationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_listDestinationsCmd).Standalone()

		iotManagedIntegrations_listDestinationsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iotManagedIntegrations_listDestinationsCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results.")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listDestinationsCmd)
}
