package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Updates the name and/or the path of the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateUserCmd).Standalone()

	iam_updateUserCmd.Flags().String("new-path", "", "New path for the IAM user.")
	iam_updateUserCmd.Flags().String("new-user-name", "", "New name for the user.")
	iam_updateUserCmd.Flags().String("user-name", "", "Name of the user to update.")
	iam_updateUserCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_updateUserCmd)
}
