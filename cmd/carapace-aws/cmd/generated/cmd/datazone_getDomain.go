package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getDomainCmd = &cobra.Command{
	Use:   "get-domain",
	Short: "Gets an Amazon DataZone domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getDomainCmd).Standalone()

	datazone_getDomainCmd.Flags().String("identifier", "", "The identifier of the specified Amazon DataZone domain.")
	datazone_getDomainCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getDomainCmd)
}
