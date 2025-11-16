package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listIngressPointsCmd = &cobra.Command{
	Use:   "list-ingress-points",
	Short: "List all ingress endpoint resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listIngressPointsCmd).Standalone()

	mailmanager_listIngressPointsCmd.Flags().String("next-token", "", "If you received a pagination token from a previous call to this API, you can provide it here to continue paginating through the next page of results.")
	mailmanager_listIngressPointsCmd.Flags().String("page-size", "", "The maximum number of ingress endpoint resources that are returned per call.")
	mailmanagerCmd.AddCommand(mailmanager_listIngressPointsCmd)
}
