package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_describeJobCmd = &cobra.Command{
	Use:   "describe-job",
	Short: "Returns information about a specific job including shipping information, job status, and other important metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_describeJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowball_describeJobCmd).Standalone()

		snowball_describeJobCmd.Flags().String("job-id", "", "The automatically generated ID for a job, for example `JID123e4567-e89b-12d3-a456-426655440000`.")
		snowball_describeJobCmd.MarkFlagRequired("job-id")
	})
	snowballCmd.AddCommand(snowball_describeJobCmd)
}
