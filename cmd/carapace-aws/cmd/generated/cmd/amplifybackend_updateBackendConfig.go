package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_updateBackendConfigCmd = &cobra.Command{
	Use:   "update-backend-config",
	Short: "Updates the AWS resources required to access the Amplify Admin UI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_updateBackendConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_updateBackendConfigCmd).Standalone()

		amplifybackend_updateBackendConfigCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_updateBackendConfigCmd.Flags().String("login-auth-config", "", "Describes the Amazon Cognito configuration for Admin UI access.")
		amplifybackend_updateBackendConfigCmd.MarkFlagRequired("app-id")
	})
	amplifybackendCmd.AddCommand(amplifybackend_updateBackendConfigCmd)
}
