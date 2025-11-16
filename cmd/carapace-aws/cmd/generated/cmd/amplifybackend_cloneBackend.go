package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_cloneBackendCmd = &cobra.Command{
	Use:   "clone-backend",
	Short: "This operation clones an existing backend.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_cloneBackendCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_cloneBackendCmd).Standalone()

		amplifybackend_cloneBackendCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_cloneBackendCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_cloneBackendCmd.Flags().String("target-environment-name", "", "The name of the destination backend environment to be created.")
		amplifybackend_cloneBackendCmd.MarkFlagRequired("app-id")
		amplifybackend_cloneBackendCmd.MarkFlagRequired("backend-environment-name")
		amplifybackend_cloneBackendCmd.MarkFlagRequired("target-environment-name")
	})
	amplifybackendCmd.AddCommand(amplifybackend_cloneBackendCmd)
}
