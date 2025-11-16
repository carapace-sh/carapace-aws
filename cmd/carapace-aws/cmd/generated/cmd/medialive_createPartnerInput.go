package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createPartnerInputCmd = &cobra.Command{
	Use:   "create-partner-input",
	Short: "Create a partner input",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createPartnerInputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_createPartnerInputCmd).Standalone()

		medialive_createPartnerInputCmd.Flags().String("input-id", "", "Unique ID of the input.")
		medialive_createPartnerInputCmd.Flags().String("request-id", "", "Unique identifier of the request to ensure the request is handled exactly once in case of retries.")
		medialive_createPartnerInputCmd.Flags().String("tags", "", "A collection of key-value pairs.")
		medialive_createPartnerInputCmd.MarkFlagRequired("input-id")
	})
	medialiveCmd.AddCommand(medialive_createPartnerInputCmd)
}
