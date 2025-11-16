package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_updateDomainNameCmd = &cobra.Command{
	Use:   "update-domain-name",
	Short: "Updates a custom `DomainName` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_updateDomainNameCmd).Standalone()

	appsync_updateDomainNameCmd.Flags().String("description", "", "A description of the `DomainName`.")
	appsync_updateDomainNameCmd.Flags().String("domain-name", "", "The domain name.")
	appsync_updateDomainNameCmd.MarkFlagRequired("domain-name")
	appsyncCmd.AddCommand(appsync_updateDomainNameCmd)
}
