package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_getBackendApimodelsCmd = &cobra.Command{
	Use:   "get-backend-apimodels",
	Short: "Gets a model introspection schema for an existing backend API resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_getBackendApimodelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_getBackendApimodelsCmd).Standalone()

		amplifybackend_getBackendApimodelsCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_getBackendApimodelsCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_getBackendApimodelsCmd.Flags().String("resource-name", "", "The name of this resource.")
		amplifybackend_getBackendApimodelsCmd.MarkFlagRequired("app-id")
		amplifybackend_getBackendApimodelsCmd.MarkFlagRequired("backend-environment-name")
		amplifybackend_getBackendApimodelsCmd.MarkFlagRequired("resource-name")
	})
	amplifybackendCmd.AddCommand(amplifybackend_getBackendApimodelsCmd)
}
