package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_modifyInstanceGroupsCmd = &cobra.Command{
	Use:   "modify-instance-groups",
	Short: "ModifyInstanceGroups modifies the number of nodes and configuration settings of an instance group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_modifyInstanceGroupsCmd).Standalone()

	emr_modifyInstanceGroupsCmd.Flags().String("cluster-id", "", "The ID of the cluster to which the instance group belongs.")
	emr_modifyInstanceGroupsCmd.Flags().String("instance-groups", "", "Instance groups to change.")
	emrCmd.AddCommand(emr_modifyInstanceGroupsCmd)
}
