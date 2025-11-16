package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_deleteDomainNameCmd = &cobra.Command{
	Use:   "delete-domain-name",
	Short: "Deletes a custom `DomainName` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_deleteDomainNameCmd).Standalone()

	appsync_deleteDomainNameCmd.Flags().String("domain-name", "", "The domain name.")
	appsync_deleteDomainNameCmd.MarkFlagRequired("domain-name")
	appsyncCmd.AddCommand(appsync_deleteDomainNameCmd)
}
