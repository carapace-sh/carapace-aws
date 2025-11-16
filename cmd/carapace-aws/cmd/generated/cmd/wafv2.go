package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2Cmd = &cobra.Command{
	Use:   "wafv2",
	Short: "WAF",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2Cmd).Standalone()

	})
	rootCmd.AddCommand(wafv2Cmd)
}
