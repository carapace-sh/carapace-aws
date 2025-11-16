package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_stopDbclusterCmd = &cobra.Command{
	Use:   "stop-dbcluster",
	Short: "Stops an Amazon Aurora DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_stopDbclusterCmd).Standalone()

	rds_stopDbclusterCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier of the Amazon Aurora DB cluster to be stopped.")
	rds_stopDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	rdsCmd.AddCommand(rds_stopDbclusterCmd)
}
