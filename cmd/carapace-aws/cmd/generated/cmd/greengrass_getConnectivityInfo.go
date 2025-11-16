package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getConnectivityInfoCmd = &cobra.Command{
	Use:   "get-connectivity-info",
	Short: "Retrieves the connectivity information for a core.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getConnectivityInfoCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_getConnectivityInfoCmd).Standalone()

		greengrass_getConnectivityInfoCmd.Flags().String("thing-name", "", "The thing name.")
		greengrass_getConnectivityInfoCmd.MarkFlagRequired("thing-name")
	})
	greengrassCmd.AddCommand(greengrass_getConnectivityInfoCmd)
}
