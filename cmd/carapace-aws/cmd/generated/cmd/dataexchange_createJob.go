package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "This operation creates a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_createJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_createJobCmd).Standalone()

		dataexchange_createJobCmd.Flags().String("details", "", "The details for the CreateJob request.")
		dataexchange_createJobCmd.Flags().String("type", "", "The type of job to be created.")
		dataexchange_createJobCmd.MarkFlagRequired("details")
		dataexchange_createJobCmd.MarkFlagRequired("type")
	})
	dataexchangeCmd.AddCommand(dataexchange_createJobCmd)
}
