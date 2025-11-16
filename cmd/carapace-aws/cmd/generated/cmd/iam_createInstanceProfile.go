package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createInstanceProfileCmd = &cobra.Command{
	Use:   "create-instance-profile",
	Short: "Creates a new instance profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createInstanceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_createInstanceProfileCmd).Standalone()

		iam_createInstanceProfileCmd.Flags().String("instance-profile-name", "", "The name of the instance profile to create.")
		iam_createInstanceProfileCmd.Flags().String("path", "", "The path to the instance profile.")
		iam_createInstanceProfileCmd.Flags().String("tags", "", "A list of tags that you want to attach to the newly created IAM instance profile.")
		iam_createInstanceProfileCmd.MarkFlagRequired("instance-profile-name")
	})
	iamCmd.AddCommand(iam_createInstanceProfileCmd)
}
