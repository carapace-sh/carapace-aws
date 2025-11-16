package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createFieldLevelEncryptionConfig2020_05_31Cmd = &cobra.Command{
	Use:   "create-field-level-encryption-config2020_05_31",
	Short: "Create a new field-level encryption configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createFieldLevelEncryptionConfig2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_createFieldLevelEncryptionConfig2020_05_31Cmd).Standalone()

		cloudfront_createFieldLevelEncryptionConfig2020_05_31Cmd.Flags().String("field-level-encryption-config", "", "The request to create a new field-level encryption configuration.")
		cloudfront_createFieldLevelEncryptionConfig2020_05_31Cmd.MarkFlagRequired("field-level-encryption-config")
	})
	cloudfrontCmd.AddCommand(cloudfront_createFieldLevelEncryptionConfig2020_05_31Cmd)
}
