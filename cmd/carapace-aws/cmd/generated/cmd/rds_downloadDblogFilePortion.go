package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_downloadDblogFilePortionCmd = &cobra.Command{
	Use:   "download-dblog-file-portion",
	Short: "Downloads all or a portion of the specified log file, up to 1 MB in size.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_downloadDblogFilePortionCmd).Standalone()

	rds_downloadDblogFilePortionCmd.Flags().String("dbinstance-identifier", "", "The customer-assigned name of the DB instance that contains the log files you want to list.")
	rds_downloadDblogFilePortionCmd.Flags().String("log-file-name", "", "The name of the log file to be downloaded.")
	rds_downloadDblogFilePortionCmd.Flags().String("marker", "", "The pagination token provided in the previous request or \"0\".")
	rds_downloadDblogFilePortionCmd.Flags().String("number-of-lines", "", "The number of lines to download.")
	rds_downloadDblogFilePortionCmd.MarkFlagRequired("dbinstance-identifier")
	rds_downloadDblogFilePortionCmd.MarkFlagRequired("log-file-name")
	rdsCmd.AddCommand(rds_downloadDblogFilePortionCmd)
}
