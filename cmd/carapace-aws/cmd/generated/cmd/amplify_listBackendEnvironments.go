package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_listBackendEnvironmentsCmd = &cobra.Command{
	Use:   "list-backend-environments",
	Short: "Lists the backend environments for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_listBackendEnvironmentsCmd).Standalone()

	amplify_listBackendEnvironmentsCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
	amplify_listBackendEnvironmentsCmd.Flags().String("environment-name", "", "The name of the backend environment")
	amplify_listBackendEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of records to list in a single response.")
	amplify_listBackendEnvironmentsCmd.Flags().String("next-token", "", "A pagination token.")
	amplify_listBackendEnvironmentsCmd.MarkFlagRequired("app-id")
	amplifyCmd.AddCommand(amplify_listBackendEnvironmentsCmd)
}
