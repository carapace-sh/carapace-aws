package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_modifyGlobalClusterCmd = &cobra.Command{
	Use:   "modify-global-cluster",
	Short: "Modify a setting for an Amazon Neptune global cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_modifyGlobalClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_modifyGlobalClusterCmd).Standalone()

		neptune_modifyGlobalClusterCmd.Flags().String("allow-major-version-upgrade", "", "A value that indicates whether major version upgrades are allowed.")
		neptune_modifyGlobalClusterCmd.Flags().String("deletion-protection", "", "Indicates whether the global database has deletion protection enabled.")
		neptune_modifyGlobalClusterCmd.Flags().String("engine-version", "", "The version number of the database engine to which you want to upgrade.")
		neptune_modifyGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The DB cluster identifier for the global cluster being modified.")
		neptune_modifyGlobalClusterCmd.Flags().String("new-global-cluster-identifier", "", "A new cluster identifier to assign to the global database.")
		neptune_modifyGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	})
	neptuneCmd.AddCommand(neptune_modifyGlobalClusterCmd)
}
