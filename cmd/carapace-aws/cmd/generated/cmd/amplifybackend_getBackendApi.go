package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_getBackendApiCmd = &cobra.Command{
	Use:   "get-backend-api",
	Short: "Gets the details for a backend API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_getBackendApiCmd).Standalone()

	amplifybackend_getBackendApiCmd.Flags().String("app-id", "", "The app ID.")
	amplifybackend_getBackendApiCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
	amplifybackend_getBackendApiCmd.Flags().String("resource-config", "", "Defines the resource configuration for the data model in your Amplify project.")
	amplifybackend_getBackendApiCmd.Flags().String("resource-name", "", "The name of this resource.")
	amplifybackend_getBackendApiCmd.MarkFlagRequired("app-id")
	amplifybackend_getBackendApiCmd.MarkFlagRequired("backend-environment-name")
	amplifybackend_getBackendApiCmd.MarkFlagRequired("resource-name")
	amplifybackendCmd.AddCommand(amplifybackend_getBackendApiCmd)
}
