package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func Test2047(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			sentence := test.Val.(string)

			got := countValidWords(sentence)
			want := test.Want.(int)

			if got != want {
				t.Errorf("got %v, want %v\n", got, want)
			}
		})
	}
}

func Tests() []T {
	return []T{
		{
			Num:  1,
			Val:  "cat and  dog",
			Want: 3,
		},
		{
			Num:  2,
			Val:  "!this  1-s b8d!",
			Want: 0,
		},
		{
			Num:  3,
			Val:  "alice and  bob are playing stone-game10",
			Want: 5,
		},
		{
			Num:  4,
			Val:  "he bought 2 pencils, 3 erasers, and 1  pencil-sharpener.",
			Want: 6,
		},
		{
			Num:  5,
			Val:  "a",
			Want: 1,
		},
		{
			Num:  6,
			Val:  "-",
			Want: 0,
		},
		{
			Num:  7,
			Val:  "!",
			Want: 1,
		},
		{
			Num:  8,
			Val:  "a-b-c",
			Want: 0,
		},
		{
			Num:  9,
			Val:  ". ! 7hk  al6 l! aon49esj35la k3 7u2tkh  7i9y5  !jyylhppd et v- h!ogsouv 5",
			Want: 4,
		},
		{
			Num:  10,
			Val:  ",ee3-jxp 6i! r  -  tk, h yh3h w-i10cws gl0   h ckd9drxyh !mr jii jaoj . na,b5a2v.!s2 pan e253yo87mvacrm ysw 7-e  7n.a!fr e6nxcxb axs  !.  ,  v2bz,p.t9iw8wu!  4q36au.zl8 39na dn rvdpfys   1qj48pi c    l6v -fqfd 3c  3 ytn,xm   !53y2a m 8fqq- 3 2ral f jhj v o  4!8  .3 p ,ijhq b2la89v5 xzax!e bjw  nq qwu! eod qqwnf 2sc mw0  ko r fi7p 0e9jc4!7bw,ki    ojf uo 7-jf1swt70! wr  3ahsb xs! v cb.h   ybb is4cx71  4r qmy  qi7rn r i0apj og z  tp545by!ct9h 8pugwy   ipyn.tfi6o 3 4 2f. l 1ex2 l9 a  5nn  s4m!xb2if 3  4dj4  7  7 4dxe g pu3 -nd1qb x x-ucx-7,455 ,cy  egdz  xvutn  p! n e,u  qgs  k48-gq7t52n wasqim8u -k yp 26z ux sgpwn5cm6 5m dfbkej pr g x t1ww10 s -d dh   10  -      -kt gb   1et !8 f!3w 8fe7czp hsxy7u6o -wu4hcbijze 27m5 6l 3vx 7oq! 1z8 7-,.t--   oat   -g2!.  o !ri72ox w7 p ko wi4kfx7vpd fq4 zffk eqvu6, xox57 f75vo1m  ln9  fw 07 d4 .  s3e ofwam1fd!e8n  p yenyky5p   09  q  pqs q1v2fdwi a4vm",
			Want: 65,
		},
	}
}
