package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_getSessionEmbedUrlCmd = &cobra.Command{
	Use:   "get-session-embed-url",
	Short: "Generates a session URL and authorization code that you can use to embed the Amazon Amazon Quick Sight console in your web server code.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_getSessionEmbedUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_getSessionEmbedUrlCmd).Standalone()

		quicksight_getSessionEmbedUrlCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account associated with your Amazon Quick Sight subscription.")
		quicksight_getSessionEmbedUrlCmd.Flags().String("entry-point", "", "The URL you use to access the embedded session.")
		quicksight_getSessionEmbedUrlCmd.Flags().String("session-lifetime-in-minutes", "", "How many minutes the session is valid.")
		quicksight_getSessionEmbedUrlCmd.Flags().String("user-arn", "", "The Amazon Quick Suite user's Amazon Resource Name (ARN), for use with `QUICKSIGHT` identity type.")
		quicksight_getSessionEmbedUrlCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_getSessionEmbedUrlCmd)
}
