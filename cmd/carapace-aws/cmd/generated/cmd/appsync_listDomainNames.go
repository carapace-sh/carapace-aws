package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_listDomainNamesCmd = &cobra.Command{
	Use:   "list-domain-names",
	Short: "Lists multiple custom domain names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_listDomainNamesCmd).Standalone()

	appsync_listDomainNamesCmd.Flags().String("max-results", "", "The maximum number of results that you want the request to return.")
	appsync_listDomainNamesCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which you can use to return the next set of items in the list.")
	appsyncCmd.AddCommand(appsync_listDomainNamesCmd)
}
