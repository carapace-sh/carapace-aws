package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_backtrackDbclusterCmd = &cobra.Command{
	Use:   "backtrack-dbcluster",
	Short: "Backtracks a DB cluster to a specific time, without creating a new DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_backtrackDbclusterCmd).Standalone()

	rds_backtrackDbclusterCmd.Flags().String("backtrack-to", "", "The timestamp of the time to backtrack the DB cluster to, specified in ISO 8601 format.")
	rds_backtrackDbclusterCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier of the DB cluster to be backtracked.")
	rds_backtrackDbclusterCmd.Flags().String("force", "", "Specifies whether to force the DB cluster to backtrack when binary logging is enabled.")
	rds_backtrackDbclusterCmd.Flags().String("use-earliest-time-on-point-in-time-unavailable", "", "Specifies whether to backtrack the DB cluster to the earliest possible backtrack time when *BacktrackTo* is set to a timestamp earlier than the earliest backtrack time.")
	rds_backtrackDbclusterCmd.MarkFlagRequired("backtrack-to")
	rds_backtrackDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	rdsCmd.AddCommand(rds_backtrackDbclusterCmd)
}
