package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteClusterParameterGroupCmd = &cobra.Command{
	Use:   "delete-cluster-parameter-group",
	Short: "Deletes a specified Amazon Redshift parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteClusterParameterGroupCmd).Standalone()

	redshift_deleteClusterParameterGroupCmd.Flags().String("parameter-group-name", "", "The name of the parameter group to be deleted.")
	redshift_deleteClusterParameterGroupCmd.MarkFlagRequired("parameter-group-name")
	redshiftCmd.AddCommand(redshift_deleteClusterParameterGroupCmd)
}
