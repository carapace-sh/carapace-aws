package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listImagePipelinesCmd = &cobra.Command{
	Use:   "list-image-pipelines",
	Short: "Returns a list of image pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listImagePipelinesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listImagePipelinesCmd).Standalone()

		imagebuilder_listImagePipelinesCmd.Flags().String("filters", "", "Use the following filters to streamline results:")
		imagebuilder_listImagePipelinesCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
		imagebuilder_listImagePipelinesCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	})
	imagebuilderCmd.AddCommand(imagebuilder_listImagePipelinesCmd)
}
