package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random password",
	Long: `Generate a random password with options
	
		For example:

		go-password generate -l 10 -d -s
	`,
	Run: func(cmd *cobra.Command, args []string) {

		length, _ := cmd.Flags().GetInt("length")

		isDigits, _ := cmd.Flags().GetBool("digits")

		isSymbols, _ := cmd.Flags().GetBool("symbols")

		charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

		if isDigits {
			charset += "0123456789"
		}

		if isSymbols {
			charset += "!@#$%^&*()_+{}[]"
		}

		password := make([]byte, length)

		for i := range password {
			password[i] = charset[rand.Intn(len(charset))]

		}

		fmt.Println(string(password))

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "Length of the password")

	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in the password")

	generateCmd.Flags().BoolP("symbols", "s", false, "Include symbols in the password")

}
