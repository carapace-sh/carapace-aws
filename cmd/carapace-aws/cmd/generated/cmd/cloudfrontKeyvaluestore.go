package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfrontKeyvaluestoreCmd = &cobra.Command{
	Use:   "cloudfront-keyvaluestore",
	Short: "Amazon CloudFront KeyValueStore Service to View and Update Data in a KVS Resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfrontKeyvaluestoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfrontKeyvaluestoreCmd).Standalone()

	})
	rootCmd.AddCommand(cloudfrontKeyvaluestoreCmd)
}
