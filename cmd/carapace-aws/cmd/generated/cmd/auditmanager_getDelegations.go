package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getDelegationsCmd = &cobra.Command{
	Use:   "get-delegations",
	Short: "Gets a list of delegations from an audit owner to a delegate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getDelegationsCmd).Standalone()

	auditmanager_getDelegationsCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
	auditmanager_getDelegationsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	auditmanagerCmd.AddCommand(auditmanager_getDelegationsCmd)
}
