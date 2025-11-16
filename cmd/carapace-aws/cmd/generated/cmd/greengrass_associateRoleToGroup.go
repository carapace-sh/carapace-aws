package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_associateRoleToGroupCmd = &cobra.Command{
	Use:   "associate-role-to-group",
	Short: "Associates a role with a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_associateRoleToGroupCmd).Standalone()

	greengrass_associateRoleToGroupCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
	greengrass_associateRoleToGroupCmd.Flags().String("role-arn", "", "The ARN of the role you wish to associate with this group.")
	greengrass_associateRoleToGroupCmd.MarkFlagRequired("group-id")
	greengrass_associateRoleToGroupCmd.MarkFlagRequired("role-arn")
	greengrassCmd.AddCommand(greengrass_associateRoleToGroupCmd)
}
