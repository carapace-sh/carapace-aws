package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listContainerRecipesCmd = &cobra.Command{
	Use:   "list-container-recipes",
	Short: "Returns a list of container recipes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listContainerRecipesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listContainerRecipesCmd).Standalone()

		imagebuilder_listContainerRecipesCmd.Flags().String("filters", "", "Use the following filters to streamline results:")
		imagebuilder_listContainerRecipesCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
		imagebuilder_listContainerRecipesCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
		imagebuilder_listContainerRecipesCmd.Flags().String("owner", "", "Returns container recipes belonging to the specified owner, that have been shared with you.")
	})
	imagebuilderCmd.AddCommand(imagebuilder_listContainerRecipesCmd)
}
