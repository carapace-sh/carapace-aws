package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_getCapabilityCmd = &cobra.Command{
	Use:   "get-capability",
	Short: "Retrieves the details for the specified capability.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_getCapabilityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_getCapabilityCmd).Standalone()

		b2bi_getCapabilityCmd.Flags().String("capability-id", "", "Specifies a system-assigned unique identifier for the capability.")
		b2bi_getCapabilityCmd.MarkFlagRequired("capability-id")
	})
	b2biCmd.AddCommand(b2bi_getCapabilityCmd)
}
