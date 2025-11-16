package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_updateBackendJobCmd = &cobra.Command{
	Use:   "update-backend-job",
	Short: "Updates a specific job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_updateBackendJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_updateBackendJobCmd).Standalone()

		amplifybackend_updateBackendJobCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_updateBackendJobCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_updateBackendJobCmd.Flags().String("job-id", "", "The ID for the job.")
		amplifybackend_updateBackendJobCmd.Flags().String("operation", "", "Filters the list of response objects to include only those with the specified operation name.")
		amplifybackend_updateBackendJobCmd.Flags().String("status", "", "Filters the list of response objects to include only those with the specified status.")
		amplifybackend_updateBackendJobCmd.MarkFlagRequired("app-id")
		amplifybackend_updateBackendJobCmd.MarkFlagRequired("backend-environment-name")
		amplifybackend_updateBackendJobCmd.MarkFlagRequired("job-id")
	})
	amplifybackendCmd.AddCommand(amplifybackend_updateBackendJobCmd)
}
