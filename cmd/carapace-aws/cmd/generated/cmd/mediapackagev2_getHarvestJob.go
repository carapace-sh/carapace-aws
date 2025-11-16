package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_getHarvestJobCmd = &cobra.Command{
	Use:   "get-harvest-job",
	Short: "Retrieves the details of a specific harvest job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_getHarvestJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_getHarvestJobCmd).Standalone()

		mediapackagev2_getHarvestJobCmd.Flags().String("channel-group-name", "", "The name of the channel group containing the channel associated with the harvest job.")
		mediapackagev2_getHarvestJobCmd.Flags().String("channel-name", "", "The name of the channel associated with the harvest job.")
		mediapackagev2_getHarvestJobCmd.Flags().String("harvest-job-name", "", "The name of the harvest job to retrieve.")
		mediapackagev2_getHarvestJobCmd.Flags().String("origin-endpoint-name", "", "The name of the origin endpoint associated with the harvest job.")
		mediapackagev2_getHarvestJobCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_getHarvestJobCmd.MarkFlagRequired("channel-name")
		mediapackagev2_getHarvestJobCmd.MarkFlagRequired("harvest-job-name")
		mediapackagev2_getHarvestJobCmd.MarkFlagRequired("origin-endpoint-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_getHarvestJobCmd)
}
