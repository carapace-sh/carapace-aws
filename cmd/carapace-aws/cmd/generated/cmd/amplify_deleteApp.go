package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_deleteAppCmd = &cobra.Command{
	Use:   "delete-app",
	Short: "Deletes an existing Amplify app specified by an app ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_deleteAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_deleteAppCmd).Standalone()

		amplify_deleteAppCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
		amplify_deleteAppCmd.MarkFlagRequired("app-id")
	})
	amplifyCmd.AddCommand(amplify_deleteAppCmd)
}
