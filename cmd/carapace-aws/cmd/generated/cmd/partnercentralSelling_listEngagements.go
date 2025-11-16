package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_listEngagementsCmd = &cobra.Command{
	Use:   "list-engagements",
	Short: "This action allows users to retrieve a list of Engagement records from Partner Central.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_listEngagementsCmd).Standalone()

	partnercentralSelling_listEngagementsCmd.Flags().String("catalog", "", "Specifies the catalog related to the request.")
	partnercentralSelling_listEngagementsCmd.Flags().String("created-by", "", "A list of AWS account IDs.")
	partnercentralSelling_listEngagementsCmd.Flags().String("engagement-identifier", "", "An array of strings representing engagement identifiers to retrieve.")
	partnercentralSelling_listEngagementsCmd.Flags().String("exclude-created-by", "", "An array of strings representing AWS Account IDs.")
	partnercentralSelling_listEngagementsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	partnercentralSelling_listEngagementsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	partnercentralSelling_listEngagementsCmd.Flags().String("sort", "", "")
	partnercentralSelling_listEngagementsCmd.MarkFlagRequired("catalog")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_listEngagementsCmd)
}
