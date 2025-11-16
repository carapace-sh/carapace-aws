package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_deprecateDomainCmd = &cobra.Command{
	Use:   "deprecate-domain",
	Short: "Deprecates the specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_deprecateDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_deprecateDomainCmd).Standalone()

		swf_deprecateDomainCmd.Flags().String("name", "", "The name of the domain to deprecate.")
		swf_deprecateDomainCmd.MarkFlagRequired("name")
	})
	swfCmd.AddCommand(swf_deprecateDomainCmd)
}
