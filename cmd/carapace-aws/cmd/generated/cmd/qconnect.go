package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnectCmd = &cobra.Command{
	Use:   "qconnect",
	Short: "- [Amazon Q actions](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations_Amazon_Q_Connect.html)\n- [Amazon Q data types](https://docs.aws.amazon.com/connect/latest/APIReference/API_Types_Amazon_Q_Connect.html)\n\n**Powered by Amazon Bedrock**: Amazon Web Services implements [automated abuse detection](https://docs.aws.amazon.com/bedrock/latest/userguide/abuse-detection.html). Because Amazon Q in Connect is built on Amazon Bedrock, users can take full advantage of the controls implemented in Amazon Bedrock to enforce safety, security, and the responsible use of artificial intelligence (AI).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnectCmd).Standalone()

	rootCmd.AddCommand(qconnectCmd)
}
