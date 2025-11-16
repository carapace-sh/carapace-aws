package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Changes the name of an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_updateApplicationCmd).Standalone()

	codedeploy_updateApplicationCmd.Flags().String("application-name", "", "The current name of the application you want to change.")
	codedeploy_updateApplicationCmd.Flags().String("new-application-name", "", "The new name to give the application.")
	codedeployCmd.AddCommand(codedeploy_updateApplicationCmd)
}
