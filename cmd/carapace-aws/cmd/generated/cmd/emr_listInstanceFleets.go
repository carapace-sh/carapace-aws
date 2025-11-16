package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_listInstanceFleetsCmd = &cobra.Command{
	Use:   "list-instance-fleets",
	Short: "Lists all available details about the instance fleets in a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_listInstanceFleetsCmd).Standalone()

	emr_listInstanceFleetsCmd.Flags().String("cluster-id", "", "The unique identifier of the cluster.")
	emr_listInstanceFleetsCmd.Flags().String("marker", "", "The pagination token that indicates the next set of results to retrieve.")
	emr_listInstanceFleetsCmd.MarkFlagRequired("cluster-id")
	emrCmd.AddCommand(emr_listInstanceFleetsCmd)
}
