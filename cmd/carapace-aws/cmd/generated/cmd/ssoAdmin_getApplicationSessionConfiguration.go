package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_getApplicationSessionConfigurationCmd = &cobra.Command{
	Use:   "get-application-session-configuration",
	Short: "Retrieves the session configuration for an application in IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_getApplicationSessionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_getApplicationSessionConfigurationCmd).Standalone()

		ssoAdmin_getApplicationSessionConfigurationCmd.Flags().String("application-arn", "", "The Amazon Resource Name (ARN) of the application for which to retrieve the session configuration.")
		ssoAdmin_getApplicationSessionConfigurationCmd.MarkFlagRequired("application-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_getApplicationSessionConfigurationCmd)
}
