package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getAccountPoolCmd = &cobra.Command{
	Use:   "get-account-pool",
	Short: "Gets the details of the account pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getAccountPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getAccountPoolCmd).Standalone()

		datazone_getAccountPoolCmd.Flags().String("domain-identifier", "", "The ID of the domain in which the account pool lives whose details are to be displayed.")
		datazone_getAccountPoolCmd.Flags().String("identifier", "", "The ID of the account pool whose details are to be displayed.")
		datazone_getAccountPoolCmd.MarkFlagRequired("domain-identifier")
		datazone_getAccountPoolCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getAccountPoolCmd)
}
