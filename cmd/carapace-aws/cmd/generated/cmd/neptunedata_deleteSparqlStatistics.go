package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_deleteSparqlStatisticsCmd = &cobra.Command{
	Use:   "delete-sparql-statistics",
	Short: "Deletes SPARQL statistics\n\nWhen invoking this operation in a Neptune cluster that has IAM authentication enabled, the IAM user or role making the request must have a policy attached that allows the [neptune-db:DeleteStatistics](https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#deletestatistics) IAM action in that cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_deleteSparqlStatisticsCmd).Standalone()

	neptunedataCmd.AddCommand(neptunedata_deleteSparqlStatisticsCmd)
}
