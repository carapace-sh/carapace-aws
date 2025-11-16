package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_tagInstanceProfileCmd = &cobra.Command{
	Use:   "tag-instance-profile",
	Short: "Adds one or more tags to an IAM instance profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_tagInstanceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_tagInstanceProfileCmd).Standalone()

		iam_tagInstanceProfileCmd.Flags().String("instance-profile-name", "", "The name of the IAM instance profile to which you want to add tags.")
		iam_tagInstanceProfileCmd.Flags().String("tags", "", "The list of tags that you want to attach to the IAM instance profile.")
		iam_tagInstanceProfileCmd.MarkFlagRequired("instance-profile-name")
		iam_tagInstanceProfileCmd.MarkFlagRequired("tags")
	})
	iamCmd.AddCommand(iam_tagInstanceProfileCmd)
}
