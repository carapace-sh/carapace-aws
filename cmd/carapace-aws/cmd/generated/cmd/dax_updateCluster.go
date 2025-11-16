package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_updateClusterCmd = &cobra.Command{
	Use:   "update-cluster",
	Short: "Modifies the settings for a DAX cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_updateClusterCmd).Standalone()

	dax_updateClusterCmd.Flags().String("cluster-name", "", "The name of the DAX cluster to be modified.")
	dax_updateClusterCmd.Flags().String("description", "", "A description of the changes being made to the cluster.")
	dax_updateClusterCmd.Flags().String("notification-topic-arn", "", "The Amazon Resource Name (ARN) that identifies the topic.")
	dax_updateClusterCmd.Flags().String("notification-topic-status", "", "The current state of the topic.")
	dax_updateClusterCmd.Flags().String("parameter-group-name", "", "The name of a parameter group for this cluster.")
	dax_updateClusterCmd.Flags().String("preferred-maintenance-window", "", "A range of time when maintenance of DAX cluster software will be performed.")
	dax_updateClusterCmd.Flags().String("security-group-ids", "", "A list of user-specified security group IDs to be assigned to each node in the DAX cluster.")
	dax_updateClusterCmd.MarkFlagRequired("cluster-name")
	daxCmd.AddCommand(dax_updateClusterCmd)
}
