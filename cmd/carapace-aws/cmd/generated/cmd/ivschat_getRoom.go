package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_getRoomCmd = &cobra.Command{
	Use:   "get-room",
	Short: "Gets the specified room.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_getRoomCmd).Standalone()

	ivschat_getRoomCmd.Flags().String("identifier", "", "Identifier of the room for which the configuration is to be retrieved.")
	ivschat_getRoomCmd.MarkFlagRequired("identifier")
	ivschatCmd.AddCommand(ivschat_getRoomCmd)
}
