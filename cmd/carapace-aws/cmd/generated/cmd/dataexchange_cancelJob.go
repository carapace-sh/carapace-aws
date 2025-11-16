package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_cancelJobCmd = &cobra.Command{
	Use:   "cancel-job",
	Short: "This operation cancels a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_cancelJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_cancelJobCmd).Standalone()

		dataexchange_cancelJobCmd.Flags().String("job-id", "", "The unique identifier for a job.")
		dataexchange_cancelJobCmd.MarkFlagRequired("job-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_cancelJobCmd)
}
