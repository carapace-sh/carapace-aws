package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_createBackendApiCmd = &cobra.Command{
	Use:   "create-backend-api",
	Short: "Creates a new backend API resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_createBackendApiCmd).Standalone()

	amplifybackend_createBackendApiCmd.Flags().String("app-id", "", "The app ID.")
	amplifybackend_createBackendApiCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
	amplifybackend_createBackendApiCmd.Flags().String("resource-config", "", "The resource configuration for this request.")
	amplifybackend_createBackendApiCmd.Flags().String("resource-name", "", "The name of this resource.")
	amplifybackend_createBackendApiCmd.MarkFlagRequired("app-id")
	amplifybackend_createBackendApiCmd.MarkFlagRequired("backend-environment-name")
	amplifybackend_createBackendApiCmd.MarkFlagRequired("resource-config")
	amplifybackend_createBackendApiCmd.MarkFlagRequired("resource-name")
	amplifybackendCmd.AddCommand(amplifybackend_createBackendApiCmd)
}
