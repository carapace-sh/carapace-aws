package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listImagesCmd = &cobra.Command{
	Use:   "list-images",
	Short: "Returns the list of images that you have access to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listImagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listImagesCmd).Standalone()

		imagebuilder_listImagesCmd.Flags().Bool("by-name", false, "Requests a list of images with a specific recipe name.")
		imagebuilder_listImagesCmd.Flags().String("filters", "", "Use the following filters to streamline results:")
		imagebuilder_listImagesCmd.Flags().String("include-deprecated", "", "Includes deprecated images in the response list.")
		imagebuilder_listImagesCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
		imagebuilder_listImagesCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
		imagebuilder_listImagesCmd.Flags().Bool("no-by-name", false, "Requests a list of images with a specific recipe name.")
		imagebuilder_listImagesCmd.Flags().String("owner", "", "The owner defines which images you want to list.")
		imagebuilder_listImagesCmd.Flag("no-by-name").Hidden = true
	})
	imagebuilderCmd.AddCommand(imagebuilder_listImagesCmd)
}
