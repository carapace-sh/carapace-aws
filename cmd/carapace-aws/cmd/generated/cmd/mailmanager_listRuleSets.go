package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listRuleSetsCmd = &cobra.Command{
	Use:   "list-rule-sets",
	Short: "List rule sets for this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listRuleSetsCmd).Standalone()

	mailmanager_listRuleSetsCmd.Flags().String("next-token", "", "If you received a pagination token from a previous call to this API, you can provide it here to continue paginating through the next page of results.")
	mailmanager_listRuleSetsCmd.Flags().String("page-size", "", "The maximum number of rule set resources that are returned per call.")
	mailmanagerCmd.AddCommand(mailmanager_listRuleSetsCmd)
}
