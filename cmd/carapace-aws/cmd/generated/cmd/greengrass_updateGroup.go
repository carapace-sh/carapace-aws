package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_updateGroupCmd = &cobra.Command{
	Use:   "update-group",
	Short: "Updates a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_updateGroupCmd).Standalone()

	greengrass_updateGroupCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
	greengrass_updateGroupCmd.Flags().String("name", "", "The name of the definition.")
	greengrass_updateGroupCmd.MarkFlagRequired("group-id")
	greengrassCmd.AddCommand(greengrass_updateGroupCmd)
}
