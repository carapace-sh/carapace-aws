package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_getApplicationCmd = &cobra.Command{
	Use:   "get-application",
	Short: "Gets information about an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_getApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_getApplicationCmd).Standalone()

		codedeploy_getApplicationCmd.Flags().String("application-name", "", "The name of an CodeDeploy application associated with the user or Amazon Web Services account.")
		codedeploy_getApplicationCmd.MarkFlagRequired("application-name")
	})
	codedeployCmd.AddCommand(codedeploy_getApplicationCmd)
}
