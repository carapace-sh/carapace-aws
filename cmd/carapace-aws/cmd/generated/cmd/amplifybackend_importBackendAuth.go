package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_importBackendAuthCmd = &cobra.Command{
	Use:   "import-backend-auth",
	Short: "Imports an existing backend authentication resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_importBackendAuthCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_importBackendAuthCmd).Standalone()

		amplifybackend_importBackendAuthCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_importBackendAuthCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_importBackendAuthCmd.Flags().String("identity-pool-id", "", "The ID of the Amazon Cognito identity pool.")
		amplifybackend_importBackendAuthCmd.Flags().String("native-client-id", "", "The ID of the Amazon Cognito native client.")
		amplifybackend_importBackendAuthCmd.Flags().String("user-pool-id", "", "The ID of the Amazon Cognito user pool.")
		amplifybackend_importBackendAuthCmd.Flags().String("web-client-id", "", "The ID of the Amazon Cognito web client.")
		amplifybackend_importBackendAuthCmd.MarkFlagRequired("app-id")
		amplifybackend_importBackendAuthCmd.MarkFlagRequired("backend-environment-name")
		amplifybackend_importBackendAuthCmd.MarkFlagRequired("native-client-id")
		amplifybackend_importBackendAuthCmd.MarkFlagRequired("user-pool-id")
		amplifybackend_importBackendAuthCmd.MarkFlagRequired("web-client-id")
	})
	amplifybackendCmd.AddCommand(amplifybackend_importBackendAuthCmd)
}
