package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_tagUserCmd = &cobra.Command{
	Use:   "tag-user",
	Short: "Adds one or more tags to an IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_tagUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_tagUserCmd).Standalone()

		iam_tagUserCmd.Flags().String("tags", "", "The list of tags that you want to attach to the IAM user.")
		iam_tagUserCmd.Flags().String("user-name", "", "The name of the IAM user to which you want to add tags.")
		iam_tagUserCmd.MarkFlagRequired("tags")
		iam_tagUserCmd.MarkFlagRequired("user-name")
	})
	iamCmd.AddCommand(iam_tagUserCmd)
}
