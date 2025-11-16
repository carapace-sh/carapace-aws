package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listImagePackagesCmd = &cobra.Command{
	Use:   "list-image-packages",
	Short: "List the Packages that are associated with an Image Build Version, as determined by Amazon Web Services Systems Manager Inventory at build time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listImagePackagesCmd).Standalone()

	imagebuilder_listImagePackagesCmd.Flags().String("image-build-version-arn", "", "Filter results for the ListImagePackages request by the Image Build Version ARN")
	imagebuilder_listImagePackagesCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
	imagebuilder_listImagePackagesCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	imagebuilder_listImagePackagesCmd.MarkFlagRequired("image-build-version-arn")
	imagebuilderCmd.AddCommand(imagebuilder_listImagePackagesCmd)
}
