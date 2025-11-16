package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listClusterEventsCmd = &cobra.Command{
	Use:   "list-cluster-events",
	Short: "Retrieves a list of event summaries for a specified HyperPod cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listClusterEventsCmd).Standalone()

	sagemaker_listClusterEventsCmd.Flags().String("cluster-name", "", "The name or Amazon Resource Name (ARN) of the HyperPod cluster for which to list events.")
	sagemaker_listClusterEventsCmd.Flags().String("event-time-after", "", "The start of the time range for filtering events.")
	sagemaker_listClusterEventsCmd.Flags().String("event-time-before", "", "The end of the time range for filtering events.")
	sagemaker_listClusterEventsCmd.Flags().String("instance-group-name", "", "The name of the instance group to filter events.")
	sagemaker_listClusterEventsCmd.Flags().String("max-results", "", "The maximum number of events to return in the response.")
	sagemaker_listClusterEventsCmd.Flags().String("next-token", "", "A token to retrieve the next set of results.")
	sagemaker_listClusterEventsCmd.Flags().String("node-id", "", "The EC2 instance ID to filter events.")
	sagemaker_listClusterEventsCmd.Flags().String("resource-type", "", "The type of resource for which to filter events.")
	sagemaker_listClusterEventsCmd.Flags().String("sort-by", "", "The field to use for sorting the event list.")
	sagemaker_listClusterEventsCmd.Flags().String("sort-order", "", "The order in which to sort the results.")
	sagemaker_listClusterEventsCmd.MarkFlagRequired("cluster-name")
	sagemakerCmd.AddCommand(sagemaker_listClusterEventsCmd)
}
