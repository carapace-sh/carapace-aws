package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteClusterSubnetGroupCmd = &cobra.Command{
	Use:   "delete-cluster-subnet-group",
	Short: "Deletes the specified cluster subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteClusterSubnetGroupCmd).Standalone()

	redshift_deleteClusterSubnetGroupCmd.Flags().String("cluster-subnet-group-name", "", "The name of the cluster subnet group name to be deleted.")
	redshift_deleteClusterSubnetGroupCmd.MarkFlagRequired("cluster-subnet-group-name")
	redshiftCmd.AddCommand(redshift_deleteClusterSubnetGroupCmd)
}
