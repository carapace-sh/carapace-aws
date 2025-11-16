package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_generateBackendApimodelsCmd = &cobra.Command{
	Use:   "generate-backend-apimodels",
	Short: "Generates a model schema for an existing backend API resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_generateBackendApimodelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_generateBackendApimodelsCmd).Standalone()

		amplifybackend_generateBackendApimodelsCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_generateBackendApimodelsCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_generateBackendApimodelsCmd.Flags().String("resource-name", "", "The name of this resource.")
		amplifybackend_generateBackendApimodelsCmd.MarkFlagRequired("app-id")
		amplifybackend_generateBackendApimodelsCmd.MarkFlagRequired("backend-environment-name")
		amplifybackend_generateBackendApimodelsCmd.MarkFlagRequired("resource-name")
	})
	amplifybackendCmd.AddCommand(amplifybackend_generateBackendApimodelsCmd)
}
