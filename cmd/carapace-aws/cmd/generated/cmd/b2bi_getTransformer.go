package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_getTransformerCmd = &cobra.Command{
	Use:   "get-transformer",
	Short: "Retrieves the details for the transformer specified by the transformer ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_getTransformerCmd).Standalone()

	b2bi_getTransformerCmd.Flags().String("transformer-id", "", "Specifies the system-assigned unique identifier for the transformer.")
	b2bi_getTransformerCmd.MarkFlagRequired("transformer-id")
	b2biCmd.AddCommand(b2bi_getTransformerCmd)
}
