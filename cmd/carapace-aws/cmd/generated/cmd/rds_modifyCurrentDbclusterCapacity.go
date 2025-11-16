package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyCurrentDbclusterCapacityCmd = &cobra.Command{
	Use:   "modify-current-dbcluster-capacity",
	Short: "Set the capacity of an Aurora Serverless v1 DB cluster to a specific value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyCurrentDbclusterCapacityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_modifyCurrentDbclusterCapacityCmd).Standalone()

		rds_modifyCurrentDbclusterCapacityCmd.Flags().String("capacity", "", "The DB cluster capacity.")
		rds_modifyCurrentDbclusterCapacityCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier for the cluster being modified.")
		rds_modifyCurrentDbclusterCapacityCmd.Flags().String("seconds-before-timeout", "", "The amount of time, in seconds, that Aurora Serverless v1 tries to find a scaling point to perform seamless scaling before enforcing the timeout action.")
		rds_modifyCurrentDbclusterCapacityCmd.Flags().String("timeout-action", "", "The action to take when the timeout is reached, either `ForceApplyCapacityChange` or `RollbackCapacityChange`.")
		rds_modifyCurrentDbclusterCapacityCmd.MarkFlagRequired("dbcluster-identifier")
	})
	rdsCmd.AddCommand(rds_modifyCurrentDbclusterCapacityCmd)
}
