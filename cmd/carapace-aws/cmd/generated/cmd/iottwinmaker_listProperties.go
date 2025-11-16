package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_listPropertiesCmd = &cobra.Command{
	Use:   "list-properties",
	Short: "This API lists the properties of a component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_listPropertiesCmd).Standalone()

	iottwinmaker_listPropertiesCmd.Flags().String("component-name", "", "The name of the component whose properties are returned by the operation.")
	iottwinmaker_listPropertiesCmd.Flags().String("component-path", "", "This string specifies the path to the composite component, starting from the top-level component.")
	iottwinmaker_listPropertiesCmd.Flags().String("entity-id", "", "The ID for the entity whose metadata (component/properties) is returned by the operation.")
	iottwinmaker_listPropertiesCmd.Flags().String("max-results", "", "The maximum number of results returned at one time.")
	iottwinmaker_listPropertiesCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	iottwinmaker_listPropertiesCmd.Flags().String("workspace-id", "", "The workspace ID.")
	iottwinmaker_listPropertiesCmd.MarkFlagRequired("entity-id")
	iottwinmaker_listPropertiesCmd.MarkFlagRequired("workspace-id")
	iottwinmakerCmd.AddCommand(iottwinmaker_listPropertiesCmd)
}
