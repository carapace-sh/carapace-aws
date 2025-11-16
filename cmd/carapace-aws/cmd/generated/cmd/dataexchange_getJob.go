package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_getJobCmd = &cobra.Command{
	Use:   "get-job",
	Short: "This operation returns information about a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_getJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_getJobCmd).Standalone()

		dataexchange_getJobCmd.Flags().String("job-id", "", "The unique identifier for a job.")
		dataexchange_getJobCmd.MarkFlagRequired("job-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_getJobCmd)
}
