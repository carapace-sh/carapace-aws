package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_describeInstanceCmd = &cobra.Command{
	Use:   "describe-instance",
	Short: "Returns the details of an instance of IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_describeInstanceCmd).Standalone()

	ssoAdmin_describeInstanceCmd.Flags().String("instance-arn", "", "The ARN of the instance of IAM Identity Center under which the operation will run.")
	ssoAdmin_describeInstanceCmd.MarkFlagRequired("instance-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_describeInstanceCmd)
}
