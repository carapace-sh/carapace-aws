package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listResourceDelegatesCmd = &cobra.Command{
	Use:   "list-resource-delegates",
	Short: "Lists the delegates associated with a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listResourceDelegatesCmd).Standalone()

	workmail_listResourceDelegatesCmd.Flags().String("max-results", "", "The number of maximum results in a page.")
	workmail_listResourceDelegatesCmd.Flags().String("next-token", "", "The token used to paginate through the delegates associated with a resource.")
	workmail_listResourceDelegatesCmd.Flags().String("organization-id", "", "The identifier for the organization that contains the resource for which delegates are listed.")
	workmail_listResourceDelegatesCmd.Flags().String("resource-id", "", "The identifier for the resource whose delegates are listed.")
	workmail_listResourceDelegatesCmd.MarkFlagRequired("organization-id")
	workmail_listResourceDelegatesCmd.MarkFlagRequired("resource-id")
	workmailCmd.AddCommand(workmail_listResourceDelegatesCmd)
}
