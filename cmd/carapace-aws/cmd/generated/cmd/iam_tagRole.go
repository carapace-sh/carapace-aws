package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_tagRoleCmd = &cobra.Command{
	Use:   "tag-role",
	Short: "Adds one or more tags to an IAM role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_tagRoleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_tagRoleCmd).Standalone()

		iam_tagRoleCmd.Flags().String("role-name", "", "The name of the IAM role to which you want to add tags.")
		iam_tagRoleCmd.Flags().String("tags", "", "The list of tags that you want to attach to the IAM role.")
		iam_tagRoleCmd.MarkFlagRequired("role-name")
		iam_tagRoleCmd.MarkFlagRequired("tags")
	})
	iamCmd.AddCommand(iam_tagRoleCmd)
}
