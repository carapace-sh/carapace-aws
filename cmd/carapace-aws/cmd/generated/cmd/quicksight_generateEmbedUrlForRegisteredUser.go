package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_generateEmbedUrlForRegisteredUserCmd = &cobra.Command{
	Use:   "generate-embed-url-for-registered-user",
	Short: "Generates an embed URL that you can use to embed an Amazon Quick Suite experience in your website.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_generateEmbedUrlForRegisteredUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_generateEmbedUrlForRegisteredUserCmd).Standalone()

		quicksight_generateEmbedUrlForRegisteredUserCmd.Flags().String("allowed-domains", "", "The domains that you want to add to the allow list for access to the generated URL that is then embedded.")
		quicksight_generateEmbedUrlForRegisteredUserCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the dashboard that you're embedding.")
		quicksight_generateEmbedUrlForRegisteredUserCmd.Flags().String("experience-configuration", "", "The experience that you want to embed.")
		quicksight_generateEmbedUrlForRegisteredUserCmd.Flags().String("session-lifetime-in-minutes", "", "How many minutes the session is valid.")
		quicksight_generateEmbedUrlForRegisteredUserCmd.Flags().String("user-arn", "", "The Amazon Resource Name for the registered user.")
		quicksight_generateEmbedUrlForRegisteredUserCmd.MarkFlagRequired("aws-account-id")
		quicksight_generateEmbedUrlForRegisteredUserCmd.MarkFlagRequired("experience-configuration")
		quicksight_generateEmbedUrlForRegisteredUserCmd.MarkFlagRequired("user-arn")
	})
	quicksightCmd.AddCommand(quicksight_generateEmbedUrlForRegisteredUserCmd)
}
