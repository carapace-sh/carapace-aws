package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_registerApplicationRevisionCmd = &cobra.Command{
	Use:   "register-application-revision",
	Short: "Registers with CodeDeploy a revision for the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_registerApplicationRevisionCmd).Standalone()

	codedeploy_registerApplicationRevisionCmd.Flags().String("application-name", "", "The name of an CodeDeploy application associated with the user or Amazon Web Services account.")
	codedeploy_registerApplicationRevisionCmd.Flags().String("description", "", "A comment about the revision.")
	codedeploy_registerApplicationRevisionCmd.Flags().String("revision", "", "Information about the application revision to register, including type and location.")
	codedeploy_registerApplicationRevisionCmd.MarkFlagRequired("application-name")
	codedeploy_registerApplicationRevisionCmd.MarkFlagRequired("revision")
	codedeployCmd.AddCommand(codedeploy_registerApplicationRevisionCmd)
}
