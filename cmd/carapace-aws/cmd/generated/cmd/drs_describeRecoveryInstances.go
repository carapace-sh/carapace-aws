package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_describeRecoveryInstancesCmd = &cobra.Command{
	Use:   "describe-recovery-instances",
	Short: "Lists all Recovery Instances or multiple Recovery Instances by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_describeRecoveryInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_describeRecoveryInstancesCmd).Standalone()

		drs_describeRecoveryInstancesCmd.Flags().String("filters", "", "A set of filters by which to return Recovery Instances.")
		drs_describeRecoveryInstancesCmd.Flags().String("max-results", "", "Maximum number of Recovery Instances to retrieve.")
		drs_describeRecoveryInstancesCmd.Flags().String("next-token", "", "The token of the next Recovery Instance to retrieve.")
	})
	drsCmd.AddCommand(drs_describeRecoveryInstancesCmd)
}
