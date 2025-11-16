package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_generateEmbedUrlForAnonymousUserCmd = &cobra.Command{
	Use:   "generate-embed-url-for-anonymous-user",
	Short: "Generates an embed URL that you can use to embed an Amazon Quick Suite dashboard or visual in your website, without having to register any reader users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_generateEmbedUrlForAnonymousUserCmd).Standalone()

	quicksight_generateEmbedUrlForAnonymousUserCmd.Flags().String("allowed-domains", "", "The domains that you want to add to the allow list for access to the generated URL that is then embedded.")
	quicksight_generateEmbedUrlForAnonymousUserCmd.Flags().String("authorized-resource-arns", "", "The Amazon Resource Names (ARNs) for the Quick Sight resources that the user is authorized to access during the lifetime of the session.")
	quicksight_generateEmbedUrlForAnonymousUserCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the dashboard that you're embedding.")
	quicksight_generateEmbedUrlForAnonymousUserCmd.Flags().String("experience-configuration", "", "The configuration of the experience that you are embedding.")
	quicksight_generateEmbedUrlForAnonymousUserCmd.Flags().String("namespace", "", "The Amazon Quick Sight namespace that the anonymous user virtually belongs to.")
	quicksight_generateEmbedUrlForAnonymousUserCmd.Flags().String("session-lifetime-in-minutes", "", "How many minutes the session is valid.")
	quicksight_generateEmbedUrlForAnonymousUserCmd.Flags().String("session-tags", "", "The session tags used for row-level security.")
	quicksight_generateEmbedUrlForAnonymousUserCmd.MarkFlagRequired("authorized-resource-arns")
	quicksight_generateEmbedUrlForAnonymousUserCmd.MarkFlagRequired("aws-account-id")
	quicksight_generateEmbedUrlForAnonymousUserCmd.MarkFlagRequired("experience-configuration")
	quicksight_generateEmbedUrlForAnonymousUserCmd.MarkFlagRequired("namespace")
	quicksightCmd.AddCommand(quicksight_generateEmbedUrlForAnonymousUserCmd)
}
