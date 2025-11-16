package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listImageBuildVersionsCmd = &cobra.Command{
	Use:   "list-image-build-versions",
	Short: "Returns a list of image build versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listImageBuildVersionsCmd).Standalone()

	imagebuilder_listImageBuildVersionsCmd.Flags().String("filters", "", "Use the following filters to streamline results:")
	imagebuilder_listImageBuildVersionsCmd.Flags().String("image-version-arn", "", "The Amazon Resource Name (ARN) of the image whose build versions you want to retrieve.")
	imagebuilder_listImageBuildVersionsCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
	imagebuilder_listImageBuildVersionsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	imagebuilderCmd.AddCommand(imagebuilder_listImageBuildVersionsCmd)
}
