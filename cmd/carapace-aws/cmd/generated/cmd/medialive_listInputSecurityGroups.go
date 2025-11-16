package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listInputSecurityGroupsCmd = &cobra.Command{
	Use:   "list-input-security-groups",
	Short: "Produces a list of Input Security Groups for an account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listInputSecurityGroupsCmd).Standalone()

	medialive_listInputSecurityGroupsCmd.Flags().String("max-results", "", "")
	medialive_listInputSecurityGroupsCmd.Flags().String("next-token", "", "")
	medialiveCmd.AddCommand(medialive_listInputSecurityGroupsCmd)
}
