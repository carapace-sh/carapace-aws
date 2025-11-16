package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a new IAM user for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_createUserCmd).Standalone()

		iam_createUserCmd.Flags().String("path", "", "The path for the user name.")
		iam_createUserCmd.Flags().String("permissions-boundary", "", "The ARN of the managed policy that is used to set the permissions boundary for the user.")
		iam_createUserCmd.Flags().String("tags", "", "A list of tags that you want to attach to the new user.")
		iam_createUserCmd.Flags().String("user-name", "", "The name of the user to create.")
		iam_createUserCmd.MarkFlagRequired("user-name")
	})
	iamCmd.AddCommand(iam_createUserCmd)
}
