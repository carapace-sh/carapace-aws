package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_describeInstanceAccessControlAttributeConfigurationCmd = &cobra.Command{
	Use:   "describe-instance-access-control-attribute-configuration",
	Short: "Returns the list of IAM Identity Center identity store attributes that have been configured to work with attributes-based access control (ABAC) for the specified IAM Identity Center instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_describeInstanceAccessControlAttributeConfigurationCmd).Standalone()

	ssoAdmin_describeInstanceAccessControlAttributeConfigurationCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_describeInstanceAccessControlAttributeConfigurationCmd.MarkFlagRequired("instance-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_describeInstanceAccessControlAttributeConfigurationCmd)
}
