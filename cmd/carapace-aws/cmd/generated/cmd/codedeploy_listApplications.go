package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Lists the applications registered with the user or Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_listApplicationsCmd).Standalone()

	codedeploy_listApplicationsCmd.Flags().String("next-token", "", "An identifier returned from the previous list applications call.")
	codedeployCmd.AddCommand(codedeploy_listApplicationsCmd)
}
