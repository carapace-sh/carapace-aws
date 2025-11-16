package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_deleteComponentTypeCmd = &cobra.Command{
	Use:   "delete-component-type",
	Short: "Deletes a component type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_deleteComponentTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_deleteComponentTypeCmd).Standalone()

		iottwinmaker_deleteComponentTypeCmd.Flags().String("component-type-id", "", "The ID of the component type to delete.")
		iottwinmaker_deleteComponentTypeCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the component type.")
		iottwinmaker_deleteComponentTypeCmd.MarkFlagRequired("component-type-id")
		iottwinmaker_deleteComponentTypeCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_deleteComponentTypeCmd)
}
