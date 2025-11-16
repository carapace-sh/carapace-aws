package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_updateConnectivityInfoCmd = &cobra.Command{
	Use:   "update-connectivity-info",
	Short: "Updates connectivity information for a Greengrass core device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_updateConnectivityInfoCmd).Standalone()

	greengrassv2_updateConnectivityInfoCmd.Flags().String("connectivity-info", "", "The connectivity information for the core device.")
	greengrassv2_updateConnectivityInfoCmd.Flags().String("thing-name", "", "The name of the core device.")
	greengrassv2_updateConnectivityInfoCmd.MarkFlagRequired("connectivity-info")
	greengrassv2_updateConnectivityInfoCmd.MarkFlagRequired("thing-name")
	greengrassv2Cmd.AddCommand(greengrassv2_updateConnectivityInfoCmd)
}
