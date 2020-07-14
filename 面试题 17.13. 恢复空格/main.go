package main

import "fmt"

// 动态规划
func respace(dictionary []string, sentence string) int {
	// 字典的索引
	m := make(map[string]bool, len(dictionary))
	for _, v := range dictionary {
		m[v] = true
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// dp[i] 到第i时未识别的字符数
	// dp[i] = min(dp[i-1]+1, dp[i-j]) // [i-j,i]在directionary中
	dp := make([]int, len(sentence)+1)
	for i := 1; i <= len(sentence); i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			if m[sentence[j:i]] {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[len(sentence)]
}

func main() {
	fmt.Println(respace([]string{"looked", "just", "like", "her", "brother"}, "jesslookedjustliketimherbrother"))
	fmt.Println(respace([]string{"ouf", "uucuocucoouoffcpuuf", "pf", "o", "fofopupoufuofffffocpocfccuofuupupcouocpocoooupcuu",
		"cf", "cffooccccuoocpfupuucufoocpocucpuouofffpoupu", "opoffuoofpupcpfouoouufpcuocufo", "fopuupco", "upocfucuucfucofufu",
		"ufoccopopuouccupooc", "fffu", "uuopuccfocopooupooucfoufop", "occ", "ppfcuu", "o", "fpp", "o",
		"oououpuccuppuococcpoucccffcpcucoffupcoppoc", "ufc", "coupo", "ufuoufofopcpfoufoouppffofffuupfco", "focfcfcfcfpuouoccupfccfpcooup",
		"ffupfffccpffufuuo", "cufoupupppocou", "upopupopccffuofpcopouofpoffopcfcuooocppofofuuc", "oo", "pccc", "oupupcccppuuucuuouocu",
		"fuop", "ppuocfuppff", "focofooffpfcpcupupuuooufu", "uofupoocpf", "opufcuffopcpcfcufpcpufuupffpp", "f", "opffp", "fpccopc"},
		"fofopupoufuofffffocpocfccuofuupupcouocpocoooupcuufffufffufpccopc"))
}
