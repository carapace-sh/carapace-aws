package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_resizeClusterCmd = &cobra.Command{
	Use:   "resize-cluster",
	Short: "Changes the size of the cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_resizeClusterCmd).Standalone()

	redshift_resizeClusterCmd.Flags().String("classic", "", "A boolean value indicating whether the resize operation is using the classic resize process.")
	redshift_resizeClusterCmd.Flags().String("cluster-identifier", "", "The unique identifier for the cluster to resize.")
	redshift_resizeClusterCmd.Flags().String("cluster-type", "", "The new cluster type for the specified cluster.")
	redshift_resizeClusterCmd.Flags().String("node-type", "", "The new node type for the nodes you are adding.")
	redshift_resizeClusterCmd.Flags().String("number-of-nodes", "", "The new number of nodes for the cluster.")
	redshift_resizeClusterCmd.Flags().String("reserved-node-id", "", "The identifier of the reserved node.")
	redshift_resizeClusterCmd.Flags().String("target-reserved-node-offering-id", "", "The identifier of the target reserved node offering.")
	redshift_resizeClusterCmd.MarkFlagRequired("cluster-identifier")
	redshiftCmd.AddCommand(redshift_resizeClusterCmd)
}
