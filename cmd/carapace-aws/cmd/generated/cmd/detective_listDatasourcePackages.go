package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_listDatasourcePackagesCmd = &cobra.Command{
	Use:   "list-datasource-packages",
	Short: "Lists data source packages in the behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_listDatasourcePackagesCmd).Standalone()

	detective_listDatasourcePackagesCmd.Flags().String("graph-arn", "", "The ARN of the behavior graph.")
	detective_listDatasourcePackagesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	detective_listDatasourcePackagesCmd.Flags().String("next-token", "", "For requests to get the next page of results, the pagination token that was returned with the previous set of results.")
	detective_listDatasourcePackagesCmd.MarkFlagRequired("graph-arn")
	detectiveCmd.AddCommand(detective_listDatasourcePackagesCmd)
}
