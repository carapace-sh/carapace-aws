package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createClusterParameterGroupCmd = &cobra.Command{
	Use:   "create-cluster-parameter-group",
	Short: "Creates an Amazon Redshift parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createClusterParameterGroupCmd).Standalone()

	redshift_createClusterParameterGroupCmd.Flags().String("description", "", "A description of the parameter group.")
	redshift_createClusterParameterGroupCmd.Flags().String("parameter-group-family", "", "The Amazon Redshift engine version to which the cluster parameter group applies.")
	redshift_createClusterParameterGroupCmd.Flags().String("parameter-group-name", "", "The name of the cluster parameter group.")
	redshift_createClusterParameterGroupCmd.Flags().String("tags", "", "A list of tag instances.")
	redshift_createClusterParameterGroupCmd.MarkFlagRequired("description")
	redshift_createClusterParameterGroupCmd.MarkFlagRequired("parameter-group-family")
	redshift_createClusterParameterGroupCmd.MarkFlagRequired("parameter-group-name")
	redshiftCmd.AddCommand(redshift_createClusterParameterGroupCmd)
}
