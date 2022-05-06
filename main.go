package main

import "fmt"
// import "piscine"

func output(str string, board [][]string) {
	fmt.Print(str);
	fmt.Println("=========\n");
	for h := 0; h < len(board); h++{ 
		for i := 0; i < len(board[h]); i++{ 
			for j := 0; j < len(board[h][i]); j++{
				fmt.Print(string(board[h][i][j]))
				fmt.Print(" ")
			}
		fmt.Println("")
	}
		// piscine.Checkmate(board[h])
		fmt.Println("dummy!!\n")
		fmt.Println("")
	}
}

func main() {
	fmt.Println("//////////////////////////////////////////");
	fmt.Println("valid maps");
	fmt.Println("//////////////////////////////////////////\n");

	pdf_v1 := [][]string{
		{
			"...2.",
			"..6.4",
			"5...6",
			"7.6..",
			".3...",
		},
	}
	output("pdf_v1", pdf_v1)

	pdf_v2 := [][]string{
		{
			".....",
			"8.8..",
			".7.7.",
			"..8.5",
			".....",
		},
	}
	output("pdf_v2", pdf_v2)

	pdf_v3 := [][]string{
		{
			"37...",
			"..8..",
			".....",
			"..8..",
			"...69",
		},
	}
	output("pdf_v3", pdf_v3)

	fmt.Println("//////////////////////////////////////////");
	fmt.Println("invalid maps---------");
	fmt.Println("//////////////////////////////////////////\n");

	pdf_inv1 := [][]string{
		{
			"...",
		},
	}
	output("pdf_inv1", pdf_inv1)

	pdf_inv2 := [][]string{
		{
			"...A.",
			"..6.4",
			"5....6",
			"7.6..",
			".3...",
		},
	}
	output("pdf_inv2", pdf_inv2)
}
