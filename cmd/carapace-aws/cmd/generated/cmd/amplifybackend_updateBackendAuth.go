package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_updateBackendAuthCmd = &cobra.Command{
	Use:   "update-backend-auth",
	Short: "Updates an existing backend authentication resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_updateBackendAuthCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_updateBackendAuthCmd).Standalone()

		amplifybackend_updateBackendAuthCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_updateBackendAuthCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_updateBackendAuthCmd.Flags().String("resource-config", "", "The resource configuration for this request object.")
		amplifybackend_updateBackendAuthCmd.Flags().String("resource-name", "", "The name of this resource.")
		amplifybackend_updateBackendAuthCmd.MarkFlagRequired("app-id")
		amplifybackend_updateBackendAuthCmd.MarkFlagRequired("backend-environment-name")
		amplifybackend_updateBackendAuthCmd.MarkFlagRequired("resource-config")
		amplifybackend_updateBackendAuthCmd.MarkFlagRequired("resource-name")
	})
	amplifybackendCmd.AddCommand(amplifybackend_updateBackendAuthCmd)
}
