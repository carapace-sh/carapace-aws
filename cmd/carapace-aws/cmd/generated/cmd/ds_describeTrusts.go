package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeTrustsCmd = &cobra.Command{
	Use:   "describe-trusts",
	Short: "Obtains information about the trust relationships for this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeTrustsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_describeTrustsCmd).Standalone()

		ds_describeTrustsCmd.Flags().String("directory-id", "", "The Directory ID of the Amazon Web Services directory that is a part of the requested trust relationship.")
		ds_describeTrustsCmd.Flags().String("limit", "", "The maximum number of objects to return.")
		ds_describeTrustsCmd.Flags().String("next-token", "", "The *DescribeTrustsResult.NextToken* value from a previous call to [DescribeTrusts]().")
		ds_describeTrustsCmd.Flags().String("trust-ids", "", "A list of identifiers of the trust relationships for which to obtain the information.")
	})
	dsCmd.AddCommand(ds_describeTrustsCmd)
}
