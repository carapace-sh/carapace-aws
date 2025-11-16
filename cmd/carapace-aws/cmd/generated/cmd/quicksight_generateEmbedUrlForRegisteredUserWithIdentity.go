package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_generateEmbedUrlForRegisteredUserWithIdentityCmd = &cobra.Command{
	Use:   "generate-embed-url-for-registered-user-with-identity",
	Short: "Generates an embed URL that you can use to embed an Amazon Quick Sight experience in your website.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_generateEmbedUrlForRegisteredUserWithIdentityCmd).Standalone()

	quicksight_generateEmbedUrlForRegisteredUserWithIdentityCmd.Flags().String("allowed-domains", "", "A list of domains to be allowed to generate the embed URL.")
	quicksight_generateEmbedUrlForRegisteredUserWithIdentityCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services registered user.")
	quicksight_generateEmbedUrlForRegisteredUserWithIdentityCmd.Flags().String("experience-configuration", "", "")
	quicksight_generateEmbedUrlForRegisteredUserWithIdentityCmd.Flags().String("session-lifetime-in-minutes", "", "The validity of the session in minutes.")
	quicksight_generateEmbedUrlForRegisteredUserWithIdentityCmd.MarkFlagRequired("aws-account-id")
	quicksight_generateEmbedUrlForRegisteredUserWithIdentityCmd.MarkFlagRequired("experience-configuration")
	quicksightCmd.AddCommand(quicksight_generateEmbedUrlForRegisteredUserWithIdentityCmd)
}
