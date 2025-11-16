package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_listStreamProcessorsCmd = &cobra.Command{
	Use:   "list-stream-processors",
	Short: "Gets a list of stream processors that you have created with [CreateStreamProcessor]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_listStreamProcessorsCmd).Standalone()

	rekognition_listStreamProcessorsCmd.Flags().String("max-results", "", "Maximum number of stream processors you want Amazon Rekognition Video to return in the response.")
	rekognition_listStreamProcessorsCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there are more stream processors to retrieve), Amazon Rekognition Video returns a pagination token in the response.")
	rekognitionCmd.AddCommand(rekognition_listStreamProcessorsCmd)
}
