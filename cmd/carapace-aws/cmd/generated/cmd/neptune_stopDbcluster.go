package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_stopDbclusterCmd = &cobra.Command{
	Use:   "stop-dbcluster",
	Short: "Stops an Amazon Neptune DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_stopDbclusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_stopDbclusterCmd).Standalone()

		neptune_stopDbclusterCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier of the Neptune DB cluster to be stopped.")
		neptune_stopDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	})
	neptuneCmd.AddCommand(neptune_stopDbclusterCmd)
}
