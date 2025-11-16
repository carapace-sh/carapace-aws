package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_createInstanceAccessControlAttributeConfigurationCmd = &cobra.Command{
	Use:   "create-instance-access-control-attribute-configuration",
	Short: "Enables the attributes-based access control (ABAC) feature for the specified IAM Identity Center instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_createInstanceAccessControlAttributeConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_createInstanceAccessControlAttributeConfigurationCmd).Standalone()

		ssoAdmin_createInstanceAccessControlAttributeConfigurationCmd.Flags().String("instance-access-control-attribute-configuration", "", "Specifies the IAM Identity Center identity store attributes to add to your ABAC configuration.")
		ssoAdmin_createInstanceAccessControlAttributeConfigurationCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_createInstanceAccessControlAttributeConfigurationCmd.MarkFlagRequired("instance-access-control-attribute-configuration")
		ssoAdmin_createInstanceAccessControlAttributeConfigurationCmd.MarkFlagRequired("instance-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_createInstanceAccessControlAttributeConfigurationCmd)
}
