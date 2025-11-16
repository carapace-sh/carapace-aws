package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_rebootDbclusterCmd = &cobra.Command{
	Use:   "reboot-dbcluster",
	Short: "You might need to reboot your DB cluster, usually for maintenance reasons.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_rebootDbclusterCmd).Standalone()

	rds_rebootDbclusterCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier.")
	rds_rebootDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	rdsCmd.AddCommand(rds_rebootDbclusterCmd)
}
