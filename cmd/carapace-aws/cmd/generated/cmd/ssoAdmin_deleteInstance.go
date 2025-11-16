package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_deleteInstanceCmd = &cobra.Command{
	Use:   "delete-instance",
	Short: "Deletes the instance of IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_deleteInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_deleteInstanceCmd).Standalone()

		ssoAdmin_deleteInstanceCmd.Flags().String("instance-arn", "", "The ARN of the instance of IAM Identity Center under which the operation will run.")
		ssoAdmin_deleteInstanceCmd.MarkFlagRequired("instance-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_deleteInstanceCmd)
}
