package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listSharesCmd = &cobra.Command{
	Use:   "list-shares",
	Short: "Retrieves the resource shares associated with an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listSharesCmd).Standalone()

	omics_listSharesCmd.Flags().String("filter", "", "Attributes that you use to filter for a specific subset of resource shares.")
	omics_listSharesCmd.Flags().String("max-results", "", "The maximum number of shares to return in one page of results.")
	omics_listSharesCmd.Flags().String("next-token", "", "Next token returned in the response of a previous ListReadSetUploadPartsRequest call.")
	omics_listSharesCmd.Flags().String("resource-owner", "", "The account that owns the resource shares.")
	omics_listSharesCmd.MarkFlagRequired("resource-owner")
	omicsCmd.AddCommand(omics_listSharesCmd)
}
