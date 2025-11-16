package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_deleteTransformerCmd = &cobra.Command{
	Use:   "delete-transformer",
	Short: "Deletes the specified transformer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_deleteTransformerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_deleteTransformerCmd).Standalone()

		b2bi_deleteTransformerCmd.Flags().String("transformer-id", "", "Specifies the system-assigned unique identifier for the transformer.")
		b2bi_deleteTransformerCmd.MarkFlagRequired("transformer-id")
	})
	b2biCmd.AddCommand(b2bi_deleteTransformerCmd)
}
