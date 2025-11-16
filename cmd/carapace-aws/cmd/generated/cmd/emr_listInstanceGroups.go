package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_listInstanceGroupsCmd = &cobra.Command{
	Use:   "list-instance-groups",
	Short: "Provides all available details about the instance groups in a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_listInstanceGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_listInstanceGroupsCmd).Standalone()

		emr_listInstanceGroupsCmd.Flags().String("cluster-id", "", "The identifier of the cluster for which to list the instance groups.")
		emr_listInstanceGroupsCmd.Flags().String("marker", "", "The pagination token that indicates the next set of results to retrieve.")
		emr_listInstanceGroupsCmd.MarkFlagRequired("cluster-id")
	})
	emrCmd.AddCommand(emr_listInstanceGroupsCmd)
}
