package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getOperationsCmd = &cobra.Command{
	Use:   "get-operations",
	Short: "Returns information about all operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getOperationsCmd).Standalone()

	lightsail_getOperationsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getOperationsCmd)
}
