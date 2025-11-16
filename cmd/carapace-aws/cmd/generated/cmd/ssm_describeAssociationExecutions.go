package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeAssociationExecutionsCmd = &cobra.Command{
	Use:   "describe-association-executions",
	Short: "Views all executions for a specific association ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeAssociationExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeAssociationExecutionsCmd).Standalone()

		ssm_describeAssociationExecutionsCmd.Flags().String("association-id", "", "The association ID for which you want to view execution history details.")
		ssm_describeAssociationExecutionsCmd.Flags().String("filters", "", "Filters for the request.")
		ssm_describeAssociationExecutionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_describeAssociationExecutionsCmd.Flags().String("next-token", "", "A token to start the list.")
		ssm_describeAssociationExecutionsCmd.MarkFlagRequired("association-id")
	})
	ssmCmd.AddCommand(ssm_describeAssociationExecutionsCmd)
}
