package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_listOriginEndpointsCmd = &cobra.Command{
	Use:   "list-origin-endpoints",
	Short: "Returns a collection of OriginEndpoint records.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_listOriginEndpointsCmd).Standalone()

	mediapackage_listOriginEndpointsCmd.Flags().String("channel-id", "", "When specified, the request will return only OriginEndpoints associated with the given Channel ID.")
	mediapackage_listOriginEndpointsCmd.Flags().String("max-results", "", "The upper bound on the number of records to return.")
	mediapackage_listOriginEndpointsCmd.Flags().String("next-token", "", "A token used to resume pagination from the end of a previous request.")
	mediapackageCmd.AddCommand(mediapackage_listOriginEndpointsCmd)
}
