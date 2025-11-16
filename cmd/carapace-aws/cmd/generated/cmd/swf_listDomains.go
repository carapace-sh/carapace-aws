package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_listDomainsCmd = &cobra.Command{
	Use:   "list-domains",
	Short: "Returns the list of domains registered in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_listDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_listDomainsCmd).Standalone()

		swf_listDomainsCmd.Flags().String("maximum-page-size", "", "The maximum number of results that are returned per call.")
		swf_listDomainsCmd.Flags().String("next-page-token", "", "If `NextPageToken` is returned there are more results available.")
		swf_listDomainsCmd.Flags().String("registration-status", "", "Specifies the registration status of the domains to list.")
		swf_listDomainsCmd.Flags().String("reverse-order", "", "When set to `true`, returns the results in reverse order.")
		swf_listDomainsCmd.MarkFlagRequired("registration-status")
	})
	swfCmd.AddCommand(swf_listDomainsCmd)
}
