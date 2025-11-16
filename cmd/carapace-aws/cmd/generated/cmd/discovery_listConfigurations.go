package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_listConfigurationsCmd = &cobra.Command{
	Use:   "list-configurations",
	Short: "Retrieves a list of configuration items as specified by the value passed to the required parameter `configurationType`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_listConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_listConfigurationsCmd).Standalone()

		discovery_listConfigurationsCmd.Flags().String("configuration-type", "", "A valid configuration identified by Application Discovery Service.")
		discovery_listConfigurationsCmd.Flags().String("filters", "", "You can filter the request using various logical operators and a *key*-*value* format.")
		discovery_listConfigurationsCmd.Flags().String("max-results", "", "The total number of items to return.")
		discovery_listConfigurationsCmd.Flags().String("next-token", "", "Token to retrieve the next set of results.")
		discovery_listConfigurationsCmd.Flags().String("order-by", "", "Certain filter criteria return output that can be sorted in ascending or descending order.")
		discovery_listConfigurationsCmd.MarkFlagRequired("configuration-type")
	})
	discoveryCmd.AddCommand(discovery_listConfigurationsCmd)
}
