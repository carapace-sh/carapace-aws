package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_getJobUnlockCodeCmd = &cobra.Command{
	Use:   "get-job-unlock-code",
	Short: "Returns the `UnlockCode` code value for the specified job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_getJobUnlockCodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowball_getJobUnlockCodeCmd).Standalone()

		snowball_getJobUnlockCodeCmd.Flags().String("job-id", "", "The ID for the job that you want to get the `UnlockCode` value for, for example `JID123e4567-e89b-12d3-a456-426655440000`.")
		snowball_getJobUnlockCodeCmd.MarkFlagRequired("job-id")
	})
	snowballCmd.AddCommand(snowball_getJobUnlockCodeCmd)
}
