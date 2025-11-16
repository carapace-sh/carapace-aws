package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_updateBackendApiCmd = &cobra.Command{
	Use:   "update-backend-api",
	Short: "Updates an existing backend API resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_updateBackendApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_updateBackendApiCmd).Standalone()

		amplifybackend_updateBackendApiCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_updateBackendApiCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_updateBackendApiCmd.Flags().String("resource-config", "", "Defines the resource configuration for the data model in your Amplify project.")
		amplifybackend_updateBackendApiCmd.Flags().String("resource-name", "", "The name of this resource.")
		amplifybackend_updateBackendApiCmd.MarkFlagRequired("app-id")
		amplifybackend_updateBackendApiCmd.MarkFlagRequired("backend-environment-name")
		amplifybackend_updateBackendApiCmd.MarkFlagRequired("resource-name")
	})
	amplifybackendCmd.AddCommand(amplifybackend_updateBackendApiCmd)
}
