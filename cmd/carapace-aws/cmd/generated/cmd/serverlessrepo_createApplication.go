package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates an application, optionally including an AWS SAM file to create the first application version in the same call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_createApplicationCmd).Standalone()

	serverlessrepo_createApplicationCmd.Flags().String("author", "", "The name of the author publishing the app.")
	serverlessrepo_createApplicationCmd.Flags().String("description", "", "The description of the application.")
	serverlessrepo_createApplicationCmd.Flags().String("home-page-url", "", "A URL with more information about the application, for example the location of your GitHub repository for the application.")
	serverlessrepo_createApplicationCmd.Flags().String("labels", "", "Labels to improve discovery of apps in search results.")
	serverlessrepo_createApplicationCmd.Flags().String("license-body", "", "A local text file that contains the license of the app that matches the spdxLicenseID value of your application.")
	serverlessrepo_createApplicationCmd.Flags().String("license-url", "", "A link to the S3 object that contains the license of the app that matches the spdxLicenseID value of your application.")
	serverlessrepo_createApplicationCmd.Flags().String("name", "", "The name of the application that you want to publish.")
	serverlessrepo_createApplicationCmd.Flags().String("readme-body", "", "A local text readme file in Markdown language that contains a more detailed description of the application and how it works.")
	serverlessrepo_createApplicationCmd.Flags().String("readme-url", "", "A link to the S3 object in Markdown language that contains a more detailed description of the application and how it works.")
	serverlessrepo_createApplicationCmd.Flags().String("semantic-version", "", "The semantic version of the application:")
	serverlessrepo_createApplicationCmd.Flags().String("source-code-archive-url", "", "A link to the S3 object that contains the ZIP archive of the source code for this version of your application.")
	serverlessrepo_createApplicationCmd.Flags().String("source-code-url", "", "A link to a public repository for the source code of your application, for example the URL of a specific GitHub commit.")
	serverlessrepo_createApplicationCmd.Flags().String("spdx-license-id", "", "A valid identifier from [https://spdx.org/licenses/](https://spdx.org/licenses/).")
	serverlessrepo_createApplicationCmd.Flags().String("template-body", "", "The local raw packaged AWS SAM template file of your application.")
	serverlessrepo_createApplicationCmd.Flags().String("template-url", "", "A link to the S3 object containing the packaged AWS SAM template of your application.")
	serverlessrepo_createApplicationCmd.MarkFlagRequired("author")
	serverlessrepo_createApplicationCmd.MarkFlagRequired("description")
	serverlessrepo_createApplicationCmd.MarkFlagRequired("name")
	serverlessrepoCmd.AddCommand(serverlessrepo_createApplicationCmd)
}
