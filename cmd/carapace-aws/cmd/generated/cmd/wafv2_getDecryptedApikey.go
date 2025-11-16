package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_getDecryptedApikeyCmd = &cobra.Command{
	Use:   "get-decrypted-apikey",
	Short: "Returns your API key in decrypted form.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_getDecryptedApikeyCmd).Standalone()

	wafv2_getDecryptedApikeyCmd.Flags().String("apikey", "", "The encrypted API key.")
	wafv2_getDecryptedApikeyCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_getDecryptedApikeyCmd.MarkFlagRequired("apikey")
	wafv2_getDecryptedApikeyCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_getDecryptedApikeyCmd)
}
