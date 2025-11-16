package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getChangeTokenCmd = &cobra.Command{
	Use:   "get-change-token",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getChangeTokenCmd).Standalone()

	wafCmd.AddCommand(waf_getChangeTokenCmd)
}
