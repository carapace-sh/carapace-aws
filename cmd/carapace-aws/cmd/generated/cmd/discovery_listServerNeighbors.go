package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_listServerNeighborsCmd = &cobra.Command{
	Use:   "list-server-neighbors",
	Short: "Retrieves a list of servers that are one network hop away from a specified server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_listServerNeighborsCmd).Standalone()

	discovery_listServerNeighborsCmd.Flags().String("configuration-id", "", "Configuration ID of the server for which neighbors are being listed.")
	discovery_listServerNeighborsCmd.Flags().String("max-results", "", "Maximum number of results to return in a single page of output.")
	discovery_listServerNeighborsCmd.Flags().String("neighbor-configuration-ids", "", "List of configuration IDs to test for one-hop-away.")
	discovery_listServerNeighborsCmd.Flags().String("next-token", "", "Token to retrieve the next set of results.")
	discovery_listServerNeighborsCmd.Flags().Bool("no-port-information-needed", false, "Flag to indicate if port and protocol information is needed as part of the response.")
	discovery_listServerNeighborsCmd.Flags().Bool("port-information-needed", false, "Flag to indicate if port and protocol information is needed as part of the response.")
	discovery_listServerNeighborsCmd.MarkFlagRequired("configuration-id")
	discovery_listServerNeighborsCmd.Flag("no-port-information-needed").Hidden = true
	discoveryCmd.AddCommand(discovery_listServerNeighborsCmd)
}
