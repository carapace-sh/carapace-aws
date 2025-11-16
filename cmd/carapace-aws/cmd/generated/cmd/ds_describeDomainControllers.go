package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeDomainControllersCmd = &cobra.Command{
	Use:   "describe-domain-controllers",
	Short: "Provides information about any domain controllers in your directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeDomainControllersCmd).Standalone()

	ds_describeDomainControllersCmd.Flags().String("directory-id", "", "Identifier of the directory for which to retrieve the domain controller information.")
	ds_describeDomainControllersCmd.Flags().String("domain-controller-ids", "", "A list of identifiers for the domain controllers whose information will be provided.")
	ds_describeDomainControllersCmd.Flags().String("limit", "", "The maximum number of items to return.")
	ds_describeDomainControllersCmd.Flags().String("next-token", "", "The *DescribeDomainControllers.NextToken* value from a previous call to [DescribeDomainControllers]().")
	ds_describeDomainControllersCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_describeDomainControllersCmd)
}
