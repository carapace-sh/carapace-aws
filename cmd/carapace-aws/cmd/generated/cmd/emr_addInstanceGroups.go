package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_addInstanceGroupsCmd = &cobra.Command{
	Use:   "add-instance-groups",
	Short: "Adds one or more instance groups to a running cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_addInstanceGroupsCmd).Standalone()

	emr_addInstanceGroupsCmd.Flags().String("instance-groups", "", "Instance groups to add.")
	emr_addInstanceGroupsCmd.Flags().String("job-flow-id", "", "Job flow in which to add the instance groups.")
	emr_addInstanceGroupsCmd.MarkFlagRequired("instance-groups")
	emr_addInstanceGroupsCmd.MarkFlagRequired("job-flow-id")
	emrCmd.AddCommand(emr_addInstanceGroupsCmd)
}
