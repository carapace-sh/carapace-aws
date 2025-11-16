package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_listKxScalingGroupsCmd = &cobra.Command{
	Use:   "list-kx-scaling-groups",
	Short: "Returns a list of scaling groups in a kdb environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_listKxScalingGroupsCmd).Standalone()

	finspace_listKxScalingGroupsCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment, for which you want to retrieve a list of scaling groups.")
	finspace_listKxScalingGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	finspace_listKxScalingGroupsCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	finspace_listKxScalingGroupsCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_listKxScalingGroupsCmd)
}
