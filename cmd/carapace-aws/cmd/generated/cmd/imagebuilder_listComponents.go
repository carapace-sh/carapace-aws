package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listComponentsCmd = &cobra.Command{
	Use:   "list-components",
	Short: "Returns the list of components that can be filtered by name, or by using the listed `filters` to streamline results.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listComponentsCmd).Standalone()

	imagebuilder_listComponentsCmd.Flags().Bool("by-name", false, "Returns the list of components for the specified name.")
	imagebuilder_listComponentsCmd.Flags().String("filters", "", "Use the following filters to streamline results:")
	imagebuilder_listComponentsCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
	imagebuilder_listComponentsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	imagebuilder_listComponentsCmd.Flags().Bool("no-by-name", false, "Returns the list of components for the specified name.")
	imagebuilder_listComponentsCmd.Flags().String("owner", "", "Filters results based on the type of owner for the component.")
	imagebuilder_listComponentsCmd.Flag("no-by-name").Hidden = true
	imagebuilderCmd.AddCommand(imagebuilder_listComponentsCmd)
}
