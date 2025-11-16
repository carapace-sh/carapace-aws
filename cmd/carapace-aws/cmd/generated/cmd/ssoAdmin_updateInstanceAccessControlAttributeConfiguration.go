package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_updateInstanceAccessControlAttributeConfigurationCmd = &cobra.Command{
	Use:   "update-instance-access-control-attribute-configuration",
	Short: "Updates the IAM Identity Center identity store attributes that you can use with the IAM Identity Center instance for attributes-based access control (ABAC).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_updateInstanceAccessControlAttributeConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_updateInstanceAccessControlAttributeConfigurationCmd).Standalone()

		ssoAdmin_updateInstanceAccessControlAttributeConfigurationCmd.Flags().String("instance-access-control-attribute-configuration", "", "Updates the attributes for your ABAC configuration.")
		ssoAdmin_updateInstanceAccessControlAttributeConfigurationCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_updateInstanceAccessControlAttributeConfigurationCmd.MarkFlagRequired("instance-access-control-attribute-configuration")
		ssoAdmin_updateInstanceAccessControlAttributeConfigurationCmd.MarkFlagRequired("instance-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_updateInstanceAccessControlAttributeConfigurationCmd)
}
