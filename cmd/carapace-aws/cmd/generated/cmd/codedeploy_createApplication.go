package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_createApplicationCmd).Standalone()

		codedeploy_createApplicationCmd.Flags().String("application-name", "", "The name of the application.")
		codedeploy_createApplicationCmd.Flags().String("compute-platform", "", "The destination platform type for the deployment (`Lambda`, `Server`, or `ECS`).")
		codedeploy_createApplicationCmd.Flags().String("tags", "", "The metadata that you apply to CodeDeploy applications to help you organize and categorize them.")
		codedeploy_createApplicationCmd.MarkFlagRequired("application-name")
	})
	codedeployCmd.AddCommand(codedeploy_createApplicationCmd)
}
