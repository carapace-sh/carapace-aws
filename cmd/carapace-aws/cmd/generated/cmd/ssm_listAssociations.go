package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listAssociationsCmd = &cobra.Command{
	Use:   "list-associations",
	Short: "Returns all State Manager associations in the current Amazon Web Services account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listAssociationsCmd).Standalone()

		ssm_listAssociationsCmd.Flags().String("association-filter-list", "", "One or more filters.")
		ssm_listAssociationsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_listAssociationsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	ssmCmd.AddCommand(ssm_listAssociationsCmd)
}
