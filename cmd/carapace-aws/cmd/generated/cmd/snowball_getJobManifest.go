package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_getJobManifestCmd = &cobra.Command{
	Use:   "get-job-manifest",
	Short: "Returns a link to an Amazon S3 presigned URL for the manifest file associated with the specified `JobId` value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_getJobManifestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowball_getJobManifestCmd).Standalone()

		snowball_getJobManifestCmd.Flags().String("job-id", "", "The ID for a job that you want to get the manifest file for, for example `JID123e4567-e89b-12d3-a456-426655440000`.")
		snowball_getJobManifestCmd.MarkFlagRequired("job-id")
	})
	snowballCmd.AddCommand(snowball_getJobManifestCmd)
}
