package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_disassociateRoleFromGroupCmd = &cobra.Command{
	Use:   "disassociate-role-from-group",
	Short: "Disassociates the role from a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_disassociateRoleFromGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_disassociateRoleFromGroupCmd).Standalone()

		greengrass_disassociateRoleFromGroupCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
		greengrass_disassociateRoleFromGroupCmd.MarkFlagRequired("group-id")
	})
	greengrassCmd.AddCommand(greengrass_disassociateRoleFromGroupCmd)
}
