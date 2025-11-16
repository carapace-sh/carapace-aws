package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_updateInstanceCmd = &cobra.Command{
	Use:   "update-instance",
	Short: "Update the details for the instance of IAM Identity Center that is owned by the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_updateInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_updateInstanceCmd).Standalone()

		ssoAdmin_updateInstanceCmd.Flags().String("encryption-configuration", "", "Specifies the encryption configuration for your IAM Identity Center instance.")
		ssoAdmin_updateInstanceCmd.Flags().String("instance-arn", "", "The ARN of the instance of IAM Identity Center under which the operation will run.")
		ssoAdmin_updateInstanceCmd.Flags().String("name", "", "Updates the instance name.")
		ssoAdmin_updateInstanceCmd.MarkFlagRequired("instance-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_updateInstanceCmd)
}
