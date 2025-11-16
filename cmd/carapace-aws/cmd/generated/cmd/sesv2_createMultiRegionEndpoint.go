package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createMultiRegionEndpointCmd = &cobra.Command{
	Use:   "create-multi-region-endpoint",
	Short: "Creates a multi-region endpoint (global-endpoint).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createMultiRegionEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_createMultiRegionEndpointCmd).Standalone()

		sesv2_createMultiRegionEndpointCmd.Flags().String("details", "", "Contains details of a multi-region endpoint (global-endpoint) being created.")
		sesv2_createMultiRegionEndpointCmd.Flags().String("endpoint-name", "", "The name of the multi-region endpoint (global-endpoint).")
		sesv2_createMultiRegionEndpointCmd.Flags().String("tags", "", "An array of objects that define the tags (keys and values) to associate with the multi-region endpoint (global-endpoint).")
		sesv2_createMultiRegionEndpointCmd.MarkFlagRequired("details")
		sesv2_createMultiRegionEndpointCmd.MarkFlagRequired("endpoint-name")
	})
	sesv2Cmd.AddCommand(sesv2_createMultiRegionEndpointCmd)
}
