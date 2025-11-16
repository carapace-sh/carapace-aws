package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listImageScanFindingAggregationsCmd = &cobra.Command{
	Use:   "list-image-scan-finding-aggregations",
	Short: "Returns a list of image scan aggregations for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listImageScanFindingAggregationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listImageScanFindingAggregationsCmd).Standalone()

		imagebuilder_listImageScanFindingAggregationsCmd.Flags().String("filter", "", "")
		imagebuilder_listImageScanFindingAggregationsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	})
	imagebuilderCmd.AddCommand(imagebuilder_listImageScanFindingAggregationsCmd)
}
