package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_listIpRoutesCmd = &cobra.Command{
	Use:   "list-ip-routes",
	Short: "Lists the address blocks that you have added to a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_listIpRoutesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_listIpRoutesCmd).Standalone()

		ds_listIpRoutesCmd.Flags().String("directory-id", "", "Identifier (ID) of the directory for which you want to retrieve the IP addresses.")
		ds_listIpRoutesCmd.Flags().String("limit", "", "Maximum number of items to return.")
		ds_listIpRoutesCmd.Flags().String("next-token", "", "The *ListIpRoutes.NextToken* value from a previous call to [ListIpRoutes]().")
		ds_listIpRoutesCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_listIpRoutesCmd)
}
