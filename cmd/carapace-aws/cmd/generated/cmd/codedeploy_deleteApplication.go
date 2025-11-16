package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_deleteApplicationCmd).Standalone()

	codedeploy_deleteApplicationCmd.Flags().String("application-name", "", "The name of an CodeDeploy application associated with the user or Amazon Web Services account.")
	codedeploy_deleteApplicationCmd.MarkFlagRequired("application-name")
	codedeployCmd.AddCommand(codedeploy_deleteApplicationCmd)
}
