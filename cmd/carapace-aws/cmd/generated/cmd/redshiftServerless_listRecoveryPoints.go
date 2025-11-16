package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listRecoveryPointsCmd = &cobra.Command{
	Use:   "list-recovery-points",
	Short: "Returns an array of recovery points.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listRecoveryPointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_listRecoveryPointsCmd).Standalone()

		redshiftServerless_listRecoveryPointsCmd.Flags().String("end-time", "", "The time when creation of the recovery point finished.")
		redshiftServerless_listRecoveryPointsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		redshiftServerless_listRecoveryPointsCmd.Flags().String("namespace-arn", "", "The Amazon Resource Name (ARN) of the namespace from which to list recovery points.")
		redshiftServerless_listRecoveryPointsCmd.Flags().String("namespace-name", "", "The name of the namespace to list recovery points for.")
		redshiftServerless_listRecoveryPointsCmd.Flags().String("next-token", "", "If your initial `ListRecoveryPoints` operation returns a `nextToken`, you can include the returned `nextToken` in following `ListRecoveryPoints` operations, which returns results in the next page.")
		redshiftServerless_listRecoveryPointsCmd.Flags().String("start-time", "", "The time when the recovery point's creation was initiated.")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_listRecoveryPointsCmd)
}
