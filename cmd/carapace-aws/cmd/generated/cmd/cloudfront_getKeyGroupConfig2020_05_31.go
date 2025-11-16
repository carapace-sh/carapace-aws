package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getKeyGroupConfig2020_05_31Cmd = &cobra.Command{
	Use:   "get-key-group-config2020_05_31",
	Short: "Gets a key group configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getKeyGroupConfig2020_05_31Cmd).Standalone()

	cloudfront_getKeyGroupConfig2020_05_31Cmd.Flags().String("id", "", "The identifier of the key group whose configuration you are getting.")
	cloudfront_getKeyGroupConfig2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_getKeyGroupConfig2020_05_31Cmd)
}
