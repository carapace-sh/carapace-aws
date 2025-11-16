package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateRoleCmd = &cobra.Command{
	Use:   "update-role",
	Short: "Updates the description or maximum session duration setting of a role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateRoleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_updateRoleCmd).Standalone()

		iam_updateRoleCmd.Flags().String("description", "", "The new description that you want to apply to the specified role.")
		iam_updateRoleCmd.Flags().String("max-session-duration", "", "The maximum session duration (in seconds) that you want to set for the specified role.")
		iam_updateRoleCmd.Flags().String("role-name", "", "The name of the role that you want to modify.")
		iam_updateRoleCmd.MarkFlagRequired("role-name")
	})
	iamCmd.AddCommand(iam_updateRoleCmd)
}
