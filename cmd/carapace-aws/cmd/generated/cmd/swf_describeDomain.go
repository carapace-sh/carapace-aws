package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_describeDomainCmd = &cobra.Command{
	Use:   "describe-domain",
	Short: "Returns information about the specified domain, including description and status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_describeDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_describeDomainCmd).Standalone()

		swf_describeDomainCmd.Flags().String("name", "", "The name of the domain to describe.")
		swf_describeDomainCmd.MarkFlagRequired("name")
	})
	swfCmd.AddCommand(swf_describeDomainCmd)
}
