package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listImageRecipesCmd = &cobra.Command{
	Use:   "list-image-recipes",
	Short: "Returns a list of image recipes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listImageRecipesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listImageRecipesCmd).Standalone()

		imagebuilder_listImageRecipesCmd.Flags().String("filters", "", "Use the following filters to streamline results:")
		imagebuilder_listImageRecipesCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
		imagebuilder_listImageRecipesCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
		imagebuilder_listImageRecipesCmd.Flags().String("owner", "", "You can specify the recipe owner to filter results by that owner.")
	})
	imagebuilderCmd.AddCommand(imagebuilder_listImageRecipesCmd)
}
