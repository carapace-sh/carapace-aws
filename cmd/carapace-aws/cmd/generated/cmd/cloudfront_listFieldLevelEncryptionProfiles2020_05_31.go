package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listFieldLevelEncryptionProfiles2020_05_31Cmd = &cobra.Command{
	Use:   "list-field-level-encryption-profiles2020_05_31",
	Short: "Request a list of field-level encryption profiles that have been created in CloudFront for this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listFieldLevelEncryptionProfiles2020_05_31Cmd).Standalone()

	cloudfront_listFieldLevelEncryptionProfiles2020_05_31Cmd.Flags().String("marker", "", "Use this when paginating results to indicate where to begin in your list of profiles.")
	cloudfront_listFieldLevelEncryptionProfiles2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of field-level encryption profiles you want in the response body.")
	cloudfrontCmd.AddCommand(cloudfront_listFieldLevelEncryptionProfiles2020_05_31Cmd)
}
