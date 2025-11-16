package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_updateClusterCmd = &cobra.Command{
	Use:   "update-cluster",
	Short: "While a cluster's `ClusterState` value is in the `AwaitingQuorum` state, you can update some of the information associated with a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_updateClusterCmd).Standalone()

	snowball_updateClusterCmd.Flags().String("address-id", "", "The ID of the updated [Address]() object.")
	snowball_updateClusterCmd.Flags().String("cluster-id", "", "The cluster ID of the cluster that you want to update, for example `CID123e4567-e89b-12d3-a456-426655440000`.")
	snowball_updateClusterCmd.Flags().String("description", "", "The updated description of this cluster.")
	snowball_updateClusterCmd.Flags().String("forwarding-address-id", "", "The updated ID for the forwarding address for a cluster.")
	snowball_updateClusterCmd.Flags().String("notification", "", "The new or updated [Notification]() object.")
	snowball_updateClusterCmd.Flags().String("on-device-service-configuration", "", "Specifies the service or services on the Snow Family device that your transferred data will be exported from or imported into.")
	snowball_updateClusterCmd.Flags().String("resources", "", "The updated arrays of [JobResource]() objects that can include updated [S3Resource]() objects or [LambdaResource]() objects.")
	snowball_updateClusterCmd.Flags().String("role-arn", "", "The new role Amazon Resource Name (ARN) that you want to associate with this cluster.")
	snowball_updateClusterCmd.Flags().String("shipping-option", "", "The updated shipping option value of this cluster's [ShippingDetails]() object.")
	snowball_updateClusterCmd.MarkFlagRequired("cluster-id")
	snowballCmd.AddCommand(snowball_updateClusterCmd)
}
