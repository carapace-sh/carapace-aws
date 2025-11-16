package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_getSoftwareUpdatesCmd = &cobra.Command{
	Use:   "get-software-updates",
	Short: "Returns an Amazon S3 presigned URL for an update file associated with a specified `JobId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_getSoftwareUpdatesCmd).Standalone()

	snowball_getSoftwareUpdatesCmd.Flags().String("job-id", "", "The ID for a job that you want to get the software update file for, for example `JID123e4567-e89b-12d3-a456-426655440000`.")
	snowball_getSoftwareUpdatesCmd.MarkFlagRequired("job-id")
	snowballCmd.AddCommand(snowball_getSoftwareUpdatesCmd)
}
