package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listWirelessGatewayTaskDefinitionsCmd = &cobra.Command{
	Use:   "list-wireless-gateway-task-definitions",
	Short: "List the wireless gateway tasks definitions registered to your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listWirelessGatewayTaskDefinitionsCmd).Standalone()

	iotwireless_listWirelessGatewayTaskDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
	iotwireless_listWirelessGatewayTaskDefinitionsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	iotwireless_listWirelessGatewayTaskDefinitionsCmd.Flags().String("task-definition-type", "", "A filter to list only the wireless gateway task definitions that use this task definition type.")
	iotwirelessCmd.AddCommand(iotwireless_listWirelessGatewayTaskDefinitionsCmd)
}
