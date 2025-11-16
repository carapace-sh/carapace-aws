package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafCmd = &cobra.Command{
	Use:   "waf",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafCmd).Standalone()

	})
	rootCmd.AddCommand(wafCmd)
}
