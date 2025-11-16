package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_getDashboardEmbedUrlCmd = &cobra.Command{
	Use:   "get-dashboard-embed-url",
	Short: "Generates a temporary session URL and authorization code(bearer token) that you can use to embed an Amazon Quick Sight read-only dashboard in your website or application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_getDashboardEmbedUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_getDashboardEmbedUrlCmd).Standalone()

		quicksight_getDashboardEmbedUrlCmd.Flags().String("additional-dashboard-ids", "", "A list of one or more dashboard IDs that you want anonymous users to have tempporary access to.")
		quicksight_getDashboardEmbedUrlCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the dashboard that you're embedding.")
		quicksight_getDashboardEmbedUrlCmd.Flags().String("dashboard-id", "", "The ID for the dashboard, also added to the Identity and Access Management (IAM) policy.")
		quicksight_getDashboardEmbedUrlCmd.Flags().String("identity-type", "", "The authentication method that the user uses to sign in.")
		quicksight_getDashboardEmbedUrlCmd.Flags().String("namespace", "", "The Amazon Quick Sight namespace that contains the dashboard IDs in this request.")
		quicksight_getDashboardEmbedUrlCmd.Flags().Bool("no-reset-disabled", false, "Remove the reset button on the embedded dashboard.")
		quicksight_getDashboardEmbedUrlCmd.Flags().Bool("no-state-persistence-enabled", false, "Adds persistence of state for the user session in an embedded dashboard.")
		quicksight_getDashboardEmbedUrlCmd.Flags().Bool("no-undo-redo-disabled", false, "Remove the undo/redo button on the embedded dashboard.")
		quicksight_getDashboardEmbedUrlCmd.Flags().Bool("reset-disabled", false, "Remove the reset button on the embedded dashboard.")
		quicksight_getDashboardEmbedUrlCmd.Flags().String("session-lifetime-in-minutes", "", "How many minutes the session is valid.")
		quicksight_getDashboardEmbedUrlCmd.Flags().Bool("state-persistence-enabled", false, "Adds persistence of state for the user session in an embedded dashboard.")
		quicksight_getDashboardEmbedUrlCmd.Flags().Bool("undo-redo-disabled", false, "Remove the undo/redo button on the embedded dashboard.")
		quicksight_getDashboardEmbedUrlCmd.Flags().String("user-arn", "", "The Amazon Quick Suite user's Amazon Resource Name (ARN), for use with `QUICKSIGHT` identity type.")
		quicksight_getDashboardEmbedUrlCmd.MarkFlagRequired("aws-account-id")
		quicksight_getDashboardEmbedUrlCmd.MarkFlagRequired("dashboard-id")
		quicksight_getDashboardEmbedUrlCmd.MarkFlagRequired("identity-type")
		quicksight_getDashboardEmbedUrlCmd.Flag("no-reset-disabled").Hidden = true
		quicksight_getDashboardEmbedUrlCmd.Flag("no-state-persistence-enabled").Hidden = true
		quicksight_getDashboardEmbedUrlCmd.Flag("no-undo-redo-disabled").Hidden = true
	})
	quicksightCmd.AddCommand(quicksight_getDashboardEmbedUrlCmd)
}
