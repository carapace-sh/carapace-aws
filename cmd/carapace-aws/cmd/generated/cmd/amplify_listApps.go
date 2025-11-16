package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_listAppsCmd = &cobra.Command{
	Use:   "list-apps",
	Short: "Returns a list of the existing Amplify apps.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_listAppsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_listAppsCmd).Standalone()

		amplify_listAppsCmd.Flags().String("max-results", "", "The maximum number of records to list in a single response.")
		amplify_listAppsCmd.Flags().String("next-token", "", "A pagination token.")
	})
	amplifyCmd.AddCommand(amplify_listAppsCmd)
}
