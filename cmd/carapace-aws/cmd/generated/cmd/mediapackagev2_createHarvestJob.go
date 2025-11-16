package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_createHarvestJobCmd = &cobra.Command{
	Use:   "create-harvest-job",
	Short: "Creates a new harvest job to export content from a MediaPackage v2 channel to an S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_createHarvestJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_createHarvestJobCmd).Standalone()

		mediapackagev2_createHarvestJobCmd.Flags().String("channel-group-name", "", "The name of the channel group containing the channel from which to harvest content.")
		mediapackagev2_createHarvestJobCmd.Flags().String("channel-name", "", "The name of the channel from which to harvest content.")
		mediapackagev2_createHarvestJobCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		mediapackagev2_createHarvestJobCmd.Flags().String("description", "", "An optional description for the harvest job.")
		mediapackagev2_createHarvestJobCmd.Flags().String("destination", "", "The S3 destination where the harvested content will be placed.")
		mediapackagev2_createHarvestJobCmd.Flags().String("harvest-job-name", "", "A name for the harvest job.")
		mediapackagev2_createHarvestJobCmd.Flags().String("harvested-manifests", "", "A list of manifests to be harvested.")
		mediapackagev2_createHarvestJobCmd.Flags().String("origin-endpoint-name", "", "The name of the origin endpoint from which to harvest content.")
		mediapackagev2_createHarvestJobCmd.Flags().String("schedule-configuration", "", "The configuration for when the harvest job should run, including start and end times.")
		mediapackagev2_createHarvestJobCmd.Flags().String("tags", "", "A collection of tags associated with the harvest job.")
		mediapackagev2_createHarvestJobCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_createHarvestJobCmd.MarkFlagRequired("channel-name")
		mediapackagev2_createHarvestJobCmd.MarkFlagRequired("destination")
		mediapackagev2_createHarvestJobCmd.MarkFlagRequired("harvested-manifests")
		mediapackagev2_createHarvestJobCmd.MarkFlagRequired("origin-endpoint-name")
		mediapackagev2_createHarvestJobCmd.MarkFlagRequired("schedule-configuration")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_createHarvestJobCmd)
}
