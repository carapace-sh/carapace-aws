package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_untagUserCmd = &cobra.Command{
	Use:   "untag-user",
	Short: "Removes the specified tags from the user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_untagUserCmd).Standalone()

	iam_untagUserCmd.Flags().String("tag-keys", "", "A list of key names as a simple array of strings.")
	iam_untagUserCmd.Flags().String("user-name", "", "The name of the IAM user from which you want to remove tags.")
	iam_untagUserCmd.MarkFlagRequired("tag-keys")
	iam_untagUserCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_untagUserCmd)
}
