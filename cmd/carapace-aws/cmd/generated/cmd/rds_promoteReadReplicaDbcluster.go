package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_promoteReadReplicaDbclusterCmd = &cobra.Command{
	Use:   "promote-read-replica-dbcluster",
	Short: "Promotes a read replica DB cluster to a standalone DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_promoteReadReplicaDbclusterCmd).Standalone()

	rds_promoteReadReplicaDbclusterCmd.Flags().String("dbcluster-identifier", "", "The identifier of the DB cluster read replica to promote.")
	rds_promoteReadReplicaDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	rdsCmd.AddCommand(rds_promoteReadReplicaDbclusterCmd)
}
