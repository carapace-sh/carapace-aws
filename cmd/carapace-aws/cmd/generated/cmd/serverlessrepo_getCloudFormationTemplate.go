package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_getCloudFormationTemplateCmd = &cobra.Command{
	Use:   "get-cloud-formation-template",
	Short: "Gets the specified AWS CloudFormation template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_getCloudFormationTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serverlessrepo_getCloudFormationTemplateCmd).Standalone()

		serverlessrepo_getCloudFormationTemplateCmd.Flags().String("application-id", "", "The Amazon Resource Name (ARN) of the application.")
		serverlessrepo_getCloudFormationTemplateCmd.Flags().String("template-id", "", "The UUID returned by CreateCloudFormationTemplate.")
		serverlessrepo_getCloudFormationTemplateCmd.MarkFlagRequired("application-id")
		serverlessrepo_getCloudFormationTemplateCmd.MarkFlagRequired("template-id")
	})
	serverlessrepoCmd.AddCommand(serverlessrepo_getCloudFormationTemplateCmd)
}
