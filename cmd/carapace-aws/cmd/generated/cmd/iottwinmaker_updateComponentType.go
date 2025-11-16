package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_updateComponentTypeCmd = &cobra.Command{
	Use:   "update-component-type",
	Short: "Updates information in a component type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_updateComponentTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_updateComponentTypeCmd).Standalone()

		iottwinmaker_updateComponentTypeCmd.Flags().String("component-type-id", "", "The ID of the component type.")
		iottwinmaker_updateComponentTypeCmd.Flags().String("component-type-name", "", "The component type name.")
		iottwinmaker_updateComponentTypeCmd.Flags().String("composite-component-types", "", "This is an object that maps strings to `compositeComponentTypes` of the `componentType`.")
		iottwinmaker_updateComponentTypeCmd.Flags().String("description", "", "The description of the component type.")
		iottwinmaker_updateComponentTypeCmd.Flags().String("extends-from", "", "Specifies the component type that this component type extends.")
		iottwinmaker_updateComponentTypeCmd.Flags().String("functions", "", "An object that maps strings to the functions in the component type.")
		iottwinmaker_updateComponentTypeCmd.Flags().Bool("is-singleton", false, "A Boolean value that specifies whether an entity can have more than one component of this type.")
		iottwinmaker_updateComponentTypeCmd.Flags().Bool("no-is-singleton", false, "A Boolean value that specifies whether an entity can have more than one component of this type.")
		iottwinmaker_updateComponentTypeCmd.Flags().String("property-definitions", "", "An object that maps strings to the property definitions in the component type.")
		iottwinmaker_updateComponentTypeCmd.Flags().String("property-groups", "", "The property groups.")
		iottwinmaker_updateComponentTypeCmd.Flags().String("workspace-id", "", "The ID of the workspace.")
		iottwinmaker_updateComponentTypeCmd.MarkFlagRequired("component-type-id")
		iottwinmaker_updateComponentTypeCmd.Flag("no-is-singleton").Hidden = true
		iottwinmaker_updateComponentTypeCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_updateComponentTypeCmd)
}
