package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_createHarvestJobCmd = &cobra.Command{
	Use:   "create-harvest-job",
	Short: "Creates a new HarvestJob record.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_createHarvestJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackage_createHarvestJobCmd).Standalone()

		mediapackage_createHarvestJobCmd.Flags().String("end-time", "", "The end of the time-window which will be harvested")
		mediapackage_createHarvestJobCmd.Flags().String("id", "", "The ID of the HarvestJob.")
		mediapackage_createHarvestJobCmd.Flags().String("origin-endpoint-id", "", "The ID of the OriginEndpoint that the HarvestJob will harvest from.")
		mediapackage_createHarvestJobCmd.Flags().String("s3-destination", "", "")
		mediapackage_createHarvestJobCmd.Flags().String("start-time", "", "The start of the time-window which will be harvested")
		mediapackage_createHarvestJobCmd.MarkFlagRequired("end-time")
		mediapackage_createHarvestJobCmd.MarkFlagRequired("id")
		mediapackage_createHarvestJobCmd.MarkFlagRequired("origin-endpoint-id")
		mediapackage_createHarvestJobCmd.MarkFlagRequired("s3-destination")
		mediapackage_createHarvestJobCmd.MarkFlagRequired("start-time")
	})
	mediapackageCmd.AddCommand(mediapackage_createHarvestJobCmd)
}
