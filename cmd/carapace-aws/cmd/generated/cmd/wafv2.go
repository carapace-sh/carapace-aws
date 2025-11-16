package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2Cmd = &cobra.Command{
	Use:   "wafv2",
	Short: "WAF\n\nThis is the latest version of the **WAF** API, released in November, 2019. The names of the entities that you use to access this API, like endpoints and namespaces, all have the versioning information added, like \"V2\" or \"v2\", to distinguish from the prior version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2Cmd).Standalone()

	})
	rootCmd.AddCommand(wafv2Cmd)
}
