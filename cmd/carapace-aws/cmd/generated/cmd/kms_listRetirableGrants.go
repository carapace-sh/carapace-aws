package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_listRetirableGrantsCmd = &cobra.Command{
	Use:   "list-retirable-grants",
	Short: "Returns information about all grants in the Amazon Web Services account and Region that have the specified retiring principal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_listRetirableGrantsCmd).Standalone()

	kms_listRetirableGrantsCmd.Flags().String("limit", "", "Use this parameter to specify the maximum number of items to return.")
	kms_listRetirableGrantsCmd.Flags().String("marker", "", "Use this parameter in a subsequent request after you receive a response with truncated results.")
	kms_listRetirableGrantsCmd.Flags().String("retiring-principal", "", "The retiring principal for which to list grants.")
	kms_listRetirableGrantsCmd.MarkFlagRequired("retiring-principal")
	kmsCmd.AddCommand(kms_listRetirableGrantsCmd)
}
