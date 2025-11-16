package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_listDomainsCmd = &cobra.Command{
	Use:   "list-domains",
	Short: "Lists all cases domains in the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_listDomainsCmd).Standalone()

	connectcases_listDomainsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connectcases_listDomainsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connectcasesCmd.AddCommand(connectcases_listDomainsCmd)
}
