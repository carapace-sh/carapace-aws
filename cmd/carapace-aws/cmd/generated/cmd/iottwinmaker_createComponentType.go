package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_createComponentTypeCmd = &cobra.Command{
	Use:   "create-component-type",
	Short: "Creates a component type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_createComponentTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_createComponentTypeCmd).Standalone()

		iottwinmaker_createComponentTypeCmd.Flags().String("component-type-id", "", "The ID of the component type.")
		iottwinmaker_createComponentTypeCmd.Flags().String("component-type-name", "", "A friendly name for the component type.")
		iottwinmaker_createComponentTypeCmd.Flags().String("composite-component-types", "", "This is an object that maps strings to `compositeComponentTypes` of the `componentType`.")
		iottwinmaker_createComponentTypeCmd.Flags().String("description", "", "The description of the component type.")
		iottwinmaker_createComponentTypeCmd.Flags().String("extends-from", "", "Specifies the parent component type to extend.")
		iottwinmaker_createComponentTypeCmd.Flags().String("functions", "", "An object that maps strings to the functions in the component type.")
		iottwinmaker_createComponentTypeCmd.Flags().Bool("is-singleton", false, "A Boolean value that specifies whether an entity can have more than one component of this type.")
		iottwinmaker_createComponentTypeCmd.Flags().Bool("no-is-singleton", false, "A Boolean value that specifies whether an entity can have more than one component of this type.")
		iottwinmaker_createComponentTypeCmd.Flags().String("property-definitions", "", "An object that maps strings to the property definitions in the component type.")
		iottwinmaker_createComponentTypeCmd.Flags().String("property-groups", "", "")
		iottwinmaker_createComponentTypeCmd.Flags().String("tags", "", "Metadata that you can use to manage the component type.")
		iottwinmaker_createComponentTypeCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the component type.")
		iottwinmaker_createComponentTypeCmd.MarkFlagRequired("component-type-id")
		iottwinmaker_createComponentTypeCmd.Flag("no-is-singleton").Hidden = true
		iottwinmaker_createComponentTypeCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_createComponentTypeCmd)
}
