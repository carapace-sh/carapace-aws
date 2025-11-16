package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getMultiRegionEndpointCmd = &cobra.Command{
	Use:   "get-multi-region-endpoint",
	Short: "Displays the multi-region endpoint (global-endpoint) configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getMultiRegionEndpointCmd).Standalone()

	sesv2_getMultiRegionEndpointCmd.Flags().String("endpoint-name", "", "The name of the multi-region endpoint (global-endpoint).")
	sesv2_getMultiRegionEndpointCmd.MarkFlagRequired("endpoint-name")
	sesv2Cmd.AddCommand(sesv2_getMultiRegionEndpointCmd)
}
