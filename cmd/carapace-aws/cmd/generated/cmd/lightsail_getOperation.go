package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getOperationCmd = &cobra.Command{
	Use:   "get-operation",
	Short: "Returns information about a specific operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getOperationCmd).Standalone()

	lightsail_getOperationCmd.Flags().String("operation-id", "", "A GUID used to identify the operation.")
	lightsail_getOperationCmd.MarkFlagRequired("operation-id")
	lightsailCmd.AddCommand(lightsail_getOperationCmd)
}
