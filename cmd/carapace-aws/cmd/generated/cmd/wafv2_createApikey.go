package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_createApikeyCmd = &cobra.Command{
	Use:   "create-apikey",
	Short: "Creates an API key that contains a set of token domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_createApikeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_createApikeyCmd).Standalone()

		wafv2_createApikeyCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_createApikeyCmd.Flags().String("token-domains", "", "The client application domains that you want to use this API key for.")
		wafv2_createApikeyCmd.MarkFlagRequired("scope")
		wafv2_createApikeyCmd.MarkFlagRequired("token-domains")
	})
	wafv2Cmd.AddCommand(wafv2_createApikeyCmd)
}
