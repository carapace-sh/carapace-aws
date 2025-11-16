package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_getAppCmd = &cobra.Command{
	Use:   "get-app",
	Short: "Returns an existing Amplify app specified by an app ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_getAppCmd).Standalone()

	amplify_getAppCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
	amplify_getAppCmd.MarkFlagRequired("app-id")
	amplifyCmd.AddCommand(amplify_getAppCmd)
}
