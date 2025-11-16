package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_batchGetGraphMemberDatasourcesCmd = &cobra.Command{
	Use:   "batch-get-graph-member-datasources",
	Short: "Gets data source package information for the behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_batchGetGraphMemberDatasourcesCmd).Standalone()

	detective_batchGetGraphMemberDatasourcesCmd.Flags().String("account-ids", "", "The list of Amazon Web Services accounts to get data source package information on.")
	detective_batchGetGraphMemberDatasourcesCmd.Flags().String("graph-arn", "", "The ARN of the behavior graph.")
	detective_batchGetGraphMemberDatasourcesCmd.MarkFlagRequired("account-ids")
	detective_batchGetGraphMemberDatasourcesCmd.MarkFlagRequired("graph-arn")
	detectiveCmd.AddCommand(detective_batchGetGraphMemberDatasourcesCmd)
}
