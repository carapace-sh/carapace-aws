package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listMetricsCmd = &cobra.Command{
	Use:   "list-metrics",
	Short: "Lists the metrics that can be exported.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listMetricsCmd).Standalone()

	resiliencehub_listMetricsCmd.Flags().String("conditions", "", "Indicates the list of all the conditions that were applied on the metrics.")
	resiliencehub_listMetricsCmd.Flags().String("data-source", "", "Indicates the data source of the metrics.")
	resiliencehub_listMetricsCmd.Flags().String("fields", "", "Indicates the list of fields in the data source.")
	resiliencehub_listMetricsCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
	resiliencehub_listMetricsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
	resiliencehub_listMetricsCmd.Flags().String("sorts", "", "(Optional) Indicates the order in which you want to sort the fields in the metrics.")
	resiliencehubCmd.AddCommand(resiliencehub_listMetricsCmd)
}
