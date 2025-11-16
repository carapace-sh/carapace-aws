package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_describeOriginEndpointCmd = &cobra.Command{
	Use:   "describe-origin-endpoint",
	Short: "Gets details about an existing OriginEndpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_describeOriginEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackage_describeOriginEndpointCmd).Standalone()

		mediapackage_describeOriginEndpointCmd.Flags().String("id", "", "The ID of the OriginEndpoint.")
		mediapackage_describeOriginEndpointCmd.MarkFlagRequired("id")
	})
	mediapackageCmd.AddCommand(mediapackage_describeOriginEndpointCmd)
}
