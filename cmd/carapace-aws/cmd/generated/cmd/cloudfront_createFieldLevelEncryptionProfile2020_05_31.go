package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createFieldLevelEncryptionProfile2020_05_31Cmd = &cobra.Command{
	Use:   "create-field-level-encryption-profile2020_05_31",
	Short: "Create a field-level encryption profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createFieldLevelEncryptionProfile2020_05_31Cmd).Standalone()

	cloudfront_createFieldLevelEncryptionProfile2020_05_31Cmd.Flags().String("field-level-encryption-profile-config", "", "The request to create a field-level encryption profile.")
	cloudfront_createFieldLevelEncryptionProfile2020_05_31Cmd.MarkFlagRequired("field-level-encryption-profile-config")
	cloudfrontCmd.AddCommand(cloudfront_createFieldLevelEncryptionProfile2020_05_31Cmd)
}
