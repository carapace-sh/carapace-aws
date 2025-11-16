package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_listInstancesCmd = &cobra.Command{
	Use:   "list-instances",
	Short: "Provides information for all active Amazon EC2 instances and Amazon EC2 instances terminated in the last 30 days, up to a maximum of 2,000.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_listInstancesCmd).Standalone()

	emr_listInstancesCmd.Flags().String("cluster-id", "", "The identifier of the cluster for which to list the instances.")
	emr_listInstancesCmd.Flags().String("instance-fleet-id", "", "The unique identifier of the instance fleet.")
	emr_listInstancesCmd.Flags().String("instance-fleet-type", "", "The node type of the instance fleet.")
	emr_listInstancesCmd.Flags().String("instance-group-id", "", "The identifier of the instance group for which to list the instances.")
	emr_listInstancesCmd.Flags().String("instance-group-types", "", "The type of instance group for which to list the instances.")
	emr_listInstancesCmd.Flags().String("instance-states", "", "A list of instance states that will filter the instances returned with this request.")
	emr_listInstancesCmd.Flags().String("marker", "", "The pagination token that indicates the next set of results to retrieve.")
	emr_listInstancesCmd.MarkFlagRequired("cluster-id")
	emrCmd.AddCommand(emr_listInstancesCmd)
}
