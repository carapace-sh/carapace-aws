package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listFieldLevelEncryptionConfigs2020_05_31Cmd = &cobra.Command{
	Use:   "list-field-level-encryption-configs2020_05_31",
	Short: "List all field-level encryption configurations that have been created in CloudFront for this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listFieldLevelEncryptionConfigs2020_05_31Cmd).Standalone()

	cloudfront_listFieldLevelEncryptionConfigs2020_05_31Cmd.Flags().String("marker", "", "Use this when paginating results to indicate where to begin in your list of configurations.")
	cloudfront_listFieldLevelEncryptionConfigs2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of field-level encryption configurations you want in the response body.")
	cloudfrontCmd.AddCommand(cloudfront_listFieldLevelEncryptionConfigs2020_05_31Cmd)
}
