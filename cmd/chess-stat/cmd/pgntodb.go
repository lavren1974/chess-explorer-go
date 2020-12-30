package cmd

import (
	"github.com/spf13/cobra"
	pgntodb "github.com/yafred/chess-stat/internal/pgntodb"
)

var pgnToDbCmd = &cobra.Command{
	Use:   "pgn [pgn file]",
	Short: "Parse a pgn file and feed mongo database",
	Long:  `Parse a pgn file and feed mongo database. Designed for chess.com and lichess.org`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pgntodb.ProcessFile(args[0])
	},
}

func init() {
	rootCmd.AddCommand(pgnToDbCmd)
}