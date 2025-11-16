package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_removeBackendConfigCmd = &cobra.Command{
	Use:   "remove-backend-config",
	Short: "Removes the AWS resources required to access the Amplify Admin UI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_removeBackendConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_removeBackendConfigCmd).Standalone()

		amplifybackend_removeBackendConfigCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_removeBackendConfigCmd.MarkFlagRequired("app-id")
	})
	amplifybackendCmd.AddCommand(amplifybackend_removeBackendConfigCmd)
}
