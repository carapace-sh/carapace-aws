package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_putApplicationSessionConfigurationCmd = &cobra.Command{
	Use:   "put-application-session-configuration",
	Short: "Updates the session configuration for an application in IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_putApplicationSessionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_putApplicationSessionConfigurationCmd).Standalone()

		ssoAdmin_putApplicationSessionConfigurationCmd.Flags().String("application-arn", "", "The Amazon Resource Name (ARN) of the application for which to update the session configuration.")
		ssoAdmin_putApplicationSessionConfigurationCmd.Flags().String("user-background-session-application-status", "", "The status of user background sessions for the application.")
		ssoAdmin_putApplicationSessionConfigurationCmd.MarkFlagRequired("application-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_putApplicationSessionConfigurationCmd)
}
