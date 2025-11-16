package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getDomainNameCmd = &cobra.Command{
	Use:   "get-domain-name",
	Short: "Retrieves a custom `DomainName` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getDomainNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_getDomainNameCmd).Standalone()

		appsync_getDomainNameCmd.Flags().String("domain-name", "", "The domain name.")
		appsync_getDomainNameCmd.MarkFlagRequired("domain-name")
	})
	appsyncCmd.AddCommand(appsync_getDomainNameCmd)
}
