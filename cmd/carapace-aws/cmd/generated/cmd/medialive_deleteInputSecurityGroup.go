package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteInputSecurityGroupCmd = &cobra.Command{
	Use:   "delete-input-security-group",
	Short: "Deletes an Input Security Group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteInputSecurityGroupCmd).Standalone()

	medialive_deleteInputSecurityGroupCmd.Flags().String("input-security-group-id", "", "The Input Security Group to delete")
	medialive_deleteInputSecurityGroupCmd.MarkFlagRequired("input-security-group-id")
	medialiveCmd.AddCommand(medialive_deleteInputSecurityGroupCmd)
}
