package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_listConnectionsCmd = &cobra.Command{
	Use:   "list-connections",
	Short: "Returns a list of App Runner connections that are associated with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_listConnectionsCmd).Standalone()

	apprunner_listConnectionsCmd.Flags().String("connection-name", "", "If specified, only this connection is returned.")
	apprunner_listConnectionsCmd.Flags().String("max-results", "", "The maximum number of results to include in each response (result page).")
	apprunner_listConnectionsCmd.Flags().String("next-token", "", "A token from a previous result page.")
	apprunnerCmd.AddCommand(apprunner_listConnectionsCmd)
}
