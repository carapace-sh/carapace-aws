package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegionalCmd = &cobra.Command{
	Use:   "waf-regional",
	Short: "This is **AWS WAF Classic Regional** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegionalCmd).Standalone()

	rootCmd.AddCommand(wafRegionalCmd)
}
