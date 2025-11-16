package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_createApplicationVersionCmd = &cobra.Command{
	Use:   "create-application-version",
	Short: "Creates an application version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_createApplicationVersionCmd).Standalone()

	serverlessrepo_createApplicationVersionCmd.Flags().String("application-id", "", "The Amazon Resource Name (ARN) of the application.")
	serverlessrepo_createApplicationVersionCmd.Flags().String("semantic-version", "", "The semantic version of the new version.")
	serverlessrepo_createApplicationVersionCmd.Flags().String("source-code-archive-url", "", "A link to the S3 object that contains the ZIP archive of the source code for this version of your application.")
	serverlessrepo_createApplicationVersionCmd.Flags().String("source-code-url", "", "A link to a public repository for the source code of your application, for example the URL of a specific GitHub commit.")
	serverlessrepo_createApplicationVersionCmd.Flags().String("template-body", "", "The raw packaged AWS SAM template of your application.")
	serverlessrepo_createApplicationVersionCmd.Flags().String("template-url", "", "A link to the packaged AWS SAM template of your application.")
	serverlessrepo_createApplicationVersionCmd.MarkFlagRequired("application-id")
	serverlessrepo_createApplicationVersionCmd.MarkFlagRequired("semantic-version")
	serverlessrepoCmd.AddCommand(serverlessrepo_createApplicationVersionCmd)
}
