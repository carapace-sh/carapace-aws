package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listImageScanFindingsCmd = &cobra.Command{
	Use:   "list-image-scan-findings",
	Short: "Returns a list of image scan findings for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listImageScanFindingsCmd).Standalone()

	imagebuilder_listImageScanFindingsCmd.Flags().String("filters", "", "An array of name value pairs that you can use to filter your results.")
	imagebuilder_listImageScanFindingsCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
	imagebuilder_listImageScanFindingsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	imagebuilderCmd.AddCommand(imagebuilder_listImageScanFindingsCmd)
}
