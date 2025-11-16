package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_listDomainsCmd = &cobra.Command{
	Use:   "list-domains",
	Short: "Lists all the domains in the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_listDomainsCmd).Standalone()

	voiceId_listDomainsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	voiceId_listDomainsCmd.Flags().String("next-token", "", "If `NextToken` is returned, there are more results available.")
	voiceIdCmd.AddCommand(voiceId_listDomainsCmd)
}
