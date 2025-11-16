package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_getApplicationRevisionCmd = &cobra.Command{
	Use:   "get-application-revision",
	Short: "Gets information about an application revision.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_getApplicationRevisionCmd).Standalone()

	codedeploy_getApplicationRevisionCmd.Flags().String("application-name", "", "The name of the application that corresponds to the revision.")
	codedeploy_getApplicationRevisionCmd.Flags().String("revision", "", "Information about the application revision to get, including type and location.")
	codedeploy_getApplicationRevisionCmd.MarkFlagRequired("application-name")
	codedeploy_getApplicationRevisionCmd.MarkFlagRequired("revision")
	codedeployCmd.AddCommand(codedeploy_getApplicationRevisionCmd)
}
