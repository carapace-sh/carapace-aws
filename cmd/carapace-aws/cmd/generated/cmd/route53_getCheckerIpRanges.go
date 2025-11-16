package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getCheckerIpRangesCmd = &cobra.Command{
	Use:   "get-checker-ip-ranges",
	Short: "Route 53 does not perform authorization for this API because it retrieves information that is already available to the public.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getCheckerIpRangesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_getCheckerIpRangesCmd).Standalone()

	})
	route53Cmd.AddCommand(route53_getCheckerIpRangesCmd)
}
