package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listAssociationVersionsCmd = &cobra.Command{
	Use:   "list-association-versions",
	Short: "Retrieves all versions of an association for a specific association ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listAssociationVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listAssociationVersionsCmd).Standalone()

		ssm_listAssociationVersionsCmd.Flags().String("association-id", "", "The association ID for which you want to view all versions.")
		ssm_listAssociationVersionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_listAssociationVersionsCmd.Flags().String("next-token", "", "A token to start the list.")
		ssm_listAssociationVersionsCmd.MarkFlagRequired("association-id")
	})
	ssmCmd.AddCommand(ssm_listAssociationVersionsCmd)
}
