package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_listOriginEndpointsCmd = &cobra.Command{
	Use:   "list-origin-endpoints",
	Short: "Retrieves all origin endpoints in a specific channel that are configured in AWS Elemental MediaPackage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_listOriginEndpointsCmd).Standalone()

	mediapackagev2_listOriginEndpointsCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
	mediapackagev2_listOriginEndpointsCmd.Flags().String("channel-name", "", "The name that describes the channel.")
	mediapackagev2_listOriginEndpointsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	mediapackagev2_listOriginEndpointsCmd.Flags().String("next-token", "", "The pagination token from the GET list request.")
	mediapackagev2_listOriginEndpointsCmd.MarkFlagRequired("channel-group-name")
	mediapackagev2_listOriginEndpointsCmd.MarkFlagRequired("channel-name")
	mediapackagev2Cmd.AddCommand(mediapackagev2_listOriginEndpointsCmd)
}
