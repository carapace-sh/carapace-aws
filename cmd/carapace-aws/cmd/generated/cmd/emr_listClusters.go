package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_listClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "Provides the status of all clusters visible to this Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_listClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_listClustersCmd).Standalone()

		emr_listClustersCmd.Flags().String("cluster-states", "", "The cluster state filters to apply when listing clusters.")
		emr_listClustersCmd.Flags().String("created-after", "", "The creation date and time beginning value filter for listing clusters.")
		emr_listClustersCmd.Flags().String("created-before", "", "The creation date and time end value filter for listing clusters.")
		emr_listClustersCmd.Flags().String("marker", "", "The pagination token that indicates the next set of results to retrieve.")
	})
	emrCmd.AddCommand(emr_listClustersCmd)
}
