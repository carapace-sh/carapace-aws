package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates application properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_updateApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_updateApplicationCmd).Standalone()

		ssoAdmin_updateApplicationCmd.Flags().String("application-arn", "", "Specifies the ARN of the application.")
		ssoAdmin_updateApplicationCmd.Flags().String("description", "", "The description of the .")
		ssoAdmin_updateApplicationCmd.Flags().String("name", "", "Specifies the updated name for the application.")
		ssoAdmin_updateApplicationCmd.Flags().String("portal-options", "", "A structure that describes the options for the portal associated with an application.")
		ssoAdmin_updateApplicationCmd.Flags().String("status", "", "Specifies whether the application is enabled or disabled.")
		ssoAdmin_updateApplicationCmd.MarkFlagRequired("application-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_updateApplicationCmd)
}
