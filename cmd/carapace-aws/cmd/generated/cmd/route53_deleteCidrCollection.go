package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_deleteCidrCollectionCmd = &cobra.Command{
	Use:   "delete-cidr-collection",
	Short: "Deletes a CIDR collection in the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_deleteCidrCollectionCmd).Standalone()

	route53_deleteCidrCollectionCmd.Flags().String("id", "", "The UUID of the collection to delete.")
	route53_deleteCidrCollectionCmd.MarkFlagRequired("id")
	route53Cmd.AddCommand(route53_deleteCidrCollectionCmd)
}
