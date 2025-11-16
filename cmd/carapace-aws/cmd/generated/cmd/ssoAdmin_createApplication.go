package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates an OAuth 2.0 customer managed application in IAM Identity Center for the given application provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_createApplicationCmd).Standalone()

		ssoAdmin_createApplicationCmd.Flags().String("application-provider-arn", "", "The ARN of the application provider under which the operation will run.")
		ssoAdmin_createApplicationCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive ID that you provide to ensure the idempotency of the request.")
		ssoAdmin_createApplicationCmd.Flags().String("description", "", "The description of the .")
		ssoAdmin_createApplicationCmd.Flags().String("instance-arn", "", "The ARN of the instance of IAM Identity Center under which the operation will run.")
		ssoAdmin_createApplicationCmd.Flags().String("name", "", "The name of the .")
		ssoAdmin_createApplicationCmd.Flags().String("portal-options", "", "A structure that describes the options for the portal associated with an application.")
		ssoAdmin_createApplicationCmd.Flags().String("status", "", "Specifies whether the application is enabled or disabled.")
		ssoAdmin_createApplicationCmd.Flags().String("tags", "", "Specifies tags to be attached to the application.")
		ssoAdmin_createApplicationCmd.MarkFlagRequired("application-provider-arn")
		ssoAdmin_createApplicationCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_createApplicationCmd.MarkFlagRequired("name")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_createApplicationCmd)
}
