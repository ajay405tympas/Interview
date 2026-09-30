// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func hasDuplicates(str string) bool {

	seen := make(map[byte]bool)

	for i := 0; i < len(str); i++ {
		if seen[str[i]] {
			return false
		}
		seen[str[i]] = true
	}
	return true
}

func lengthOfLongestSubstring(s string) int {

	temp := ""
	lstr := ""

	for k := 1; k <= len(s); k++ {
		for i := 0; i < len(s)-k+1; i++ {
			for j := i; j <= k+i; j++ {
				temp = s[i:j]
			}

			if hasDuplicates(temp) {
				lstr = temp
				fmt.Println("lstr:", lstr)
			}
		}

	}

	return len(lstr)
}

func main() {

	s := "/\\y_E2?p%@Aw?5Stv4:(ivY+\\]aJ&.1)Qg2IRh5S`@rHI'2@r-oIDrcUKIc$YxNHAk(GJVzJ8L-w,q)Qs\\y0\\<t5CG+MLT/wTNt\"(qL9Uba+?nUk:8R:xSXtTqi>@wXcn .O;s#U-[dK7f'(oZ51Y}%+2h4f)@/sb/)+f,69.(.Pn]y|G!Luf>]^y!7>!^<6tK xi)W;&6}7L4T.P@i ?vtJDA4u?TW%PMAK/8eo7V.(CF>4LHB^j_)feW/@/m!i<^g}y9z5 h >g<1Ebz5x+{UBKVDaz;_LD8;@{,/q706uW=47\"i2* H+8z];r%:#LFrR)|GCN0}K1#{bKEKfyCFZ*='WV*vAr|iQNdkqlG/k%txjM-+]wH8r#\\yIgthKw\\ L),FdIWAF|3809+g\"$ VULtf)gF$vE!dWK]ueCiPg^A^2<MFx'9J9-rJcc]l $Ym$[TYY`b7gJ aG -u8w\"@Yz*f^sP<h:'CUXDP}ylsUwE-'u,>*zFWSy01W2Gi7Q\"0YvD9VI6b`p8ylK&&0yTT1K;s9r0E?p>T\\9V3=~xA)EKUGA}dwk]>bzE;ScV6`Bwn\"Xc"
	fmt.Println("Max Length is:", lengthOfLongestSubstring(s))

	fmt.Println("Hello, 世界")
}
