package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_deleteCapabilityCmd = &cobra.Command{
	Use:   "delete-capability",
	Short: "Deletes the specified capability.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_deleteCapabilityCmd).Standalone()

	b2bi_deleteCapabilityCmd.Flags().String("capability-id", "", "Specifies a system-assigned unique identifier for the capability.")
	b2bi_deleteCapabilityCmd.MarkFlagRequired("capability-id")
	b2biCmd.AddCommand(b2bi_deleteCapabilityCmd)
}
