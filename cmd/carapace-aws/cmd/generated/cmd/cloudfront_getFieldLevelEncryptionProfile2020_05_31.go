package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getFieldLevelEncryptionProfile2020_05_31Cmd = &cobra.Command{
	Use:   "get-field-level-encryption-profile2020_05_31",
	Short: "Get the field-level encryption profile information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getFieldLevelEncryptionProfile2020_05_31Cmd).Standalone()

	cloudfront_getFieldLevelEncryptionProfile2020_05_31Cmd.Flags().String("id", "", "Get the ID for the field-level encryption profile information.")
	cloudfront_getFieldLevelEncryptionProfile2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_getFieldLevelEncryptionProfile2020_05_31Cmd)
}
