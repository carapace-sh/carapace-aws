package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_batchGetMembershipDatasourcesCmd = &cobra.Command{
	Use:   "batch-get-membership-datasources",
	Short: "Gets information on the data source package history for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_batchGetMembershipDatasourcesCmd).Standalone()

	detective_batchGetMembershipDatasourcesCmd.Flags().String("graph-arns", "", "The ARN of the behavior graph.")
	detective_batchGetMembershipDatasourcesCmd.MarkFlagRequired("graph-arns")
	detectiveCmd.AddCommand(detective_batchGetMembershipDatasourcesCmd)
}
