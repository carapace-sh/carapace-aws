package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Deletes a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_deleteGroupCmd).Standalone()

	greengrass_deleteGroupCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
	greengrass_deleteGroupCmd.MarkFlagRequired("group-id")
	greengrassCmd.AddCommand(greengrass_deleteGroupCmd)
}
