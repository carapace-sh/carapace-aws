package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listMembersCmd = &cobra.Command{
	Use:   "list-members",
	Short: "List members associated with the Amazon Inspector delegated administrator for your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_listMembersCmd).Standalone()

		inspector2_listMembersCmd.Flags().String("max-results", "", "The maximum number of results the response can return.")
		inspector2_listMembersCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
		inspector2_listMembersCmd.Flags().Bool("no-only-associated", false, "Specifies whether to list only currently associated members if `True` or to list all members within the organization if `False`.")
		inspector2_listMembersCmd.Flags().Bool("only-associated", false, "Specifies whether to list only currently associated members if `True` or to list all members within the organization if `False`.")
		inspector2_listMembersCmd.Flag("no-only-associated").Hidden = true
	})
	inspector2Cmd.AddCommand(inspector2_listMembersCmd)
}
