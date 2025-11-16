package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listPositionConfigurationsCmd = &cobra.Command{
	Use:   "list-position-configurations",
	Short: "List position configurations for a given resource, such as positioning solvers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listPositionConfigurationsCmd).Standalone()

	iotwireless_listPositionConfigurationsCmd.Flags().String("max-results", "", "")
	iotwireless_listPositionConfigurationsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	iotwireless_listPositionConfigurationsCmd.Flags().String("resource-type", "", "Resource type for which position configurations are listed.")
	iotwirelessCmd.AddCommand(iotwireless_listPositionConfigurationsCmd)
}
