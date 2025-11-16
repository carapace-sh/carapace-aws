package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyGlobalClusterCmd = &cobra.Command{
	Use:   "modify-global-cluster",
	Short: "Modifies a setting for an Amazon Aurora global database cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyGlobalClusterCmd).Standalone()

	rds_modifyGlobalClusterCmd.Flags().String("allow-major-version-upgrade", "", "Specifies whether to allow major version upgrades.")
	rds_modifyGlobalClusterCmd.Flags().String("deletion-protection", "", "Specifies whether to enable deletion protection for the global database cluster.")
	rds_modifyGlobalClusterCmd.Flags().String("engine-version", "", "The version number of the database engine to which you want to upgrade.")
	rds_modifyGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The cluster identifier for the global cluster to modify.")
	rds_modifyGlobalClusterCmd.Flags().String("new-global-cluster-identifier", "", "The new cluster identifier for the global database cluster.")
	rds_modifyGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	rdsCmd.AddCommand(rds_modifyGlobalClusterCmd)
}
