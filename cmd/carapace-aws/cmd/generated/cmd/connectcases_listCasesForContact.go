package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_listCasesForContactCmd = &cobra.Command{
	Use:   "list-cases-for-contact",
	Short: "Lists cases for a given contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_listCasesForContactCmd).Standalone()

	connectcases_listCasesForContactCmd.Flags().String("contact-arn", "", "A unique identifier of a contact in Amazon Connect.")
	connectcases_listCasesForContactCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_listCasesForContactCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connectcases_listCasesForContactCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connectcases_listCasesForContactCmd.MarkFlagRequired("contact-arn")
	connectcases_listCasesForContactCmd.MarkFlagRequired("domain-id")
	connectcasesCmd.AddCommand(connectcases_listCasesForContactCmd)
}
