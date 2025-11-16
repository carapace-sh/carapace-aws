package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_deleteInstanceAccessControlAttributeConfigurationCmd = &cobra.Command{
	Use:   "delete-instance-access-control-attribute-configuration",
	Short: "Disables the attributes-based access control (ABAC) feature for the specified IAM Identity Center instance and deletes all of the attribute mappings that have been configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_deleteInstanceAccessControlAttributeConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_deleteInstanceAccessControlAttributeConfigurationCmd).Standalone()

		ssoAdmin_deleteInstanceAccessControlAttributeConfigurationCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_deleteInstanceAccessControlAttributeConfigurationCmd.MarkFlagRequired("instance-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_deleteInstanceAccessControlAttributeConfigurationCmd)
}
