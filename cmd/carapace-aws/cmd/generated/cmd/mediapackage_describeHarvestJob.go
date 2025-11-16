package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_describeHarvestJobCmd = &cobra.Command{
	Use:   "describe-harvest-job",
	Short: "Gets details about an existing HarvestJob.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_describeHarvestJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackage_describeHarvestJobCmd).Standalone()

		mediapackage_describeHarvestJobCmd.Flags().String("id", "", "The ID of the HarvestJob.")
		mediapackage_describeHarvestJobCmd.MarkFlagRequired("id")
	})
	mediapackageCmd.AddCommand(mediapackage_describeHarvestJobCmd)
}
