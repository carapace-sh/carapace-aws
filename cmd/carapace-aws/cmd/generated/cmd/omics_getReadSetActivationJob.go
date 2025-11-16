package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getReadSetActivationJobCmd = &cobra.Command{
	Use:   "get-read-set-activation-job",
	Short: "Returns detailed information about the status of a read set activation job in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getReadSetActivationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_getReadSetActivationJobCmd).Standalone()

		omics_getReadSetActivationJobCmd.Flags().String("id", "", "The job's ID.")
		omics_getReadSetActivationJobCmd.Flags().String("sequence-store-id", "", "The job's sequence store ID.")
		omics_getReadSetActivationJobCmd.MarkFlagRequired("id")
		omics_getReadSetActivationJobCmd.MarkFlagRequired("sequence-store-id")
	})
	omicsCmd.AddCommand(omics_getReadSetActivationJobCmd)
}
