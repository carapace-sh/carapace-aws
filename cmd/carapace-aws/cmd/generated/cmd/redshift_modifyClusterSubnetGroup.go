package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyClusterSubnetGroupCmd = &cobra.Command{
	Use:   "modify-cluster-subnet-group",
	Short: "Modifies a cluster subnet group to include the specified list of VPC subnets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyClusterSubnetGroupCmd).Standalone()

	redshift_modifyClusterSubnetGroupCmd.Flags().String("cluster-subnet-group-name", "", "The name of the subnet group to be modified.")
	redshift_modifyClusterSubnetGroupCmd.Flags().String("description", "", "A text description of the subnet group to be modified.")
	redshift_modifyClusterSubnetGroupCmd.Flags().String("subnet-ids", "", "An array of VPC subnet IDs.")
	redshift_modifyClusterSubnetGroupCmd.MarkFlagRequired("cluster-subnet-group-name")
	redshift_modifyClusterSubnetGroupCmd.MarkFlagRequired("subnet-ids")
	redshiftCmd.AddCommand(redshift_modifyClusterSubnetGroupCmd)
}
