package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeRoleCustomPermissionCmd = &cobra.Command{
	Use:   "describe-role-custom-permission",
	Short: "Describes all custom permissions that are mapped to a role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeRoleCustomPermissionCmd).Standalone()

	quicksight_describeRoleCustomPermissionCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that you want to create a group in.")
	quicksight_describeRoleCustomPermissionCmd.Flags().String("namespace", "", "The namespace that contains the role.")
	quicksight_describeRoleCustomPermissionCmd.Flags().String("role", "", "The name of the role whose permissions you want described.")
	quicksight_describeRoleCustomPermissionCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeRoleCustomPermissionCmd.MarkFlagRequired("namespace")
	quicksight_describeRoleCustomPermissionCmd.MarkFlagRequired("role")
	quicksightCmd.AddCommand(quicksight_describeRoleCustomPermissionCmd)
}
