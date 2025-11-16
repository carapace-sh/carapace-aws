package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteMultiRegionEndpointCmd = &cobra.Command{
	Use:   "delete-multi-region-endpoint",
	Short: "Deletes a multi-region endpoint (global-endpoint).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteMultiRegionEndpointCmd).Standalone()

	sesv2_deleteMultiRegionEndpointCmd.Flags().String("endpoint-name", "", "The name of the multi-region endpoint (global-endpoint) to be deleted.")
	sesv2_deleteMultiRegionEndpointCmd.MarkFlagRequired("endpoint-name")
	sesv2Cmd.AddCommand(sesv2_deleteMultiRegionEndpointCmd)
}
