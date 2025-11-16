package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_untagInstanceProfileCmd = &cobra.Command{
	Use:   "untag-instance-profile",
	Short: "Removes the specified tags from the IAM instance profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_untagInstanceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_untagInstanceProfileCmd).Standalone()

		iam_untagInstanceProfileCmd.Flags().String("instance-profile-name", "", "The name of the IAM instance profile from which you want to remove tags.")
		iam_untagInstanceProfileCmd.Flags().String("tag-keys", "", "A list of key names as a simple array of strings.")
		iam_untagInstanceProfileCmd.MarkFlagRequired("instance-profile-name")
		iam_untagInstanceProfileCmd.MarkFlagRequired("tag-keys")
	})
	iamCmd.AddCommand(iam_untagInstanceProfileCmd)
}
