package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeInputSecurityGroupCmd = &cobra.Command{
	Use:   "describe-input-security-group",
	Short: "Produces a summary of an Input Security Group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeInputSecurityGroupCmd).Standalone()

	medialive_describeInputSecurityGroupCmd.Flags().String("input-security-group-id", "", "The id of the Input Security Group to describe")
	medialive_describeInputSecurityGroupCmd.MarkFlagRequired("input-security-group-id")
	medialiveCmd.AddCommand(medialive_describeInputSecurityGroupCmd)
}
