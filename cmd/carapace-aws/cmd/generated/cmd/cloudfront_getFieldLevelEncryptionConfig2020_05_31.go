package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getFieldLevelEncryptionConfig2020_05_31Cmd = &cobra.Command{
	Use:   "get-field-level-encryption-config2020_05_31",
	Short: "Get the field-level encryption configuration information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getFieldLevelEncryptionConfig2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getFieldLevelEncryptionConfig2020_05_31Cmd).Standalone()

		cloudfront_getFieldLevelEncryptionConfig2020_05_31Cmd.Flags().String("id", "", "Request the ID for the field-level encryption configuration information.")
		cloudfront_getFieldLevelEncryptionConfig2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_getFieldLevelEncryptionConfig2020_05_31Cmd)
}
