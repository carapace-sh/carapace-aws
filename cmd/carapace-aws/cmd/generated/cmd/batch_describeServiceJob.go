package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_describeServiceJobCmd = &cobra.Command{
	Use:   "describe-service-job",
	Short: "The details of a service job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_describeServiceJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_describeServiceJobCmd).Standalone()

		batch_describeServiceJobCmd.Flags().String("job-id", "", "The job ID for the service job to describe.")
		batch_describeServiceJobCmd.MarkFlagRequired("job-id")
	})
	batchCmd.AddCommand(batch_describeServiceJobCmd)
}
