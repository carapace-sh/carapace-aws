package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listWirelessGatewaysCmd = &cobra.Command{
	Use:   "list-wireless-gateways",
	Short: "Lists the wireless gateways registered to your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listWirelessGatewaysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_listWirelessGatewaysCmd).Standalone()

		iotwireless_listWirelessGatewaysCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
		iotwireless_listWirelessGatewaysCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	})
	iotwirelessCmd.AddCommand(iotwireless_listWirelessGatewaysCmd)
}
