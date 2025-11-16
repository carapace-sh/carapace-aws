package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_startJobCmd = &cobra.Command{
	Use:   "start-job",
	Short: "This operation starts a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_startJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_startJobCmd).Standalone()

		dataexchange_startJobCmd.Flags().String("job-id", "", "The unique identifier for a job.")
		dataexchange_startJobCmd.MarkFlagRequired("job-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_startJobCmd)
}
