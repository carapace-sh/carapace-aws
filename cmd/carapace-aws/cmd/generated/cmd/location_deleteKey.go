package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_deleteKeyCmd = &cobra.Command{
	Use:   "delete-key",
	Short: "Deletes the specified API key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_deleteKeyCmd).Standalone()

	location_deleteKeyCmd.Flags().Bool("force-delete", false, "ForceDelete bypasses an API key's expiry conditions and deletes the key.")
	location_deleteKeyCmd.Flags().String("key-name", "", "The name of the API key to delete.")
	location_deleteKeyCmd.Flags().Bool("no-force-delete", false, "ForceDelete bypasses an API key's expiry conditions and deletes the key.")
	location_deleteKeyCmd.MarkFlagRequired("key-name")
	location_deleteKeyCmd.Flag("no-force-delete").Hidden = true
	locationCmd.AddCommand(location_deleteKeyCmd)
}
