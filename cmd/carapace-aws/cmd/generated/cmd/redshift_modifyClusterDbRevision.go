package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyClusterDbRevisionCmd = &cobra.Command{
	Use:   "modify-cluster-db-revision",
	Short: "Modifies the database revision of a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyClusterDbRevisionCmd).Standalone()

	redshift_modifyClusterDbRevisionCmd.Flags().String("cluster-identifier", "", "The unique identifier of a cluster whose database revision you want to modify.")
	redshift_modifyClusterDbRevisionCmd.Flags().String("revision-target", "", "The identifier of the database revision.")
	redshift_modifyClusterDbRevisionCmd.MarkFlagRequired("cluster-identifier")
	redshift_modifyClusterDbRevisionCmd.MarkFlagRequired("revision-target")
	redshiftCmd.AddCommand(redshift_modifyClusterDbRevisionCmd)
}
