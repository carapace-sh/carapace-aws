package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_startReadSetActivationJobCmd = &cobra.Command{
	Use:   "start-read-set-activation-job",
	Short: "Activates an archived read set and returns its metadata in a JSON formatted output.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_startReadSetActivationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_startReadSetActivationJobCmd).Standalone()

		omics_startReadSetActivationJobCmd.Flags().String("client-token", "", "To ensure that jobs don't run multiple times, specify a unique token for each job.")
		omics_startReadSetActivationJobCmd.Flags().String("sequence-store-id", "", "The read set's sequence store ID.")
		omics_startReadSetActivationJobCmd.Flags().String("sources", "", "The job's source files.")
		omics_startReadSetActivationJobCmd.MarkFlagRequired("sequence-store-id")
		omics_startReadSetActivationJobCmd.MarkFlagRequired("sources")
	})
	omicsCmd.AddCommand(omics_startReadSetActivationJobCmd)
}
