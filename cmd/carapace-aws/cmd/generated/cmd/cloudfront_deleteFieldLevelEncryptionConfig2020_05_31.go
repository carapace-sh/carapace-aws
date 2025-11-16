package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteFieldLevelEncryptionConfig2020_05_31Cmd = &cobra.Command{
	Use:   "delete-field-level-encryption-config2020_05_31",
	Short: "Remove a field-level encryption configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteFieldLevelEncryptionConfig2020_05_31Cmd).Standalone()

	cloudfront_deleteFieldLevelEncryptionConfig2020_05_31Cmd.Flags().String("id", "", "The ID of the configuration you want to delete from CloudFront.")
	cloudfront_deleteFieldLevelEncryptionConfig2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the configuration identity to delete.")
	cloudfront_deleteFieldLevelEncryptionConfig2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_deleteFieldLevelEncryptionConfig2020_05_31Cmd)
}
