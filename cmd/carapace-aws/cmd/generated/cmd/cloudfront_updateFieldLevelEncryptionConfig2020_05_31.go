package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateFieldLevelEncryptionConfig2020_05_31Cmd = &cobra.Command{
	Use:   "update-field-level-encryption-config2020_05_31",
	Short: "Update a field-level encryption configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateFieldLevelEncryptionConfig2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_updateFieldLevelEncryptionConfig2020_05_31Cmd).Standalone()

		cloudfront_updateFieldLevelEncryptionConfig2020_05_31Cmd.Flags().String("field-level-encryption-config", "", "Request to update a field-level encryption configuration.")
		cloudfront_updateFieldLevelEncryptionConfig2020_05_31Cmd.Flags().String("id", "", "The ID of the configuration you want to update.")
		cloudfront_updateFieldLevelEncryptionConfig2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the configuration identity to update.")
		cloudfront_updateFieldLevelEncryptionConfig2020_05_31Cmd.MarkFlagRequired("field-level-encryption-config")
		cloudfront_updateFieldLevelEncryptionConfig2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_updateFieldLevelEncryptionConfig2020_05_31Cmd)
}
