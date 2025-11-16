package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_cancelHarvestJobCmd = &cobra.Command{
	Use:   "cancel-harvest-job",
	Short: "Cancels an in-progress harvest job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_cancelHarvestJobCmd).Standalone()

	mediapackagev2_cancelHarvestJobCmd.Flags().String("channel-group-name", "", "The name of the channel group containing the channel from which the harvest job is running.")
	mediapackagev2_cancelHarvestJobCmd.Flags().String("channel-name", "", "The name of the channel from which the harvest job is running.")
	mediapackagev2_cancelHarvestJobCmd.Flags().String("etag", "", "The current Entity Tag (ETag) associated with the harvest job.")
	mediapackagev2_cancelHarvestJobCmd.Flags().String("harvest-job-name", "", "The name of the harvest job to cancel.")
	mediapackagev2_cancelHarvestJobCmd.Flags().String("origin-endpoint-name", "", "The name of the origin endpoint that the harvest job is harvesting from.")
	mediapackagev2_cancelHarvestJobCmd.MarkFlagRequired("channel-group-name")
	mediapackagev2_cancelHarvestJobCmd.MarkFlagRequired("channel-name")
	mediapackagev2_cancelHarvestJobCmd.MarkFlagRequired("harvest-job-name")
	mediapackagev2_cancelHarvestJobCmd.MarkFlagRequired("origin-endpoint-name")
	mediapackagev2Cmd.AddCommand(mediapackagev2_cancelHarvestJobCmd)
}
