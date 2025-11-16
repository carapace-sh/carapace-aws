package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSellingCmd = &cobra.Command{
	Use:   "partnercentral-selling",
	Short: "AWS Partner Central API for Selling\n\n**AWS Partner Central API for Selling Reference Guide**\n\nThis Amazon Web Services (AWS) Partner Central API reference is designed to help [AWS Partners](http://aws.amazon.com/partners/programs/) integrate Customer Relationship Management (CRM) systems with AWS Partner Central.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSellingCmd).Standalone()

	rootCmd.AddCommand(partnercentralSellingCmd)
}
