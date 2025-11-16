package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_createCloudFormationTemplateCmd = &cobra.Command{
	Use:   "create-cloud-formation-template",
	Short: "Creates an AWS CloudFormation template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_createCloudFormationTemplateCmd).Standalone()

	serverlessrepo_createCloudFormationTemplateCmd.Flags().String("application-id", "", "The Amazon Resource Name (ARN) of the application.")
	serverlessrepo_createCloudFormationTemplateCmd.Flags().String("semantic-version", "", "The semantic version of the application:")
	serverlessrepo_createCloudFormationTemplateCmd.MarkFlagRequired("application-id")
	serverlessrepoCmd.AddCommand(serverlessrepo_createCloudFormationTemplateCmd)
}
