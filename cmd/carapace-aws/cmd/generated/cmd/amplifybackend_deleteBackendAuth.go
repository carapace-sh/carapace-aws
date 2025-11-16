package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_deleteBackendAuthCmd = &cobra.Command{
	Use:   "delete-backend-auth",
	Short: "Deletes an existing backend authentication resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_deleteBackendAuthCmd).Standalone()

	amplifybackend_deleteBackendAuthCmd.Flags().String("app-id", "", "The app ID.")
	amplifybackend_deleteBackendAuthCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
	amplifybackend_deleteBackendAuthCmd.Flags().String("resource-name", "", "The name of this resource.")
	amplifybackend_deleteBackendAuthCmd.MarkFlagRequired("app-id")
	amplifybackend_deleteBackendAuthCmd.MarkFlagRequired("backend-environment-name")
	amplifybackend_deleteBackendAuthCmd.MarkFlagRequired("resource-name")
	amplifybackendCmd.AddCommand(amplifybackend_deleteBackendAuthCmd)
}
