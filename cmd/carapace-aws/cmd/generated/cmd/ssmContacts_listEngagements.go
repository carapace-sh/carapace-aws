package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_listEngagementsCmd = &cobra.Command{
	Use:   "list-engagements",
	Short: "Lists all engagements that have happened in an incident.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_listEngagementsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_listEngagementsCmd).Standalone()

		ssmContacts_listEngagementsCmd.Flags().String("incident-id", "", "The Amazon Resource Name (ARN) of the incident you're listing engagements for.")
		ssmContacts_listEngagementsCmd.Flags().String("max-results", "", "The maximum number of engagements per page of results.")
		ssmContacts_listEngagementsCmd.Flags().String("next-token", "", "The pagination token to continue to the next page of results.")
		ssmContacts_listEngagementsCmd.Flags().String("time-range-value", "", "The time range to lists engagements for an incident.")
	})
	ssmContactsCmd.AddCommand(ssmContacts_listEngagementsCmd)
}
