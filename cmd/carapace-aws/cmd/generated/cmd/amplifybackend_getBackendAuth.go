package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_getBackendAuthCmd = &cobra.Command{
	Use:   "get-backend-auth",
	Short: "Gets a backend auth details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_getBackendAuthCmd).Standalone()

	amplifybackend_getBackendAuthCmd.Flags().String("app-id", "", "The app ID.")
	amplifybackend_getBackendAuthCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
	amplifybackend_getBackendAuthCmd.Flags().String("resource-name", "", "The name of this resource.")
	amplifybackend_getBackendAuthCmd.MarkFlagRequired("app-id")
	amplifybackend_getBackendAuthCmd.MarkFlagRequired("backend-environment-name")
	amplifybackend_getBackendAuthCmd.MarkFlagRequired("resource-name")
	amplifybackendCmd.AddCommand(amplifybackend_getBackendAuthCmd)
}
