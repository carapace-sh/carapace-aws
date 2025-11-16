package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_requestUploadCredentialsCmd = &cobra.Command{
	Use:   "request-upload-credentials",
	Short: "**This API works with the following fleet types:** EC2\n\nRetrieves a fresh set of credentials for use when uploading a new set of game build files to Amazon GameLift Servers's Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_requestUploadCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_requestUploadCredentialsCmd).Standalone()

		gamelift_requestUploadCredentialsCmd.Flags().String("build-id", "", "A unique identifier for the build to get credentials for.")
		gamelift_requestUploadCredentialsCmd.MarkFlagRequired("build-id")
	})
	gameliftCmd.AddCommand(gamelift_requestUploadCredentialsCmd)
}
