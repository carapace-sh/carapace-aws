package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createRoleCmd = &cobra.Command{
	Use:   "create-role",
	Short: "Creates a new role for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createRoleCmd).Standalone()

	iam_createRoleCmd.Flags().String("assume-role-policy-document", "", "The trust relationship policy document that grants an entity permission to assume the role.")
	iam_createRoleCmd.Flags().String("description", "", "A description of the role.")
	iam_createRoleCmd.Flags().String("max-session-duration", "", "The maximum session duration (in seconds) that you want to set for the specified role.")
	iam_createRoleCmd.Flags().String("path", "", "The path to the role.")
	iam_createRoleCmd.Flags().String("permissions-boundary", "", "The ARN of the managed policy that is used to set the permissions boundary for the role.")
	iam_createRoleCmd.Flags().String("role-name", "", "The name of the role to create.")
	iam_createRoleCmd.Flags().String("tags", "", "A list of tags that you want to attach to the new role.")
	iam_createRoleCmd.MarkFlagRequired("assume-role-policy-document")
	iam_createRoleCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_createRoleCmd)
}
