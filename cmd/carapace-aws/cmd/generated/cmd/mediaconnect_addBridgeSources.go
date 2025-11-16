package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_addBridgeSourcesCmd = &cobra.Command{
	Use:   "add-bridge-sources",
	Short: "Adds sources to an existing bridge.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_addBridgeSourcesCmd).Standalone()

	mediaconnect_addBridgeSourcesCmd.Flags().String("bridge-arn", "", "The Amazon Resource Name (ARN) of the bridge that you want to update.")
	mediaconnect_addBridgeSourcesCmd.Flags().String("sources", "", "The sources that you want to add to this bridge.")
	mediaconnect_addBridgeSourcesCmd.MarkFlagRequired("bridge-arn")
	mediaconnect_addBridgeSourcesCmd.MarkFlagRequired("sources")
	mediaconnectCmd.AddCommand(mediaconnect_addBridgeSourcesCmd)
}
