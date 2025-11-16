package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteFieldLevelEncryptionProfile2020_05_31Cmd = &cobra.Command{
	Use:   "delete-field-level-encryption-profile2020_05_31",
	Short: "Remove a field-level encryption profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteFieldLevelEncryptionProfile2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_deleteFieldLevelEncryptionProfile2020_05_31Cmd).Standalone()

		cloudfront_deleteFieldLevelEncryptionProfile2020_05_31Cmd.Flags().String("id", "", "Request the ID of the profile you want to delete from CloudFront.")
		cloudfront_deleteFieldLevelEncryptionProfile2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the profile to delete.")
		cloudfront_deleteFieldLevelEncryptionProfile2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_deleteFieldLevelEncryptionProfile2020_05_31Cmd)
}
