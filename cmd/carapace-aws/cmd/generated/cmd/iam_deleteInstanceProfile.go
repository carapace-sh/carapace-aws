package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteInstanceProfileCmd = &cobra.Command{
	Use:   "delete-instance-profile",
	Short: "Deletes the specified instance profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteInstanceProfileCmd).Standalone()

	iam_deleteInstanceProfileCmd.Flags().String("instance-profile-name", "", "The name of the instance profile to delete.")
	iam_deleteInstanceProfileCmd.MarkFlagRequired("instance-profile-name")
	iamCmd.AddCommand(iam_deleteInstanceProfileCmd)
}
