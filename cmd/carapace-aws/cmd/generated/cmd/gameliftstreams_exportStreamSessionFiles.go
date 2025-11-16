package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_exportStreamSessionFilesCmd = &cobra.Command{
	Use:   "export-stream-session-files",
	Short: "Export the files that your application modifies or generates in a stream session, which can help you debug or verify your application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_exportStreamSessionFilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreams_exportStreamSessionFilesCmd).Standalone()

		gameliftstreams_exportStreamSessionFilesCmd.Flags().String("identifier", "", "An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the stream group resource.")
		gameliftstreams_exportStreamSessionFilesCmd.Flags().String("output-uri", "", "The S3 bucket URI where Amazon GameLift Streams uploads the set of compressed exported files for this stream session.")
		gameliftstreams_exportStreamSessionFilesCmd.Flags().String("stream-session-identifier", "", "An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the stream session resource.")
		gameliftstreams_exportStreamSessionFilesCmd.MarkFlagRequired("identifier")
		gameliftstreams_exportStreamSessionFilesCmd.MarkFlagRequired("output-uri")
		gameliftstreams_exportStreamSessionFilesCmd.MarkFlagRequired("stream-session-identifier")
	})
	gameliftstreamsCmd.AddCommand(gameliftstreams_exportStreamSessionFilesCmd)
}
