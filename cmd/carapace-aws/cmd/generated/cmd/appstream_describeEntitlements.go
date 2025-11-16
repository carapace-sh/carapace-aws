package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeEntitlementsCmd = &cobra.Command{
	Use:   "describe-entitlements",
	Short: "Retrieves a list that describes one of more entitlements.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeEntitlementsCmd).Standalone()

	appstream_describeEntitlementsCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
	appstream_describeEntitlementsCmd.Flags().String("name", "", "The name of the entitlement.")
	appstream_describeEntitlementsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	appstream_describeEntitlementsCmd.Flags().String("stack-name", "", "The name of the stack with which the entitlement is associated.")
	appstream_describeEntitlementsCmd.MarkFlagRequired("stack-name")
	appstreamCmd.AddCommand(appstream_describeEntitlementsCmd)
}
