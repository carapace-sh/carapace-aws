package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_deleteJobCmd = &cobra.Command{
	Use:   "delete-job",
	Short: "Deletes a single Job by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_deleteJobCmd).Standalone()

	drs_deleteJobCmd.Flags().String("job-id", "", "The ID of the Job to be deleted.")
	drs_deleteJobCmd.MarkFlagRequired("job-id")
	drsCmd.AddCommand(drs_deleteJobCmd)
}
