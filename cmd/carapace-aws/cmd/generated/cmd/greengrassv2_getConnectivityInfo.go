package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_getConnectivityInfoCmd = &cobra.Command{
	Use:   "get-connectivity-info",
	Short: "Retrieves connectivity information for a Greengrass core device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_getConnectivityInfoCmd).Standalone()

	greengrassv2_getConnectivityInfoCmd.Flags().String("thing-name", "", "The name of the core device.")
	greengrassv2_getConnectivityInfoCmd.MarkFlagRequired("thing-name")
	greengrassv2Cmd.AddCommand(greengrassv2_getConnectivityInfoCmd)
}
