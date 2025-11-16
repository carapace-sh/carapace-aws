package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createClusterSubnetGroupCmd = &cobra.Command{
	Use:   "create-cluster-subnet-group",
	Short: "Creates a new Amazon Redshift subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createClusterSubnetGroupCmd).Standalone()

	redshift_createClusterSubnetGroupCmd.Flags().String("cluster-subnet-group-name", "", "The name for the subnet group.")
	redshift_createClusterSubnetGroupCmd.Flags().String("description", "", "A description for the subnet group.")
	redshift_createClusterSubnetGroupCmd.Flags().String("subnet-ids", "", "An array of VPC subnet IDs.")
	redshift_createClusterSubnetGroupCmd.Flags().String("tags", "", "A list of tag instances.")
	redshift_createClusterSubnetGroupCmd.MarkFlagRequired("cluster-subnet-group-name")
	redshift_createClusterSubnetGroupCmd.MarkFlagRequired("description")
	redshift_createClusterSubnetGroupCmd.MarkFlagRequired("subnet-ids")
	redshiftCmd.AddCommand(redshift_createClusterSubnetGroupCmd)
}
