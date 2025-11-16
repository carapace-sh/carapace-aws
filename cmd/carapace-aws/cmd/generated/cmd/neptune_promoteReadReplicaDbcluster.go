package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_promoteReadReplicaDbclusterCmd = &cobra.Command{
	Use:   "promote-read-replica-dbcluster",
	Short: "Not supported.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_promoteReadReplicaDbclusterCmd).Standalone()

	neptune_promoteReadReplicaDbclusterCmd.Flags().String("dbcluster-identifier", "", "Not supported.")
	neptune_promoteReadReplicaDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	neptuneCmd.AddCommand(neptune_promoteReadReplicaDbclusterCmd)
}
