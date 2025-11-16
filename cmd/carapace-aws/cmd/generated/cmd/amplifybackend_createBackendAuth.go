package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_createBackendAuthCmd = &cobra.Command{
	Use:   "create-backend-auth",
	Short: "Creates a new backend authentication resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_createBackendAuthCmd).Standalone()

	amplifybackend_createBackendAuthCmd.Flags().String("app-id", "", "The app ID.")
	amplifybackend_createBackendAuthCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
	amplifybackend_createBackendAuthCmd.Flags().String("resource-config", "", "The resource configuration for this request object.")
	amplifybackend_createBackendAuthCmd.Flags().String("resource-name", "", "The name of this resource.")
	amplifybackend_createBackendAuthCmd.MarkFlagRequired("app-id")
	amplifybackend_createBackendAuthCmd.MarkFlagRequired("backend-environment-name")
	amplifybackend_createBackendAuthCmd.MarkFlagRequired("resource-config")
	amplifybackend_createBackendAuthCmd.MarkFlagRequired("resource-name")
	amplifybackendCmd.AddCommand(amplifybackend_createBackendAuthCmd)
}
