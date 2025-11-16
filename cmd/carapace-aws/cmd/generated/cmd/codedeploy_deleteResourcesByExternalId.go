package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_deleteResourcesByExternalIdCmd = &cobra.Command{
	Use:   "delete-resources-by-external-id",
	Short: "Deletes resources linked to an external ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_deleteResourcesByExternalIdCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_deleteResourcesByExternalIdCmd).Standalone()

		codedeploy_deleteResourcesByExternalIdCmd.Flags().String("external-id", "", "The unique ID of an external resource (for example, a CloudFormation stack ID) that is linked to one or more CodeDeploy resources.")
	})
	codedeployCmd.AddCommand(codedeploy_deleteResourcesByExternalIdCmd)
}
