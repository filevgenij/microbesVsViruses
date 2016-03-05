package main

import "fmt"
import "math/big"
import "flag"

func main() {

	microbesAmountPtr := flag.Int("microbesAmount", 100, "Amount of microbes at the begining")
	virusesAmountPtr := flag.Int("virusesAmount", 1, "Amount of viruses at the begining")

	flag.Parse()

	microbesCnt := big.NewInt(int64(*microbesAmountPtr))
	virusesCnt := big.NewInt(int64(*virusesAmountPtr))
	fmt.Println("Lets play :)")
	fmt.Println("Microbes amount -", microbesCnt, ", viruses amount -", virusesCnt)
	steps := 0
	//for microbesCnt >= 0 {
	for microbesCnt.Sign() > 0 {
		//microbesCnt -= virusesCnt
		microbesCnt.Sub(microbesCnt, virusesCnt)
		if microbesCnt.Sign() == -1 {
			microbesCnt.Set(big.NewInt(0))
		}
		//microbesCnt *= 2
		microbesCnt.Mul(microbesCnt, big.NewInt(2))
		//virusesCnt *= 2
		virusesCnt.Mul(virusesCnt, big.NewInt(2))
		fmt.Println("Microbes amount -", microbesCnt, ", viruses amount -", virusesCnt)
		steps += 1
	}
	fmt.Println("Viruses win in", steps, "steps")
}
