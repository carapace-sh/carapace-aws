package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateFieldLevelEncryptionProfile2020_05_31Cmd = &cobra.Command{
	Use:   "update-field-level-encryption-profile2020_05_31",
	Short: "Update a field-level encryption profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateFieldLevelEncryptionProfile2020_05_31Cmd).Standalone()

	cloudfront_updateFieldLevelEncryptionProfile2020_05_31Cmd.Flags().String("field-level-encryption-profile-config", "", "Request to update a field-level encryption profile.")
	cloudfront_updateFieldLevelEncryptionProfile2020_05_31Cmd.Flags().String("id", "", "The ID of the field-level encryption profile request.")
	cloudfront_updateFieldLevelEncryptionProfile2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the profile identity to update.")
	cloudfront_updateFieldLevelEncryptionProfile2020_05_31Cmd.MarkFlagRequired("field-level-encryption-profile-config")
	cloudfront_updateFieldLevelEncryptionProfile2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_updateFieldLevelEncryptionProfile2020_05_31Cmd)
}
