package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_updateConnectivityInfoCmd = &cobra.Command{
	Use:   "update-connectivity-info",
	Short: "Updates the connectivity information for the core.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_updateConnectivityInfoCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_updateConnectivityInfoCmd).Standalone()

		greengrass_updateConnectivityInfoCmd.Flags().String("connectivity-info", "", "A list of connectivity info.")
		greengrass_updateConnectivityInfoCmd.Flags().String("thing-name", "", "The thing name.")
		greengrass_updateConnectivityInfoCmd.MarkFlagRequired("thing-name")
	})
	greengrassCmd.AddCommand(greengrass_updateConnectivityInfoCmd)
}
