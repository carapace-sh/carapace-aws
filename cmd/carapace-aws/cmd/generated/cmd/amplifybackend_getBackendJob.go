package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_getBackendJobCmd = &cobra.Command{
	Use:   "get-backend-job",
	Short: "Returns information about a specific job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_getBackendJobCmd).Standalone()

	amplifybackend_getBackendJobCmd.Flags().String("app-id", "", "The app ID.")
	amplifybackend_getBackendJobCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
	amplifybackend_getBackendJobCmd.Flags().String("job-id", "", "The ID for the job.")
	amplifybackend_getBackendJobCmd.MarkFlagRequired("app-id")
	amplifybackend_getBackendJobCmd.MarkFlagRequired("backend-environment-name")
	amplifybackend_getBackendJobCmd.MarkFlagRequired("job-id")
	amplifybackendCmd.AddCommand(amplifybackend_getBackendJobCmd)
}
