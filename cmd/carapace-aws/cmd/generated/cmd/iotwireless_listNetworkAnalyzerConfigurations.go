package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listNetworkAnalyzerConfigurationsCmd = &cobra.Command{
	Use:   "list-network-analyzer-configurations",
	Short: "Lists the network analyzer configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listNetworkAnalyzerConfigurationsCmd).Standalone()

	iotwireless_listNetworkAnalyzerConfigurationsCmd.Flags().String("max-results", "", "")
	iotwireless_listNetworkAnalyzerConfigurationsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	iotwirelessCmd.AddCommand(iotwireless_listNetworkAnalyzerConfigurationsCmd)
}
