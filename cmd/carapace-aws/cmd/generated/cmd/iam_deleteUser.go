package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteUserCmd).Standalone()

	iam_deleteUserCmd.Flags().String("user-name", "", "The name of the user to delete.")
	iam_deleteUserCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_deleteUserCmd)
}
