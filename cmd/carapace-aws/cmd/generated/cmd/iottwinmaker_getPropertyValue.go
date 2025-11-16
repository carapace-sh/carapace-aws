package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_getPropertyValueCmd = &cobra.Command{
	Use:   "get-property-value",
	Short: "Gets the property values for a component, component type, entity, or workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_getPropertyValueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_getPropertyValueCmd).Standalone()

		iottwinmaker_getPropertyValueCmd.Flags().String("component-name", "", "The name of the component whose property values the operation returns.")
		iottwinmaker_getPropertyValueCmd.Flags().String("component-path", "", "This string specifies the path to the composite component, starting from the top-level component.")
		iottwinmaker_getPropertyValueCmd.Flags().String("component-type-id", "", "The ID of the component type whose property values the operation returns.")
		iottwinmaker_getPropertyValueCmd.Flags().String("entity-id", "", "The ID of the entity whose property values the operation returns.")
		iottwinmaker_getPropertyValueCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iottwinmaker_getPropertyValueCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
		iottwinmaker_getPropertyValueCmd.Flags().String("property-group-name", "", "The property group name.")
		iottwinmaker_getPropertyValueCmd.Flags().String("selected-properties", "", "The properties whose values the operation returns.")
		iottwinmaker_getPropertyValueCmd.Flags().String("tabular-conditions", "", "The tabular conditions.")
		iottwinmaker_getPropertyValueCmd.Flags().String("workspace-id", "", "The ID of the workspace whose values the operation returns.")
		iottwinmaker_getPropertyValueCmd.MarkFlagRequired("selected-properties")
		iottwinmaker_getPropertyValueCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_getPropertyValueCmd)
}
