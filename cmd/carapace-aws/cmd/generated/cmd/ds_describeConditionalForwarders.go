package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeConditionalForwardersCmd = &cobra.Command{
	Use:   "describe-conditional-forwarders",
	Short: "Obtains information about the conditional forwarders for this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeConditionalForwardersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_describeConditionalForwardersCmd).Standalone()

		ds_describeConditionalForwardersCmd.Flags().String("directory-id", "", "The directory ID for which to get the list of associated conditional forwarders.")
		ds_describeConditionalForwardersCmd.Flags().String("remote-domain-names", "", "The fully qualified domain names (FQDN) of the remote domains for which to get the list of associated conditional forwarders.")
		ds_describeConditionalForwardersCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_describeConditionalForwardersCmd)
}
