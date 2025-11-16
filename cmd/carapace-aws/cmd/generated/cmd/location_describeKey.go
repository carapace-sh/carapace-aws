package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_describeKeyCmd = &cobra.Command{
	Use:   "describe-key",
	Short: "Retrieves the API key resource details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_describeKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_describeKeyCmd).Standalone()

		location_describeKeyCmd.Flags().String("key-name", "", "The name of the API key resource.")
		location_describeKeyCmd.MarkFlagRequired("key-name")
	})
	locationCmd.AddCommand(location_describeKeyCmd)
}
