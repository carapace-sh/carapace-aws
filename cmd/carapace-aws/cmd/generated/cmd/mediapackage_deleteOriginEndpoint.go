package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_deleteOriginEndpointCmd = &cobra.Command{
	Use:   "delete-origin-endpoint",
	Short: "Deletes an existing OriginEndpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_deleteOriginEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackage_deleteOriginEndpointCmd).Standalone()

		mediapackage_deleteOriginEndpointCmd.Flags().String("id", "", "The ID of the OriginEndpoint to delete.")
		mediapackage_deleteOriginEndpointCmd.MarkFlagRequired("id")
	})
	mediapackageCmd.AddCommand(mediapackage_deleteOriginEndpointCmd)
}
