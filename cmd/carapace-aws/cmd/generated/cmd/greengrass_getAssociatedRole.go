package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getAssociatedRoleCmd = &cobra.Command{
	Use:   "get-associated-role",
	Short: "Retrieves the role associated with a particular group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getAssociatedRoleCmd).Standalone()

	greengrass_getAssociatedRoleCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
	greengrass_getAssociatedRoleCmd.MarkFlagRequired("group-id")
	greengrassCmd.AddCommand(greengrass_getAssociatedRoleCmd)
}
