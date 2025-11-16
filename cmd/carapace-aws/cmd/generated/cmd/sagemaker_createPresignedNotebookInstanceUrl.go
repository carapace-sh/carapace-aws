package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createPresignedNotebookInstanceUrlCmd = &cobra.Command{
	Use:   "create-presigned-notebook-instance-url",
	Short: "Returns a URL that you can use to connect to the Jupyter server from a notebook instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createPresignedNotebookInstanceUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createPresignedNotebookInstanceUrlCmd).Standalone()

		sagemaker_createPresignedNotebookInstanceUrlCmd.Flags().String("notebook-instance-name", "", "The name of the notebook instance.")
		sagemaker_createPresignedNotebookInstanceUrlCmd.Flags().String("session-expiration-duration-in-seconds", "", "The duration of the session, in seconds.")
		sagemaker_createPresignedNotebookInstanceUrlCmd.MarkFlagRequired("notebook-instance-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createPresignedNotebookInstanceUrlCmd)
}
