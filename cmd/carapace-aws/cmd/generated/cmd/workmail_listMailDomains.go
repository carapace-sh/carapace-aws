package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listMailDomainsCmd = &cobra.Command{
	Use:   "list-mail-domains",
	Short: "Lists the mail domains in a given WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listMailDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_listMailDomainsCmd).Standalone()

		workmail_listMailDomainsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		workmail_listMailDomainsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		workmail_listMailDomainsCmd.Flags().String("organization-id", "", "The WorkMail organization for which to list domains.")
		workmail_listMailDomainsCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_listMailDomainsCmd)
}
