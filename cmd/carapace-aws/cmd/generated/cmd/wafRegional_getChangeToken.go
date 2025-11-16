package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getChangeTokenCmd = &cobra.Command{
	Use:   "get-change-token",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getChangeTokenCmd).Standalone()

	wafRegionalCmd.AddCommand(wafRegional_getChangeTokenCmd)
}
