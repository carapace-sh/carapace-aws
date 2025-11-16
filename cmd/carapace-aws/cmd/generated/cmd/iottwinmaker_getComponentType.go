package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_getComponentTypeCmd = &cobra.Command{
	Use:   "get-component-type",
	Short: "Retrieves information about a component type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_getComponentTypeCmd).Standalone()

	iottwinmaker_getComponentTypeCmd.Flags().String("component-type-id", "", "The ID of the component type.")
	iottwinmaker_getComponentTypeCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the component type.")
	iottwinmaker_getComponentTypeCmd.MarkFlagRequired("component-type-id")
	iottwinmaker_getComponentTypeCmd.MarkFlagRequired("workspace-id")
	iottwinmakerCmd.AddCommand(iottwinmaker_getComponentTypeCmd)
}
