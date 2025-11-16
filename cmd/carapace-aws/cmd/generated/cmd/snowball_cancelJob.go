package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_cancelJobCmd = &cobra.Command{
	Use:   "cancel-job",
	Short: "Cancels the specified job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_cancelJobCmd).Standalone()

	snowball_cancelJobCmd.Flags().String("job-id", "", "The 39-character job ID for the job that you want to cancel, for example `JID123e4567-e89b-12d3-a456-426655440000`.")
	snowball_cancelJobCmd.MarkFlagRequired("job-id")
	snowballCmd.AddCommand(snowball_cancelJobCmd)
}
