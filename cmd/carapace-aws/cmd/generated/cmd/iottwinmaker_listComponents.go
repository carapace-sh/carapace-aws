package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_listComponentsCmd = &cobra.Command{
	Use:   "list-components",
	Short: "This API lists the components of an entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_listComponentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_listComponentsCmd).Standalone()

		iottwinmaker_listComponentsCmd.Flags().String("component-path", "", "This string specifies the path to the composite component, starting from the top-level component.")
		iottwinmaker_listComponentsCmd.Flags().String("entity-id", "", "The ID for the entity whose metadata (component/properties) is returned by the operation.")
		iottwinmaker_listComponentsCmd.Flags().String("max-results", "", "The maximum number of results returned at one time.")
		iottwinmaker_listComponentsCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
		iottwinmaker_listComponentsCmd.Flags().String("workspace-id", "", "The workspace ID.")
		iottwinmaker_listComponentsCmd.MarkFlagRequired("entity-id")
		iottwinmaker_listComponentsCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_listComponentsCmd)
}
